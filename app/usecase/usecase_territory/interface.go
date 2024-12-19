package usecase_territory

import "net/http"

type UseCase interface {
	GetProvince(w http.ResponseWriter, r *http.Request)
	GetCity(w http.ResponseWriter, r *http.Request)
	GetSubdistrict(w http.ResponseWriter, r *http.Request)
	GetVillage(w http.ResponseWriter, r *http.Request)
}
