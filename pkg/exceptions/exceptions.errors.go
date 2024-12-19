package exceptions

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/jolebo/e-canteen-cashier-api/handler"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"github.com/jolebo/e-canteen-cashier-api/pkg/validations"
)

func ErrorHadler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}

	if badRequestError(w, r, err) {
		return
	}

	if unAuthorizedError(w, r, err) {
		return
	}

	if validationError(w, r, err) {
		return
	}

	if conflictError(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func conflictError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(ConflictError)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)

		webResponse := handler.WebResponse{
			Error:   true,
			Message: exception.Error,
		}
		helpers.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)

		webResponse := handler.WebResponse{
			Error:   true,
			Message: exception.Error,
		}
		helpers.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func badRequestError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(BadRequestError)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		webResponse := handler.WebResponse{
			Error:   true,
			Message: exception.Error,
		}
		helpers.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	webResponse := handler.WebResponse{
		Error:   true,
		Message: "Internal Server Error",
		// Data:    err,
	}
	helpers.WriteToResponseBody(w, webResponse)
	return true
}

func unAuthorizedError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(UnAuthorizedError)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		webResponse := handler.WebResponse{
			Error:   true,
			Message: exception.Error,
		}
		helpers.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func validationError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := handler.WebResponse{
			Error:   true,
			Message: "Data yang dikirim belum sesuai",
			Data:    validations.GetValidationMessage(exception),
		}
		helpers.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}
