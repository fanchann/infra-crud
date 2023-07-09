package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"infra-book-crud/common/helpers"
	"infra-book-crud/common/services"
	"infra-book-crud/common/types"
)

var (
	service = bookService()
)

type CategoryController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type BookController struct{}

func (cntrl *BookController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	form := types.Book{}
	helpers.ReadFromRequestBody(request, &form)

	bookResponse := service.AddNewBook(form)
	if form.Title == "" || len(form.Title) <= 5 || len(form.Title) >= 200 {
		helpers.WriteToResponseBody(writer, helpers.ApiResponse(http.StatusBadRequest, "bad request", nil))
		return
	}

	helpers.WriteToResponseBody(writer, helpers.ApiResponse(http.StatusOK, "success create book", bookResponse))
}

func (cntrl *BookController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	/**
	get id
	**/
	form := new(types.BookWithId)
	id := params.ByName("id")
	idInt, _ := strconv.Atoi(id)
	form.Id = idInt

	bookResponse, errResponse := service.UpdateBook(*form)
	if errResponse != nil {
		helpers.WriteToResponseBody(writer, helpers.ApiResponse(http.StatusBadRequest, "bad request", errResponse))
		return
	}
	helpers.WriteToResponseBody(writer, helpers.ApiResponse(http.StatusOK, "success edit book by id", bookResponse))

}

func (cntrl *BookController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	form := new(types.BookWithId)
	id := params.ByName("id")
	idInt, _ := strconv.Atoi(id)
	form.Id = idInt

	errResponse := service.DeleteBook(idInt)
	if errResponse != nil {
		helpers.WriteToResponseBody(writer, helpers.ApiResponse(http.StatusNotFound, "not found", errResponse))
		return
	}
	helpers.WriteToResponseBody(writer, helpers.ApiResponse(http.StatusOK, "success delete data", nil))
}

func (cntrl *BookController) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	id := params.ByName("id")
	idInt, _ := strconv.Atoi(id)

	bookResponse, errResponse := service.GetBookById(idInt)
	if errResponse != nil {
		helpers.WriteToResponseBody(writer, helpers.ApiResponse(http.StatusNotFound, "nothing data", errResponse))
		return
	}
	helpers.WriteToResponseBody(writer, helpers.ApiResponse(http.StatusOK, "success get data by id", bookResponse))
}

func (cntrl *BookController) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	booksResponse := service.GetAllBooks()
	helpers.WriteToResponseBody(writer, helpers.ApiResponse(http.StatusOK, "success get all books", booksResponse))
}

func bookService() services.BookServiceImpl {
	return services.BookServiceImpl{}
}
