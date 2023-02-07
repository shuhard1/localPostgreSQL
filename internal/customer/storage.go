package customer

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, customer *Customer) error
	FindAll(ctx context.Context) (u []Customer, err error)
	FindOne(ctx context.Context, id string) (Customer, error)
	Update(ctx context.Context, user Customer) error
	Delete(ctx context.Context, id string) error
}
