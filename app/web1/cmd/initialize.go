package main

import (
	"github.com/dokyan1989/g1/app/web1/config"

	"github.com/dokyan1989/g1/app/web1/internal/models"
	"github.com/dokyan1989/g1/app/web1/internal/routes"
	"github.com/dokyan1989/g1/app/web1/internal/store"
	"github.com/dokyan1989/g1/lib/router"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func newRouter(cfg *config.Config) (*router.Router, error) {
	db, err := newDB(cfg.MySQL.DSN())
	if err != nil {
		return nil, err
	}

	err = autoMigration(db)
	if err != nil {
		return nil, err
	}

	store := store.NewStore(db)

	return routes.RegisterRoutes(store), nil
}

func autoMigration(db *gorm.DB) error {
	return db.AutoMigrate(&models.Product{})
}
