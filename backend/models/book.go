package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	id             uuid.UUID
	author         string
	title          string
	genre          string
	description    *string
	rating         *int
	date_published *time.Time
}
