package usecase_transaction

import "net/http"

type UseCase interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	GetTransDetail(w http.ResponseWriter, r *http.Request)
	GetTransactionSummary(w http.ResponseWriter, r *http.Request)
}
