package repository

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type BookRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewBookRepository(db *gorm.DB, rdb *redis.Client) *BookRepository {
	return &BookRepository{db: db, rdb: rdb}
}
