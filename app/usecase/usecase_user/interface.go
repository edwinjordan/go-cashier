package usecase_user

import "net/http"

type UseCase interface {
	DoLogin(w http.ResponseWriter, r *http.Request)
	DoLogout(w http.ResponseWriter, r *http.Request)
	GetVersionAdmin(w http.ResponseWriter, r *http.Request)
	GetVersionShop(w http.ResponseWriter, r *http.Request)
	CheckMaintenanceMode(w http.ResponseWriter, r *http.Request)
}
