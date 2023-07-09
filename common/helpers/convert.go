package helpers

import (
	"infra-book-crud/common/types"
	"infra-book-crud/models"
)

/*
*
service to controller
*
*/
func TypesToModels(data types.Book) models.Books {
	return models.Books{
		Title: data.Title,
	}
}

/*
*
convert models.Book{} to types.BookWithId{}
*
*/
func ModelToTypes(data models.Books) types.BookWithId {
	return types.BookWithId{
		Id: int(data.ID),
		Book: types.Book{
			Title: data.Title,
		},
	}
}

func ModelsToTypes(data []models.Books) []types.BookWithId {
	var result []types.BookWithId
	for _, v := range data {
		result = append(result, ModelToTypes(v))
	}
	return result
}

/*
*
convert types.BookWithId{} to models.Book{}
*
*/

func TypesToModel(data types.BookWithId) models.Books {
	return models.Books{
		Title: data.Book.Title,
	}
}
