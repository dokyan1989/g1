package store

import (
	"context"
	"fmt"
	"log"

	sqlUtils "github.com/dokyan1989/g1/lib/utils/sql"
)

type Product struct {
	ID    uint64  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ListProductsParams struct {
	IDs    []uint64 `json:"ids"`
	Names  []string `json:"names"`
	Limit  uint32   `json:"limit"`
	Offset uint64   `json:"offset"`
}

func (s *SQLStore) ListProducts(ctx context.Context, params ListProductsParams) ([]Product, error) {
	sql, vals, err := s.makeListProductsSql(params)
	if err != nil {
		return nil, err
	}

	var products []Product
	rows, err := s.db.QueryContext(ctx, sql, vals...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *SQLStore) makeListProductsSql(params ListProductsParams) (string, []interface{}, error) {
	sql := "select id, name, price from products where 1 = 1"
	vals := []interface{}{}

	if len(params.IDs) > 0 {
		expr, args, err := sqlUtils.MakeSqlIn("id", params.IDs)
		if err != nil {
			return "", []interface{}{}, err
		}
		sql += fmt.Sprintf(" and %s", expr)
		vals = append(vals, args...)
	}

	if len(params.Names) > 0 {
		expr, args, err := sqlUtils.MakeSqlIn("name", params.Names)
		if err != nil {
			return "", []interface{}{}, err
		}
		sql += fmt.Sprintf(" and %s", expr)
		vals = append(vals, args...)
	}

	sql += fmt.Sprintf(" %s", sqlUtils.MakeSqlLimitOffset(params.Limit, params.Offset))

	log.Println(sql)
	log.Println(vals)

	return sql, vals, nil
}

type CreateProductParams struct {
	Name  string
	Price float64
}

func (s *SQLStore) CreateProduct(ctx context.Context, params CreateProductParams) (uint64, error) {
	stmt, err := s.db.PrepareContext(ctx, "insert into products(name, price) values(?,?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.ExecContext(ctx, params.Name, params.Price)
	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}
