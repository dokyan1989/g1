package store

import (
	"context"
	"fmt"
	"log"
	"time"

	sqlUtils "github.com/dokyan1989/g1/lib/utils/sql"
)

type Employee struct {
	EmpNo     uint64    `json:"emp_no"`
	BirthDate time.Time `json:"birth_date"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Gender    string    `json:"gender"`
	HireDate  time.Time `json:"hire_date"`
}

type ListEmployeesParams struct {
	EmpNos        []uint64  `json:"emp_nos"`
	FromBirthDate time.Time `json:"from_birth_date"`
	ToBirthDate   time.Time `json:"to_birth_date"`
	Names         []string  `json:"names"`
	Gender        string    `json:"gender"`
	FromHireDate  time.Time `json:"from_hire_date"`
	ToHireDate    time.Time `json:"to_hire_date"`
	Limit         uint32    `json:"limit"`
	Offset        uint64    `json:"offset"`
}

func (s *SQLStore) ListEmployees(ctx context.Context, params ListEmployeesParams) ([]Employee, error) {
	sql, vals, err := s.buildListEmployeesSql(params)
	if err != nil {
		return nil, err
	}

	var employees []Employee
	rows, err := s.db.QueryContext(ctx, sql, vals...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e Employee
		err := rows.Scan(&e.EmpNo, &e.BirthDate, &e.FirstName, &e.LastName, &e.Gender, &e.HireDate)
		if err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (s *SQLStore) buildListEmployeesSql(params ListEmployeesParams) (string, []interface{}, error) {
	sql := `select emp_no, birth_date, first_name, last_name, gender, hire_date
	from employees where 1=1`
	vals := []interface{}{}

	if len(params.EmpNos) > 0 {
		expr, args, err := sqlUtils.BuildSqlIn("emp_no", params.EmpNos)
		if err != nil {
			return "", []interface{}{}, err
		}
		sql = sqlUtils.AppendClause(sql, fmt.Sprintf("and %s", expr))
		vals = append(vals, args...)
	}

	if len(params.Names) > 0 {
		expr1, args1, err := sqlUtils.BuildSqlIn("first_name", params.Names)
		if err != nil {
			return "", []interface{}{}, err
		}
		expr2, args2, err := sqlUtils.BuildSqlIn("last_name", params.Names)
		if err != nil {
			return "", []interface{}{}, err
		}

		cls := fmt.Sprintf("(%s or %s)", expr1, expr2)
		sql = sqlUtils.AppendClause(sql, fmt.Sprintf("and %s", cls))
		vals = append(vals, args1...)
		vals = append(vals, args2...)
	}

	if !params.FromBirthDate.IsZero() {
		sql = sqlUtils.AppendClause(sql, "and birth_date >= ?")
		vals = append(vals, params.FromBirthDate.Format(time.RFC3339))
	}

	if !params.ToBirthDate.IsZero() {
		sql = sqlUtils.AppendClause(sql, "and birth_date <= ?")
		vals = append(vals, params.ToBirthDate.Format(time.RFC3339))
	}

	if params.Gender != "" {
		sql = sqlUtils.AppendClause(sql, "and gender = ?")
		vals = append(vals, params.Gender)
	}

	if !params.FromHireDate.IsZero() {
		sql = sqlUtils.AppendClause(sql, "and hire_date >= ?")
		vals = append(vals, params.FromHireDate.Format(time.RFC3339))
	}

	if !params.ToHireDate.IsZero() {
		sql = sqlUtils.AppendClause(sql, "and hire_date <= ?")
		vals = append(vals, params.ToHireDate.Format(time.RFC3339))
	}

	cls, limit, offset := sqlUtils.BuildSqlLimitOffset(params.Limit, params.Offset)
	sql = sqlUtils.AppendClause(sql, cls)
	vals = append(vals, limit, offset)

	log.Println(sql)
	log.Println(vals)

	return sql, vals, nil
}

type CreateEmployeeParams struct {
	BirthDate time.Time `json:"birth_date"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Gender    string    `json:"gender"`
	HireDate  time.Time `json:"hire_date"`
}

func (s *SQLStore) CreateEmployee(ctx context.Context, params CreateEmployeeParams) (uint64, error) {
	sql := `insert into employees(birth_date, first_name, last_name, gender, hire_date) values(?,?,?,?,?)`
	stmt, err := s.db.PrepareContext(ctx, sql)
	if err != nil {
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, params.BirthDate, params.FirstName, params.LastName, params.Gender, params.HireDate)
	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}

type UpdateEmployeeParams struct {
	EmpNo     uint64    `json:"emp_no"`
	BirthDate time.Time `json:"birth_date"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Gender    string    `json:"gender"`
	HireDate  time.Time `json:"hire_date"`
}

func (s *SQLStore) UpdateEmployee(ctx context.Context, params UpdateEmployeeParams) (uint64, error) {
	sql := `update employees 
	set birth_date = ?, first_name = ?, last_name = ?, gender = ?, hire_date = ?
	where emp_no = ?`
	stmt, err := s.db.PrepareContext(ctx, sql)
	if err != nil {
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, params.BirthDate, params.FirstName, params.LastName, params.Gender, params.HireDate, params.EmpNo)
	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}

func (s *SQLStore) DeleteEmployee(ctx context.Context, empNo uint64) (uint64, error) {
	sql := `update from employees where emp_no = ?`
	stmt, err := s.db.PrepareContext(ctx, sql)
	if err != nil {
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, empNo)
	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}
