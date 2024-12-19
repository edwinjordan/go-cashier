package usecase_customer

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"html"
	"math/rand"
	"net/http"
	"strconv"
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
	CustomerRepository repository.CustomerRepository
	UserOTPRepository  repository.UserOTPRepository
	TempCartRepository repository.TempCartRepository
	UserLogRepository  repository.UserLogRepository
	Validate           *validator.Validate
}

func NewUseCase(customerRepo repository.CustomerRepository, otpRepo repository.UserOTPRepository, tempCartRepo repository.TempCartRepository, userLogRepo repository.UserLogRepository, validate *validator.Validate) UseCase {
	return &UseCaseImpl{
		Validate:           validate,
		CustomerRepository: customerRepo,
		UserOTPRepository:  otpRepo,
		TempCartRepository: tempCartRepo,
		UserLogRepository:  userLogRepo,
	}
}

func (controller *UseCaseImpl) Register(w http.ResponseWriter, r *http.Request) {
	dataRequest := entity.Customer{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)

	/* check if customer phonenumber exist */
	dataPhone := controller.CustomerRepository.FindSpesificData(r.Context(), entity.Customer{
		CustomerPhonenumber: dataRequest.CustomerPhonenumber,
	})

	if dataPhone != nil {
		panic(exceptions.NewConflictError("Nomor hp sudah digunakan, silahkan gunakan nomor hp lain atau masuk menggunakan nomor hp terdaftar"))
	}

	/* check if customer email exist */

	dataEmail := controller.CustomerRepository.FindSpesificData(r.Context(), entity.Customer{
		CustomerEmail: dataRequest.CustomerEmail,
	})

	if dataEmail != nil {
		panic(exceptions.NewConflictError("Email sudah digunakan, silahkan gunakan email lain"))
	}

	dataRequest.CustomerPassword = helpers.EncryptPassword(dataRequest.CustomerPassword)
	dataRequest.CustomerCode = controller.CustomerRepository.GenCustCode(r.Context())
	dataRequest.CustomerName = html.EscapeString(dataRequest.CustomerName)
	dataRequest.CustomerEmail = html.EscapeString(dataRequest.CustomerEmail)
	dataRequest.CustomerPhonenumber = html.EscapeString(dataRequest.CustomerPhonenumber)
	customer := controller.CustomerRepository.Create(r.Context(), dataRequest)
	/* otp */
	rand.Seed(time.Now().UTC().UnixNano())
	rdm := fmt.Sprint(rand.Int())

	otp := controller.UserOTPRepository.Create(r.Context(), entity.UserOTP{
		OtpCustomerId: customer.CustomerId,
		OtpNumber:     rdm[:6],
	})
	/* sent otp */
	helpers.SendWhatsapp(r.Context(), map[string]interface{}{
		"phonenumber": customer.CustomerPhonenumber,
		"text":        "eCanteen\n\nKode Verifikasi Anda : *" + otp.OtpNumber + "*\nakan berlaku selama 15 menit .\n\nPENTING !!! DEMI KEAMANAN AKUN ANDA, JANGAN BERIKAN KODE RAHASIA INI KEPADA SIAPAPUN,  Terima Kasih.",
	})

	dataResponse := map[string]interface{}{
		"otp":      otp,
		"customer": customer,
	}
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessCreateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customerId"]
	dataRequest := map[string]interface{}{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	customer, err := controller.CustomerRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	password := customer.CustomerPassword

	/* check jika pengguna ingin mengubah kata sandinya */
	if dataRequest["customer_new_password"].(string) != "" {
		/* check kata sandi lamanya */
		checkPassword := bcrypt.CompareHashAndPassword([]byte(customer.CustomerPassword), []byte(dataRequest["customer_old_password"].(string)))
		if checkPassword != nil {
			panic(exceptions.NewBadRequestError("Tidak dapat mengubah kata sandi karena kata sandi lama tidak cocok"))
		}

		password = helpers.EncryptPassword(dataRequest["customer_new_password"].(string))
	}
	dataCustomer := entity.Customer{
		CustomerName:        html.EscapeString(dataRequest["customer_name"].(string)),
		CustomerGender:      dataRequest["customer_gender"].(string),
		CustomerPhonenumber: customer.CustomerPhonenumber,
		CustomerEmail:       customer.CustomerEmail,
		CustomerDob: sql.NullString{
			String: dataRequest["customer_dob"].(string),
			Valid:  true,
		},
		CustomerPassword: password,
		CustomerClass:    dataRequest["customer_class"].(string),
		CustomerMajorId:  dataRequest["customer_major_id"].(string),
		CustomerUpdateAt: helpers.CreateDateTime(),
	}
	dataResponse := controller.CustomerRepository.Update(r.Context(), []string{"customer_name", "customer_gender", "customer_phonenumber", "customer_email", "customer_dob", "customer_password", "customer_class", "customer_major_id", "customer_update_at"}, dataCustomer, id)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessUpdateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customerId"]
	_, err := controller.CustomerRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	controller.CustomerRepository.Delete(r.Context(), id)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessDeleteData,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customerId"]
	dataResponse, err := controller.CustomerRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) VerifyOtp(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	helpers.ReadFromRequestBody(r, &data)

	otp := controller.UserOTPRepository.FindSpesificData(r.Context(), entity.UserOTP{
		OtpCustomerId: data["customer_id"].(string),
		OtpStatus:     0,
		OtpNumber:     data["otp"].(string),
	})

	/* check apakah otp masih ada */
	if otp == nil {
		panic(exceptions.NewNotFoundError("Data tidak ditemukan, silahkan login kembali"))
	}

	/* check apakah masa aktif otp masih berlaku */

	if helpers.CreateDateTime().Unix() > otp[0].OtpExpired.Unix() {
		panic(exceptions.NewBadRequestError("Kode OTP anda sudah kadaluarsa, silahakan mengirim ulang OTP"))
	}

	/* update status di customer */
	controller.CustomerRepository.Update(r.Context(), "customer_status", entity.Customer{
		CustomerStatus: 1,
	}, data["customer_id"].(string))

	controller.UserOTPRepository.Update(r.Context(), entity.UserOTP{
		OtpId:         otp[0].OtpId,
		OtpCustomerId: otp[0].OtpCustomerId,
		OtpNumber:     otp[0].OtpNumber,
		OtpStatus:     1,
		OtpExpired:    otp[0].OtpExpired,
	}, otp[0].OtpId)

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		// Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	Qlimit := query.Get("limit")
	Qoffset := query.Get("offset")
	search := query.Get("search")

	if Qlimit == "" {
		Qlimit = "10"
	}

	if Qoffset == "" {
		Qoffset = "0"
	}

	limit, _ := strconv.Atoi(Qlimit)
	offset, _ := strconv.Atoi(Qoffset)

	nextOffset := limit + offset

	conf := map[string]interface{}{
		"limit":    limit,
		"offset":   offset,
		"search":   search,
		"customer": query.Get("customer"),
	}

	w.Header().Add("offset", fmt.Sprint(nextOffset))
	w.Header().Add("Access-Control-Expose-Headers", "offset")

	dataResponse := controller.CustomerRepository.FindAll(r.Context(), conf)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) DoLogin(w http.ResponseWriter, r *http.Request) {
	dataLogin := map[string]interface{}{}
	helpers.ReadFromRequestBody(r, &dataLogin)

	user := controller.CustomerRepository.FindSpesificData(r.Context(), entity.Customer{
		CustomerPhonenumber: dataLogin["phonenumber"].(string),
	})
	if user == nil {
		panic(exceptions.NewUnAuthorizedError("Gagal melakukan login, periksa kembali nomor hp dan password anda!"))
	}

	checkPassword := bcrypt.CompareHashAndPassword([]byte(user[0].CustomerPassword), []byte(dataLogin["password"].(string)))

	if checkPassword != nil {
		panic(exceptions.NewUnAuthorizedError("Gagal melakukan login, periksa kembali nomor hp dan password anda!"))
	}

	if user[0].CustomerStatus == 2 {
		panic(exceptions.NewUnAuthorizedError("Akun anda sudah diblokir, silahkan hubungi administator untuk membuka blokir"))
	}

	/* success login */
	/* remove tempcart on this user */
	controller.TempCartRepository.DeleteSpesificData(r.Context(), entity.TempCart{
		TempCartUserId: user[0].CustomerId,
	})

	/* add fcm token if user active */
	if user[0].CustomerStatus == 1 {
		userLog := controller.UserLogRepository.FindSpesificData(r.Context(), entity.UserLog{
			LogUserUserId: user[0].CustomerId,
			LogUserToken:  dataLogin["UserFcmToken"].(string),
		})
		if userLog == nil || userLog[0].LogUserLogoutDate.IsZero() {
			decodedByte, _ := base64.StdEncoding.DecodeString(dataLogin["UserDeviceMetadata"].(string))
			/* insert token to table */
			controller.UserLogRepository.Create(r.Context(), entity.UserLog{
				LogUserUserId:   user[0].CustomerId,
				LogUserToken:    dataLogin["UserFcmToken"].(string),
				LogUserMetadata: string(decodedByte),
			})
		}
	}
	otp := entity.UserOTP{}
	if user[0].CustomerStatus == 0 {
		/* otp */
		rand.Seed(time.Now().UTC().UnixNano())
		rdm := fmt.Sprint(rand.Int())

		otp = controller.UserOTPRepository.Create(r.Context(), entity.UserOTP{
			OtpCustomerId: user[0].CustomerId,
			OtpNumber:     rdm[:6],
		})
		/* sent otp */
		helpers.SendWhatsapp(r.Context(), map[string]interface{}{
			"phonenumber": user[0].CustomerPhonenumber,
			"text":        "eCanteen\n\nKode Verifikasi Anda : *" + otp.OtpNumber + "*\nakan berlaku selama 15 menit.\n\nPENTING !!! DEMI KEAMANAN AKUN ANDA, JANGAN BERIKAN KODE RAHASIA INI KEPADA SIAPAPUN,  Terima Kasih.",
		})
	}

	claims := &jwt.MapClaims{
		"customer_id":          user[0].CustomerId,
		"customer_phonenumber": user[0].CustomerPhonenumber,
		"customer_name":        user[0].CustomerName,
		"customer_status":      user[0].CustomerStatus,
		"customer":             user[0],
		"Otp":                  otp,
		"HasAccessCashier":     0,
		"exp":                  time.Now().Add(1000 * time.Hour).Unix(),
		"iss":                  config.GetEnv("APP_NAME"),
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

func (controller *UseCaseImpl) AddLog(w http.ResponseWriter, r *http.Request) {
	dataLogin := map[string]interface{}{}
	helpers.ReadFromRequestBody(r, &dataLogin)

	userLog := controller.UserLogRepository.FindSpesificData(r.Context(), entity.UserLog{
		LogUserUserId: dataLogin["CustomerId"].(string),
		LogUserToken:  dataLogin["UserFcmToken"].(string),
	})
	if userLog == nil || userLog[0].LogUserLogoutDate.IsZero() {
		decodedByte, _ := base64.StdEncoding.DecodeString(dataLogin["UserDeviceMetadata"].(string))
		/* insert token to table */
		controller.UserLogRepository.Create(r.Context(), entity.UserLog{
			LogUserUserId:   dataLogin["CustomerId"].(string),
			LogUserToken:    dataLogin["UserFcmToken"].(string),
			LogUserMetadata: string(decodedByte),
		})
	}

	webResponse := handler.WebResponse{
		Error:   false,
		Message: "Berhasil menyimpan data",
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) DoLogout(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	helpers.ReadFromRequestBody(r, &data)

	/* get data user token */

	log := controller.UserLogRepository.FindSpesificData(r.Context(), entity.UserLog{
		LogUserUserId: data["customer_id"].(string),
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

func (controller *UseCaseImpl) SendOTPResetPassword(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	helpers.ReadFromRequestBody(r, &data)
	customer := controller.CustomerRepository.FindSpesificData(r.Context(), entity.Customer{
		CustomerPhonenumber: data["CustomerPhonenumber"].(string),
	})

	if customer == nil {
		panic(exceptions.NewNotFoundError("Tidak ada akun yang tertaut dengan nomor yang anda masukkan"))
	}

	if customer[0].CustomerStatus == 2 {
		panic(exceptions.NewUnAuthorizedError("Akun anda sudah diblokir, silahkan hubungi administator untuk membuka blokir"))
	}
	/* otp */
	rand.Seed(time.Now().UTC().UnixNano())
	rdm := fmt.Sprint(rand.Int())

	otp := controller.UserOTPRepository.Create(r.Context(), entity.UserOTP{
		OtpCustomerId: customer[0].CustomerId,
		OtpNumber:     rdm[:6],
	})
	/* sent otp */
	helpers.SendWhatsapp(r.Context(), map[string]interface{}{
		"phonenumber": customer[0].CustomerPhonenumber,
		"text":        "eCanteen\n\nKode verifikasi untuk mengubah kata sandi anda : *" + otp.OtpNumber + "*\nakan berlaku sampai *" + otp.OtpExpired.String() + "*.\n\nPENTING !!! DEMI KEAMANAN AKUN ANDA, JANGAN BERIKAN KODE RAHASIA INI KEPADA SIAPAPUN,  Terima Kasih.",
	})

	dataResponse := map[string]interface{}{
		"otp":      otp,
		"customer": customer[0],
	}
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessCreateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) ChangePassword(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	helpers.ReadFromRequestBody(r, &data)
	dataResponse := controller.CustomerRepository.Update(r.Context(), "customer_password", entity.Customer{
		CustomerPassword: helpers.EncryptPassword(data["customer_password"].(string)),
	}, data["customer_id"].(string))
	webResponse := handler.WebResponse{
		Error:   false,
		Message: "Kata sandi berhasil dirubah, silahkan masuk untuk melanjutkan",
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
