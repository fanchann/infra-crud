package repositories_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"infra-book-crud/models"
	"infra-book-crud/repositories"
)

var (
	mockRepo = repositories.MockBooksRepositories{Mock: mock.Mock{}}

	fakeData = []models.Books{
		{
			Title: "Book 1",
		},
		{
			Title: "Book 2",
		},
	}
)

func TestCreateBook(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo.Mock.On("AddNewBook", fakeData[0]).Return(fakeData[0])
		result := mockRepo.AddNewBook(fakeData[0])
		assert.Equal(t, fakeData[0], result)
	})
	t.Run("fail", func(t *testing.T) {
		mockRepo.Mock.On("AddNewBook", models.Books{}).Return(models.Books{})
		result := mockRepo.AddNewBook(models.Books{})
		assert.Equal(t, models.Books{}, result)
	})
}

func TestUpdateBook(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo.Mock.On("UpdateBook", fakeData[0]).Return(fakeData[0])
		result := mockRepo.UpdateBook(fakeData[0])
		assert.Equal(t, fakeData[0], result)
	})
	t.Run("fail", func(t *testing.T) {
		mockRepo.Mock.On("UpdateBook", models.Books{}).Return(models.Books{})
		result := mockRepo.UpdateBook(models.Books{})
		assert.Equal(t, models.Books{}, result)
	})
}

func TestDeleteBook(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo.Mock.On("DeleteBook", 1).Return(nil)
		result := mockRepo.DeleteBook(1)
		assert.Equal(t, nil, result)
	})
	t.Run("fail", func(t *testing.T) {
		exceptError := errors.New("book not found")
		mockRepo.Mock.On("DeleteBook", 2).Return(exceptError)
		result := mockRepo.DeleteBook(2)
		assert.Equal(t, exceptError, result)
	})
}

func TestGetAllBooks(t *testing.T) {
	mockRepo.Mock.On("GetAllBooks").Return(fakeData)
	result := mockRepo.GetAllBooks()
	assert.Equal(t, fakeData, result)

}

func TestGetBookById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo.Mock.On("GetBookById", 1).Return(fakeData[0], nil)
		result, err := mockRepo.GetBookById(1)
		assert.Equal(t, fakeData[0], result)
		assert.Equal(t, nil, err)
	})

	t.Run("fail", func(t *testing.T) {
		mockRepo.Mock.On("GetBookById", 2).Return(models.Books{}, errors.New("book not found"))
		result, err := mockRepo.GetBookById(2)
		assert.Equal(t, models.Books{}, result)
		assert.Equal(t, errors.New("book not found"), err)
	})
}
