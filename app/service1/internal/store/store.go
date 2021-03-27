package store

import (
	"context"
	"database/sql"
)

type Store interface {
	// Employee
	ListEmployees(ctx context.Context, params ListEmployeesParams) ([]Employee, error)
	CreateEmployee(ctx context.Context, params CreateEmployeeParams) (uint64, error)
	UpdateEmployee(ctx context.Context, params UpdateEmployeeParams) (uint64, error)
	DeleteEmployee(ctx context.Context, empNo uint64) (uint64, error)
}

type SQLStore struct {
	db *sql.DB
}

func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{db}
}
