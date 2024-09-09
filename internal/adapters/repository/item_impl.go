package repository

import (
	"github.com/dodirepository/product-svc/infrastructure/database"
	domain "github.com/dodirepository/product-svc/internal/domain/repository"
	domainUsecase "github.com/dodirepository/product-svc/internal/domain/usecases"

	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

// AccountRepositoryHandler :nodoc:
func ItemRepositoryHandler() domain.ItemRepositoryInterface {
	return &ItemRepository{
		db: database.GetConnection(),
	}
}

func (i *ItemRepository) List(filter domainUsecase.SearchItemList) ([]domain.ItemsCategory, error) {
	items := []domain.ItemsCategory{}
	query := i.db.Table("items as i").Select("i.id, i.name, i.description, i.price, i.qty, c.id as category_id, c.name as category_name")
	query = query.Joins("inner join category as c ON i.category_id = c.id and c.is_deleted = false and c.is_active = true")
	query = query.Where("i.is_active = ? AND i.is_deleted = false", true).
		Scan(&items)
	if query.RecordNotFound() {
		return nil, nil
	}
	return items, nil
}
