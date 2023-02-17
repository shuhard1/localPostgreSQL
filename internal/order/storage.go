package order

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, id string, info string) error
	FindAll(ctx context.Context) (u []Order, err error)
	FindOne(ctx context.Context, id string) (Order, error)
}
