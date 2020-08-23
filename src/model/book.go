package model

import (
	"ufc.com/deti/go-dad/src/database"

	"github.com/jinzhu/gorm"
)

// Book model
type Book struct {
	gorm.Model
	Name    string  `json:"nome"`
	Authors string  `json:"autores"`
	Year    string  `json:"data_lancamento"`
	Preco   float64 `json:"preco"`
	Cover   string  `json:"nome_capa"`
}

type Books []Book

// Migrate -- Migrate
func Migrate() {

	db, _ := database.NewConnection()
	db.AutoMigrate(&Book{})
	db.Close()
}

func GetAll() Books {

	var books Books
	db, _ := database.NewConnection()
	defer db.Close()
	db.Find(&books)

	return books
}

func GetOne(id int) (*Book, error) {

	db, _ := database.NewConnection()
	defer db.Close()
	var book Book
	err := db.Where("id = ?", id).Find(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil

}

func Store(book Book) *Book {

	db, _ := database.NewConnection()
	defer db.Close()
	db.Create(&book)
	return &book

}

func Delete(id int) error {

	db, _ := database.NewConnection()
	defer db.Close()
	var book Book
	err := db.Where("id = ?", id).Find(&book).Error
	db.Delete(&book)
	if err != nil {
		
	}
	return nil
}
