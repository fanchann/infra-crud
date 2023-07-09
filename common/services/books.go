package services

import "infra-book-crud/common/types"

type IBookService interface {
	AddNewBook(request types.Book) types.BookWithId
	GetAllBooks() []types.BookWithId
	GetBookById(id int) (types.BookWithId, error)
	UpdateBook(request types.BookWithId) (types.BookWithId, error)
	DeleteBook(id int) error
}
