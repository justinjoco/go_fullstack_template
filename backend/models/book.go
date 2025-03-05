package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	Id            uuid.UUID  `json:"id" gorm:"primary_key"`
	Author        string     `json:"author"`
	Title         string     `json:"title"`
	Genre         string     `json:"genre"`
	Description   *string    `json:"description"`
	Rating        *int       `json:"rating"`
	DatePublished *time.Time `json:"date_published"`
}
