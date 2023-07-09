package repositories

import (
	"log"
	"os"

	"github.com/zakirkun/infra-go/pkg/database"
	"gorm.io/gorm"

	"infra-book-crud/models"
)

type BookRepositoryImpl struct{}

func (repo *BookRepositoryImpl) AddNewBook(book models.Books) models.Books {
	db := openDB()
	if db.Create(&book).Error != nil {
		log.Printf("Error: %v", db.Create(&book).Error)
		return models.Books{}
	}
	return book

}

func (repo *BookRepositoryImpl) GetAllBooks() []models.Books {
	var books []models.Books
	db := openDB()
	if db.Find(&books).Error != nil {
		log.Printf("Error: %v", db.Find(&books).Error)
		os.Exit(1)
	}
	return books
}

func (repo *BookRepositoryImpl) GetBookById(id int) (models.Books, error) {
	var book models.Books
	db := openDB()
	if db.First(&book, id).Error != nil {
		log.Printf("Error: %v", db.First(&book, id).Error)
		return models.Books{}, db.First(&book, id).Error
	}
	return book, nil
}

func (repo *BookRepositoryImpl) UpdateBook(book models.Books) models.Books {
	db := openDB()
	if db.Save(&book).Error != nil {
		log.Printf("Error: %v", db.Save(&book).Error)
		return models.Books{}
	}
	return book
}

func (repo *BookRepositoryImpl) DeleteBook(id int) error {
	db := openDB()
	if db.Delete(&models.Books{}, id).Error != nil {
		log.Printf("Error: %v", db.Delete(&models.Books{}, id).Error)
		return db.Delete(&models.Books{}, id).Error
	}
	return nil
}

func openDB() *gorm.DB {
	db, err := database.DB.OpenDB()
	if err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}

	return db
}
