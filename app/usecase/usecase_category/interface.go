package usecase_category

import "net/http"

type UseCase interface {
	FindById(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
}
