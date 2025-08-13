package handlers

import (
	"github.com/munishchaudhary/book-store-app/src/generated/restapi/operations"
)

func RegisterAPIHandlers(api *operations.BookStoreAPI) {
	api.CreateBookHandler = operations.CreateBookHandlerFunc(CreateBookHandler)
	api.GetBookHandler = operations.GetBookHandlerFunc(GetBookHandler)
	api.GetBooksHandler = operations.GetBooksHandlerFunc(GetBooksHandler)
	api.UpdateBookHandler = operations.UpdateBookHandlerFunc(UpdateBookHandler)
	api.DeleteBookHandler = operations.DeleteBookHandlerFunc(DeleteBookHandler)
	api.HealthCheckHandler = operations.HealthCheckHandlerFunc(HealthCheckHandler)
}
