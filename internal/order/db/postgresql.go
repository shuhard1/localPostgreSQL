package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/shuhard1/localPostgreSQL/internal/order"
	"github.com/shuhard1/localPostgreSQL/pkg/client/postgresql"
)

type repository struct {
	client postgresql.Client
}

func (r *repository) Create(ctx context.Context, id string, info string) error {
	q := `INSERT INTO orders (id, info) VALUES ($1, $2) RETURNING id`

	err := r.client.QueryRow(ctx, q, id, info).Scan()
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			return newErr
		}
		return err
	}
	return nil
}

func (r *repository) FindAll(ctx context.Context) (u []order.Order, err error) {
	q := `SELECT id, info FROM orders;`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	orders := make([]order.Order, 0)

	for rows.Next() {
		var ath order.Order

		err = rows.Scan(&ath.ID, &ath.Info)
		if err != nil {
			return nil, err
		}

		orders = append(orders, ath)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (order.Order, error) {

	q := `SELECT id, info FROM orders WHERE id = $1`

	var ath order.Order
	err := r.client.QueryRow(ctx, q, id).Scan(&ath.ID, &ath.Info)
	if err != nil {
		return order.Order{}, err
	}

	return ath, nil
}

func NewRepository(client postgresql.Client) order.Repository {
	return &repository{
		client: client,
	}
}
