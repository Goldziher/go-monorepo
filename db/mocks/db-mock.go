package mocks

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PgxRowMock struct{}

func (p PgxRowMock) Scan(dest ...any) error {
	fmt.Printf("Row Scan: %+v", dest)
	return nil
}

type PgxRowsMock struct{}

func (p PgxRowsMock) Close() {
	fmt.Println("Closing connection")
}

func (p PgxRowsMock) Err() error {
	return errors.New("error running rows transaction")
}

func (p PgxRowsMock) CommandTag() pgconn.CommandTag {
	return pgconn.CommandTag{}
}

func (p PgxRowsMock) FieldDescriptions() []pgconn.FieldDescription {
	return []pgconn.FieldDescription{}
}

func (p PgxRowsMock) Next() bool {
	return false
}

func (p PgxRowsMock) Scan(dest ...any) error {
	fmt.Printf("Rows Scan: %+v", dest)
	return nil
}

func (p PgxRowsMock) Values() ([]any, error) {
	return []any{}, nil
}

func (p PgxRowsMock) RawValues() [][]byte {
	return [][]byte{}
}

func (p PgxRowsMock) Conn() *pgx.Conn {
	return &pgx.Conn{}
}

type QueryMock struct{}

func (*QueryMock) QueryRow(_ context.Context, s string, i ...any) pgx.Row {
	fmt.Printf("Query: %s \nArgs: %+v", s, i)
	return PgxRowMock{}
}

func (*QueryMock) Exec(
	_ context.Context,
	s string,
	i ...any,
) (pgconn.CommandTag, error) {
	fmt.Printf("Query: %s \nArgs: %+v", s, i)
	return pgconn.CommandTag{}, nil
}

func (*QueryMock) Query(_ context.Context, s string, i ...any) (pgx.Rows, error) {
	fmt.Printf("Query: %s; \nArgs: %+v", s, i)
	return PgxRowsMock{}, nil
}
