package usecase_user

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/jolebo/e-canteen-cashier-api/app/repository"
	"github.com/jolebo/e-canteen-cashier-api/config"
	"github.com/jolebo/e-canteen-cashier-api/entity"
	"github.com/jolebo/e-canteen-cashier-api/handler"
	"github.com/jolebo/e-canteen-cashier-api/pkg/exceptions"
	"github.com/jolebo/e-canteen-cashier-api/pkg/helpers"
	"golang.org/x/crypto/bcrypt"
)

type UseCaseImpl struct {
	UserRepository     repository.UserRepository
	TempCartRepository repository.TempCartRepository
	UserLogRepository  repository.UserLogRepository
	VersionRepository  repository.VersionRepository
	Validate           *validator.Validate
}

func NewUseCase(userRepo repository.UserRepository, userLogRepo repository.UserLogRepository, tempCartRepo repository.TempCartRepository, versionRepo repository.VersionRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:           validate,
		UserRepository:     userRepo,
		TempCartRepository: tempCartRepo,
		UserLogRepository:  userLogRepo,
		VersionRepository:  versionRepo,
	}
}

func (controller *UseCaseImpl) DoLogin(w http.ResponseWriter, r *http.Request) {
	dataLogin := entity.Login{}
	helpers.ReadFromRequestBody(r, &dataLogin)
	err := controller.Validate.Struct(dataLogin)
	helpers.PanicIfError(err)

	user := controller.UserRepository.FindSpesificData(r.Context(), entity.User{
		UserEmail: dataLogin.UserEmail,
	})
	if user == nil {
		panic(exceptions.NewUnAuthorizedError("Gagal melakukan login, periksa kembali email dan password apanic(exceptions.NewUnAuthorizedError("Gagal melakukan login, periksa kembali email dan password anda!"))
	}

	checkPassword := bcrypt.CompareHashAndPassword([]byte(user[0].UserPassword), []byte(dataLogin.UserPassword))

	if checkPassword != nil {
		panic(exceptions.NewUnAuthorizedError("Gagal melakukan login, periksa kembali email dan password anda!"))
	}

	if user[0].UserHasMobileAccess != 1 {
		panic(exceptions.NewUnAuthorizedError("Anda tidak diijinkan untuk mengakses aplikasi ini!"))
	}

	/* success login */
	/* remove tempcart on this user */
	controller.TempCartRepository.DeleteSpesificData(r.Context(), entity.TempCart{
		TempCartUserId: user[0].UserId,
	})
	decodedByte, _ := base64.StdEncoding.DecodeString(dataLogin.UserDeviceMetadata)
	/* insert token to table */
	controller.UserLogRepository.Create(r.Context(), entity.UserLog{
		LogUserUserId:   user[0].UserId,
		LogUserToken:    dataLogin.UserFcmToken,
		LogUserMetadata: string(decodedByte),
	})

	claims := &jwt.MapClaims{
		"UserId":           user[0].UserId,
		"UserEmail":        user[0].UserEmail,
		"UserNama":         user[0].UserName,
		"Pegawai":          user[0].Pegawai,
		"HasAccessCashier": user[0].UserHasMobileAccess,
		"exp":              time.Now().Add(1000 * time.Hour).Unix(),
		"iss":              config.GetEnv("APP_NAME"),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte(config.GetEnv("SECRET_KEY")))

	helpers.PanicIfError(err)

	webResponse := handler.WebResponse{
		Error:   false,
		Message: "Berhasil login",
		Data:    signedToken,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) DoLogout(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	helpers.ReadFromRequestBody(r, &data)

	/* get data user token */

	log := controller.UserLogRepository.FindSpesificData(r.Context(), entity.UserLog{
		LogUserUserId: data["user_id"].(string),
		LogUserToken:  data["fcmtoken"].(string),
	})

	if log != nil {
		controller.UserLogRepository.Update(r.Context(), entity.UserLog{
			LogUserId:         log[0].LogUserId,
			LogUserUserId:     log[0].LogUserUserId,
			LogUserToken:      log[0].LogUserToken,
			LogUserMetadata:   log[0].LogUserMetadata,
			LogUserLoginDate:  log[0].LogUserLoginDate,
			LogUserLogoutDate: helpers.CreateDateTime(),
		}, log[0].LogUserId)
	}

	webResponse := handler.WebResponse{
		Error:   false,
		Message: "Berhasil logout",
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) GetVersionAdmin(w http.ResponseWriter, r *http.Request) {
	data := controller.VersionRepository.GetVersionAdmin(r.Context())

	webResponse := handler.WebResponse{
		Error:   false,
		Message: "Berhasil mendapatkan data",
		Data:    data,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) GetVersionShop(w http.ResponseWriter, r *http.Request) {
	data := controller.VersionRepository.GetVersionShop(r.Context())

	webResponse := handler.WebResponse{
		Error:   false,
		Message: "Berhasil mendapatkan data",
		Data:    data,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) CheckMaintenanceMode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["confCode"]
	data := controller.UserRepository.CheckMaintenanceMode(r.Context(), map[string]interface{}{
		"conf_code": code,
	})

	webResponse := handler.WebResponse{
		Error:   false,
		Message: "Berhasil mendapatkan data",
		Data:    data,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
