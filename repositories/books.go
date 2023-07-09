package repositories

import "infra-book-crud/models"

type IBookRepository interface {
	AddNewBook(book models.Books) models.Books
	GetAllBooks() []models.Books
	GetBookById(id int) (models.Books, error)
	UpdateBook(book models.Books) models.Books
	DeleteBook(id int) error
}
