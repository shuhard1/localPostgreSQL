package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	customer "github.com/shuhard1/localPostgreSQL/internal/customer"
	"github.com/shuhard1/localPostgreSQL/pkg/client/postgresql"
)

type repository struct {
	client postgresql.Client
}

// Create implements author.Repository
func (r *repository) Create(ctx context.Context, customer *customer.Customer) error {
	q := `INSERT INTO author (name) VALUES ($1) RETURNING id`

	err := r.client.QueryRow(ctx, q, customer.Info).Scan(&customer.ID)
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
func (r *repository) FindAll(ctx context.Context) (u []customer.Customer, err error) {
	q := `SELECT info ->> 'payment' AS payment FROM customers;`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	authors := make([]customer.Customer, 0)

	//проходится по рядам в таблице
	for rows.Next() {
		var ath customer.Customer

		//записывает в структуру данные из БД
		err = rows.Scan(&ath.Info)
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
func (r *repository) FindOne(ctx context.Context, id string) (customer.Customer, error) {
	q := `SELECT id, name FROM customers WHERE id = $1`

	var ath customer.Customer
	err := r.client.QueryRow(ctx, q, id).Scan(&ath.ID, &ath.Info)
	if err != nil {
		return customer.Customer{}, err
	}

	return ath, nil
}

// Update implements author.Repository
func (*repository) Update(ctx context.Context, user customer.Customer) error {
	panic("unimplemented")
}

func NewRepository(client postgresql.Client) customer.Repository {
	return &repository{
		client: client,
	}
}
