package repositories

import (
	"context"

	"github.com/Voldemort-Project/sga-service/src/domain/entities"
	"github.com/Voldemort-Project/sga-service/utils"
)

type CheckinRepositoryImpl interface {
	BaseRepositoryImpl
	GetCheckinGuestList(
		ctx context.Context,
		pagination *utils.PaginationDto,
	) (rows []entities.CheckinEntity, total int64, err error)
}
