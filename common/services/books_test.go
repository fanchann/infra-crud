package services_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"infra-book-crud/common/services"
	"infra-book-crud/common/types"
)

var (
	mockServices  = services.MockBooksService{Mock: mock.Mock{}}
	fakeTitleBook = []types.Book{
		{
			Title: "Book 1",
		},
		{
			Title: "Book 2",
		},
	}

	fakeTitleBookWithId = []types.BookWithId{
		{
			Id:   1,
			Book: fakeTitleBook[0],
		},
		{
			Id:   2,
			Book: fakeTitleBook[1],
		},
	}
)

func TestCreateNewBook(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockServices.Mock.On("AddNewBook", fakeTitleBook[0]).Return(fakeTitleBookWithId[0])
		result := mockServices.AddNewBook(fakeTitleBook[0])
		assert.Equal(t, fakeTitleBookWithId[0], result)
	})

	t.Run("fail", func(t *testing.T) {
		mockServices.Mock.On("AddNewBook", types.Book{}).Return(types.BookWithId{})
		result := mockServices.AddNewBook(types.Book{})
		assert.Equal(t, types.BookWithId{}, result)
	})
}

func TestGetBookById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockServices.Mock.On("GetBookById", 1).Return(fakeTitleBookWithId[0], nil)
		result, err := mockServices.GetBookById(1)
		assert.Equal(t, fakeTitleBookWithId[0], result)
		assert.Nil(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		mockServices.Mock.On("GetBookById", 2).Return(types.BookWithId{}, errors.New("book not found"))
		result, err := mockServices.GetBookById(2)
		assert.Equal(t, types.BookWithId{}, result)
		assert.Equal(t, errors.New("book not found"), err)
	})
}

func TestGetAllBooks(t *testing.T) {
	mockServices.Mock.On("GetAllBooks").Return(fakeTitleBookWithId)
	result := mockServices.GetAllBooks()
	assert.Equal(t, fakeTitleBookWithId, result)
}

func TestUpdateBook(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockServices.Mock.On("UpdateBook", fakeTitleBookWithId[0]).Return(fakeTitleBookWithId[0], nil)
		result, err := mockServices.UpdateBook(fakeTitleBookWithId[0])
		assert.Equal(t, fakeTitleBookWithId[0], result)
		assert.Nil(t, err)
	})
	t.Run("fail", func(t *testing.T) {
		mockServices.Mock.On("UpdateBook", types.BookWithId{}).Return(types.BookWithId{}, errors.New("book not found"))
		result, err := mockServices.UpdateBook(types.BookWithId{})
		assert.Equal(t, types.BookWithId{}, result)
		assert.Equal(t, errors.New("book not found"), err)
	})
}

func TestDeleteBook(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockServices.Mock.On("DeleteBook", 1).Return(nil)
		result := mockServices.DeleteBook(1)
		assert.Nil(t, result)
	})
	t.Run("fail", func(t *testing.T) {
		mockServices.Mock.On("DeleteBook", 2).Return(errors.New("book not found"))
		result := mockServices.DeleteBook(2)
		assert.Equal(t, errors.New("book not found"), result)
	})
}
