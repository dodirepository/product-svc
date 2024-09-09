package usecases

import (
	"context"
	"net/http"

	"github.com/dodirepository/product-svc/internal/adapters/repository"
	domainRepo "github.com/dodirepository/product-svc/internal/domain/repository"
	domain "github.com/dodirepository/product-svc/internal/domain/usecases"
	"github.com/sirupsen/logrus"
)

type ItemUsecase struct {
	itemRepository domainRepo.ItemRepositoryInterface
}

func ItemUsecaseHandler() domain.ItemsInterface {
	return &ItemUsecase{
		itemRepository: repository.ItemRepositoryHandler(),
	}
}

func (u *ItemUsecase) List(ctx context.Context, filter domain.SearchItemList) (interface{}, *domain.ErrorResponse) {

	items, err := u.itemRepository.List(filter)
	if err != nil {
		logrus.WithError(err).Error("failed to get items")
		return nil, &domain.ErrorResponse{
			Message: http.StatusText(http.StatusInternalServerError),
			Status:  http.StatusInternalServerError,
		}
	}
	result := []domain.ListProducts{}
	for _, v := range items {
		result = append(result, domain.ListProducts{
			ID:         v.ID,
			CategoryID: v.CategoryID,
			Name:       v.Name,
			Desc:       v.Desc,
			Price:      v.Price,
			Qty:        v.Qty,
			IsActive:   v.IsActive,
			Category: domain.Category{
				ID:   v.CategoryID,
				Name: v.CategoryName,
			},
		})
	}
	return result, nil

}
