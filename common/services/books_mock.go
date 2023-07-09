package services

import (
	"github.com/stretchr/testify/mock"

	"infra-book-crud/common/types"
)

type MockBooksService struct {
	Mock mock.Mock
}

func (m *MockBooksService) AddNewBook(request types.Book) types.BookWithId {
	args := m.Mock.Called(request)
	return args.Get(0).(types.BookWithId)
}

func (m *MockBooksService) GetAllBooks() []types.BookWithId {
	args := m.Mock.Called()
	return args.Get(0).([]types.BookWithId)
}

func (m *MockBooksService) GetBookById(id int) (types.BookWithId, error) {
	args := m.Mock.Called(id)
	return args.Get(0).(types.BookWithId), args.Error(1)
}

func (m *MockBooksService) UpdateBook(request types.BookWithId) (types.BookWithId, error) {
	args := m.Mock.Called(request)
	return args.Get(0).(types.BookWithId), args.Error(1)
}

func (m *MockBooksService) DeleteBook(id int) error {
	args := m.Mock.Called(id)
	return args.Error(0)
}
