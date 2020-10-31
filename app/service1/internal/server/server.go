package server

import (
	"database/sql"

	"github.com/dokyan1989/g1/app/service1/config"
	"github.com/dokyan1989/g1/app/service1/internal/store"
	"github.com/dokyan1989/g1/app/service1/pb"

	_ "github.com/go-sql-driver/mysql"
)

// Server ...
type Server struct {
	// GRPC
	pb.UnimplementedProductServiceServer
	pb.UnimplementedEmployeeServiceServer

	// Common
	store store.Store
}

// New ...
func New(cfg *config.Config) *Server {
	db, err := sql.Open("mysql", cfg.MySQL.DSN())
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	store := store.NewSQLStore(db)
	return &Server{store: store}
}
