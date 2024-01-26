package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

type RepositoryParams struct {
	Url string
}

func NewRepository(p RepositoryParams) *Repository {
	db, err := gorm.Open(postgres.Open(p.Url), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Repository{
		Db: db,
	}
}
