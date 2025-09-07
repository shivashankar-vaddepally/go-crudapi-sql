package models

import (
	"github.com/jinzhu/gorm"
	app "github.com/vaddepally-shiva-shankar/go-crudapi-sql/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model         // gorm.Model adds fields ID, CreatedAt, UpdatedAt, DeletedAt automatically
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	app.Connect()
	db = app.GetDB()
	db.AutoMigrate(&Book{})
}
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func (b *Book) GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func (b *Book) GetBookbyId(Id int64) *Book {
	var book Book
	db.Find(&book, Id)
	return &book
}

func (b *Book) DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}
