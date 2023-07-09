package services

import (
	"infra-book-crud/common/helpers"
	"infra-book-crud/common/types"
	"infra-book-crud/repositories"
)

type BookServiceImpl struct{}

var (
	bookRepo = repositories.BookRepositoryImpl{}
)

func (service *BookServiceImpl) AddNewBook(request types.Book) types.BookWithId {
	book := bookRepo.AddNewBook(helpers.TypesToModels(request))
	return helpers.ModelToTypes(book)
}

func (service *BookServiceImpl) GetAllBooks() []types.BookWithId {
	books := bookRepo.GetAllBooks()
	return helpers.ModelsToTypes(books)
}

func (service *BookServiceImpl) GetBookById(id int) (types.BookWithId, error) {
	book, err := bookRepo.GetBookById(id)
	if err != nil {
		return types.BookWithId{}, err
	}
	return helpers.ModelToTypes(book), nil
}

func (service *BookServiceImpl) UpdateBook(request types.BookWithId) (types.BookWithId, error) {
	book, err := bookRepo.GetBookById(request.Id)
	if err != nil {
		return types.BookWithId{}, err
	}
	if book.ID == 0 {
		return types.BookWithId{}, err
	}
	book = bookRepo.UpdateBook(helpers.TypesToModel(request))
	return helpers.ModelToTypes(book), nil
}

func (service *BookServiceImpl) DeleteBook(id int) error {
	book, err := bookRepo.GetBookById(id)
	if err != nil {
		return err
	}
	if book.ID == 0 {
		return err
	}

	bookRepo.DeleteBook(id)
	return nil
}
