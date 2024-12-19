package usecase_order

import (
	"net/http"
)

type UseCase interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	GetOrderDetail(w http.ResponseWriter, r *http.Request)
	OrderCanceled(w http.ResponseWriter, r *http.Request)
	OrderProcessed(w http.ResponseWriter, r *http.Request)
}
