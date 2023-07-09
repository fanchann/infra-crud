package routers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"infra-book-crud/common/controller"
)

func InitRouters() http.Handler {
	router := httprouter.New()

	controller := controller.BookController{}

	router.GET("/api/v1/books", controller.FindAll)
	router.GET("/api/v1/book/:id", controller.FindById)
	router.POST("/api/v1/book/new", controller.Create)
	router.PUT("/api/v1/book/:id", controller.Update)
	router.DELETE("/api/v1/book/:id", controller.Delete)

	return router
}
