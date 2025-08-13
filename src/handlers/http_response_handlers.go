/*
 * CSO - Configuration Management API handlers
 */

package handlers

import (
	"net/http"
	"strings"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/munishchaudhary/book-store-app/src/generated/models"
)

type HandlerResponse struct {
	payload *models.APIResponse
}

// NewHandlerResponse creates HandlerResponse with the payload value
func NewHandlerResponse(payload *models.APIResponse) *HandlerResponse {
	return &HandlerResponse{payload}
}

// API response handler
func (hdlrResponse *HandlerResponse) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	apiResponse := hdlrResponse.payload
	if apiResponse.Code >= 200 && apiResponse.Code <= 299 {
		apiResponse.Success = true
	} else {
		apiResponse.Success = false
		// don't propagate invalid HTTP error codes and SQL errors
		if apiResponse.Code <= 0 || strings.Contains(apiResponse.Error, "sql:") {
			apiResponse.Code = http.StatusInternalServerError
			apiResponse.Error = "Internal Server Error"
		}
	}
	rw.WriteHeader(int(apiResponse.Code))
	if err := producer.Produce(rw, apiResponse); err != nil {
		panic(err) // let the recovery middleware deal with this: to be implemented
	}
}

func HttpError(code int64, message string) middleware.Responder {
	return NewHandlerResponse(&models.APIResponse{Code: code, Data: nil, Error: message, Success: false})
}

func HttpResponse(payload *models.APIResponse) middleware.Responder {
	return NewHandlerResponse(payload)
}

func HttpSuccess(code int64, data interface{}) middleware.Responder {
	return NewHandlerResponse(&models.APIResponse{Code: code, Data: data, Error: "", Success: true})
}

func StatusOK(data interface{}) middleware.Responder {
	return HttpSuccess(http.StatusOK, data)
}

func StatusCreated(data interface{}) middleware.Responder {
	return HttpSuccess(http.StatusCreated, data)
}

func StatusBadRequest(message string) middleware.Responder {
	return HttpError(http.StatusBadRequest, message)
}

func StatusNotFound(message string) middleware.Responder {
	return HttpError(http.StatusNotFound, message)
}

func StatusInternalServerError(message string) middleware.Responder {
	return HttpError(http.StatusInternalServerError, message)
}
