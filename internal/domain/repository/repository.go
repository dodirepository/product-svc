package domain

import domain "github.com/dodirepository/product-svc/internal/domain/usecases"

type ItemsCategory struct {
	ID           int64   `gorm:"column:id;primaryKey"`
	Name         string  `gorm:"column:name"`
	Desc         string  `gorm:"column:description"`
	Price        float64 `gorm:"column:price"`
	Qty          int     `gorm:"column:qty"`
	IsActive     bool    `gorm:"column:is_active"`
	CategoryID   int64   `gorm:"column:category_id"`
	CategoryName string  `gorm:"column:category_name"`
}

type Category struct {
	ID   int64  `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
}

type ItemRepositoryInterface interface {
	List(filter domain.SearchItemList) ([]ItemsCategory, error)
}
