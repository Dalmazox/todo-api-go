package db

import (
	"context"
	"database/sql"

	"github.com/dalmazox/todo-api/configs"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Repository struct{}

var (
	builder    goqu.DialectWrapper
	connection *sql.DB
)

func init() {
	connectionString := configs.GetConfig().ConnectionString

	if connectionString == "" {
		panic("CONNECTION_STRING is not set")
	}

	builder = goqu.Dialect("postgres")
	pgsqlConn, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	err = pgsqlConn.Ping()
	if err != nil {
		panic(err)
	}

	connection = pgsqlConn
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreateTodo(ctx context.Context, description string, done bool) error {
	record := goqu.Record{
		"id":          uuid.NewString(),
		"description": description,
		"done":        done,
	}
	sql, _, err := builder.Insert("todos").Rows(record).ToSQL()
	if err != nil {
		return err
	}

	_, err = connection.ExecContext(ctx, sql)
	if err != nil {
		return err
	}

	return nil
}
