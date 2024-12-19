package usecase_customer

import "net/http"

type UseCase interface {
	Register(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	VerifyOtp(w http.ResponseWriter, r *http.Request)
	DoLogin(w http.ResponseWriter, r *http.Request)
	AddLog(w http.ResponseWriter, r *http.Request)
	DoLogout(w http.ResponseWriter, r *http.Request)
	SendOTPResetPassword(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
}
