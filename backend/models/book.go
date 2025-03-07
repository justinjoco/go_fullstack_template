package models

import (
	"reflect"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	Id            uuid.UUID  `json:"id" gorm:"primary_key"`
	Author        string     `json:"author"`
	Title         string     `json:"title"`
	Genre         string     `json:"genre"`
	Description   *string    `json:"description"`
	Rating        *float32   `json:"rating"`
	DatePublished *time.Time `json:"date_published"`
}

func (Book) TableName() string {
	return "book"
}

func (book *Book) BeforeCreate(tx *gorm.DB) (err error) {
	if book.Id == uuid.Nil { // Check if the ID is not already set
		book.Id = uuid.New() // Generate a new UUID
	}
	return nil
}

// mergeStructs merges two structs, overwriting values in the first struct with those from the second
func (target *Book) Merge(source *Book) {
	targetVal := reflect.ValueOf(target).Elem()
	sourceVal := reflect.ValueOf(source).Elem()

	for i := 0; i < targetVal.NumField(); i++ {
		// Get the field name and value
		fieldName := targetVal.Type().Field(i).Name
		sourceField := sourceVal.FieldByName(fieldName)

		// If the source field is not zero (not empty), copy the value to the target
		if !sourceField.IsZero() {
			targetVal.Field(i).Set(sourceField)
		}
	}
}
