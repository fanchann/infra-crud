package repositories

import (
	"github.com/stretchr/testify/mock"

	"infra-book-crud/models"
)

type MockBooksRepositories struct {
	Mock mock.Mock
}

func (m *MockBooksRepositories) AddNewBook(book models.Books) models.Books {
	args := m.Mock.Called(book)
	return args.Get(0).(models.Books)
}

func (m *MockBooksRepositories) GetAllBooks() []models.Books {
	args := m.Mock.Called()
	return args.Get(0).([]models.Books)
}

func (m *MockBooksRepositories) GetBookById(id int) (models.Books, error) {
	args := m.Mock.Called(id)
	return args.Get(0).(models.Books), args.Error(1)
}

func (m *MockBooksRepositories) UpdateBook(book models.Books) models.Books {
	args := m.Mock.Called(book)
	return args.Get(0).(models.Books)
}

func (m *MockBooksRepositories) DeleteBook(id int) error {
	args := m.Mock.Called(id)
	return args.Error(0)
}
