package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (book *Book) CreateBook() error {
	fmt.Printf("%+v\n", book)
	return nil
}
