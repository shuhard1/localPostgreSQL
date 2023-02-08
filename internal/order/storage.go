package order

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, customer *Order) error
	FindAll(ctx context.Context) (u []Order, err error)
	FindOne(ctx context.Context, id string) (Order, error)
	Update(ctx context.Context, user Order) error
	Delete(ctx context.Context, id string) error
}
