package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/shuhard1/localPostgreSQL/internal/author"
	"github.com/shuhard1/localPostgreSQL/pkg/client/postgresql"
)

type repository struct {
	client postgresql.Client
}

// Create implements author.Repository
func (r *repository) Create(ctx context.Context, author *author.Author) error {
	q := `INSERT INTO author (name) VALUES ($1) RETURNING id`

	err := r.client.QueryRow(ctx, q, author.Name).Scan(&author.ID)
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
func (r *repository) FindAll(ctx context.Context) (u []author.Author, err error) {
	q := `SELECT info ->> 'payment' AS payment FROM customers;`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	authors := make([]author.Author, 0)

	//проходится по рядам в таблице
	for rows.Next() {
		var ath author.Author

		//записывает в структуру данные из БД
		err = rows.Scan(&ath.Name)
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
func (r *repository) FindOne(ctx context.Context, id string) (author.Author, error) {
	q := `SELECT id, name FROM customers WHERE id = $1`

	var ath author.Author
	err := r.client.QueryRow(ctx, q, id).Scan(&ath.ID, &ath.Name)
	if err != nil {
		return author.Author{}, err
	}

	return ath, nil
}

// Update implements author.Repository
func (*repository) Update(ctx context.Context, user author.Author) error {
	panic("unimplemented")
}

func NewRepository(client postgresql.Client) author.Repository {
	return &repository{
		client: client,
	}
}
