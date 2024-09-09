package domain

import (
	"context"
)

type ItemsInterface interface {
	List(ctx context.Context, filter SearchItemList) (interface{}, *ErrorResponse)
}
