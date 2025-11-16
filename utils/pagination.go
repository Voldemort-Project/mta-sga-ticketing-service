package utils

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

const (
	DefaultPage  = 1
	DefaultLimit = 10

	MaxLimit = 100

	FieldPage    = "page"
	FieldPerPage = "perPage"
	FieldKeyword = "keyword"
	FieldOrder   = "order"
)

type OrderDirection string

const (
	OrderDirectionAsc  OrderDirection = "asc"
	OrderDirectionDesc OrderDirection = "desc"
)

func (o OrderDirection) String() string {
	return string(o)
}

func ParseOrderDirection(direction string) OrderDirection {
	switch strings.ToLower(direction) {
	case "asc":
		return OrderDirectionAsc
	case "desc":
		return OrderDirectionDesc
	}
	return OrderDirectionAsc
}

type PaginationRequest struct {
	Page    int
	PerPage int
	Keyword string
	Order   string
}

type PaginationOrder struct {
	Column    string
	Direction OrderDirection
}

type PaginationDto struct {
	Page    int
	PerPage int
	Keyword string
	Order   []PaginationOrder
}

func (p *PaginationDto) WithScope() func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		limit := p.PerPage
		offset := (p.Page - 1) * limit
		return tx.Limit(limit).Offset(offset)
	}
}

func NewPagination(c fiber.Ctx) *PaginationDto {
	var (
		page    = c.Query(FieldPage, strconv.Itoa(DefaultPage))
		perPage = c.Query(FieldPerPage, strconv.Itoa(DefaultLimit))
		keyword = c.Query(FieldKeyword, "")
		order   = c.Query(FieldOrder, "")
	)

	intPage, _ := strconv.Atoi(page)
	intPerPage, _ := strconv.Atoi(perPage)

	orderSlice := strings.Split(order, ";")
	orderMaker := make([]PaginationOrder, len(orderSlice))

	for i, odr := range orderSlice {
		odrSlice := strings.Split(odr, ":")
		if len(orderSlice) != 2 {
			continue
		}
		column := odrSlice[0]
		direction := ParseOrderDirection(odrSlice[1])
		orderMaker[i] = PaginationOrder{
			Column:    column,
			Direction: direction,
		}
	}

	return &PaginationDto{
		Page:    intPage,
		PerPage: intPerPage,
		Keyword: keyword,
		Order:   orderMaker,
	}
}
