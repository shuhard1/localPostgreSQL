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

// Create implements author.Repository
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

// Delete implements author.Repository
func (*repository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindAll implements author.Repository
func (r *repository) FindAll(ctx context.Context) (u []order.Order, err error) {
	q := `SELECT id, info FROM orders;`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	authors := make([]order.Order, 0)

	//проходится по рядам в таблице
	for rows.Next() {
		var ath order.Order

		//записывает в структуру данные из БД
		err = rows.Scan(&ath.ID, &ath.Info)
		if err != nil {
			return nil, err
		}

		authors = append(authors, ath)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

// FindOne implements author.Repository
func (r *repository) FindOne(ctx context.Context, id string) (order.Order, error) {

	q := `SELECT id, info FROM orders WHERE id = $1`

	var ath order.Order
	err := r.client.QueryRow(ctx, q, id).Scan(&ath.ID, &ath.Info)
	if err != nil {
		return order.Order{}, err
	}

	return ath, nil
}

// Update implements author.Repository
func (*repository) Update(ctx context.Context, user order.Order) error {
	panic("unimplemented")
}

func NewRepository(client postgresql.Client) order.Repository {
	return &repository{
		client: client,
	}
}
