package pgservice

import (
	"context"
	"database/sql"
	"errors"
	"sync"

	"github.com/Voldemort-Project/sga-service/src/domain/entities"
	"github.com/Voldemort-Project/sga-service/src/domain/repositories"
	pgmodels "github.com/Voldemort-Project/sga-service/src/infra/db/postgres/models"
	"github.com/Voldemort-Project/sga-service/utils"
	"gorm.io/gorm"
)

type checkinService struct {
	db *gorm.DB
}

func NewCheckinService(db *gorm.DB) repositories.CheckinRepositoryImpl {
	return &checkinService{db: db}
}

func (s *checkinService) Upsert(ctx context.Context, data any, tx *gorm.DB) error {
	trx := s.db
	if tx != nil {
		trx = tx
	}

	payload, ok := data.(*entities.CheckinEntity)
	if !ok {
		return errors.New("data is not a checkin entity")
	}
	model := pgmodels.CheckinModel{}
	model.FromEntity(payload)

	if err := trx.
		WithContext(ctx).
		Session(&gorm.Session{FullSaveAssociations: true}).
		Save(&model).Error; err != nil {
		return err
	}
	return nil
}

func (s *checkinService) Detail(ctx context.Context, id string, tx *gorm.DB) (any, error) {
	trx := s.db
	if tx != nil {
		trx = tx
	}

	model := pgmodels.CheckinModel{}
	if err := trx.
		WithContext(ctx).
		Session(&gorm.Session{QueryFields: true}).
		Where("id = ?", id).
		First(&model).Error; err != nil {
		return nil, err
	}

	return model.ToEntity(), nil
}

func (s *checkinService) GetCheckinGuestList(
	ctx context.Context,
	pagination *utils.PaginationDto,
) ([]entities.CheckinEntity, int64, error) {
	var (
		rows        = make([]entities.CheckinEntity, 0)
		sqlRows     = make([]pgmodels.CheckinModel, 0)
		errChannels = make(chan error, 2)

		total int64
		err   error
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	tx := s.db.WithContext(ctx).Session(&gorm.Session{QueryFields: true, PrepareStmt: true})
	tx = tx.Preload("User").Preload("Room").Preload("Organization")

	if pagination.Keyword != "" {
		tx = tx.Where("guest_name ILIKE @keyword OR guest_email ILIKE @keyword OR guest_phone ILIKE @keyword OR guest_id_card ILIKE @keyword", sql.Named("keyword", "%"+pagination.Keyword+"%"))
	}

	execSelect := func(t *gorm.DB, wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()

		mu.Lock()
		defer mu.Unlock()

		if err := t.Find(&sqlRows).Error; err != nil {
			errChannels <- err
		}
	}

	execCount := func(t *gorm.DB, wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()

		mu.Lock()
		defer mu.Unlock()

		if err := t.Model(&pgmodels.CheckinModel{}).Count(&total).Error; err != nil {
			errChannels <- err
		}
	}

	wg.Add(2)
	go execSelect(tx, &wg, &mutex)
	go execCount(tx, &wg, &mutex)

	go func() {
		wg.Wait()

		close(errChannels)

		mutex.Lock()
		defer mutex.Unlock()

		for errs := range errChannels {
			err = errs
			total = 0
			break
		}
	}()

	wg.Wait()

	if err != nil {
		return nil, 0, err
	}

	for _, row := range sqlRows {
		rows = append(rows, *row.ToEntity())
	}

	return rows, total, nil
}
