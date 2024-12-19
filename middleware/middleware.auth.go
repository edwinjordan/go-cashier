package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/jolebo/e-canteen-cashier-api/config"
	"github.com/jolebo/e-canteen-cashier-api/pkg/exceptions"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/* lanjutkan jika tidak mengarah ke menu utama */
		byPassUrl := []string{
			"/api/kasir/login",
			"/api/kasir/logout",
			"/api/customer/login",
			"/api/customer",
			"/api/customer/verifyOtp",
			"/api/customer/sentOTPResetPassword",
			"/api/customer/change_password",
			"/api/customer/logout",
			"/api/kasir/version",
			"/api/shop/version",
		}
		for _, v := range byPassUrl {
			if v == r.URL.Path {
				next.ServeHTTP(w, r)
				return
			}
		}

		if r.URL.Path == "/api/customer" && r.Method == "POST" {
			next.ServeHTTP(w, r)
			return
		}

		/* ambil data token */
		authHeader := r.Header.Get("Authorization")

		/* check apakah ada string bearer */
		if !strings.Contains(authHeader, "Bearer") {
			panic(exceptions.NewUnAuthorizedError("Anda tidak memiliki akses ke aplikasi ini"))
		}

		/* hapus string bearer */
		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

		/* proses parse token */
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			/* jika method tidak sesuai maka tolak */
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Anda tidak memiliki akses ke aplikasi ini")
			} else if method != jwt.SigningMethodHS256 {
				return nil, errors.New("Anda tidak memiliki akses ke aplikasi ini")
			}
			return []byte(config.GetEnv("SECRET_KEY")), nil
		})

		if err != nil {
			panic(exceptions.NewUnAuthorizedError(err.Error()))
		}

		/* ambil data payload */
		claims, ok := token.Claims.(jwt.MapClaims)

		/* jika token tidak valid atau payload tidak ada */
		if !token.Valid || !ok {
			panic(exceptions.NewUnAuthorizedError("Anda tidak memiliki akses ke aplikasi ini"))
		}
		/* masukkan ke context agar bisa digunakan */
		ctx := context.WithValue(r.Context(), "userslogin", claims)
		r = r.WithContext(ctx)

		if UserAuth(r) {
			next.ServeHTTP(w, r)
		} else {
			panic(exceptions.NewUnAuthorizedError("invalid access"))
		}

		/* get data payload from context */
		/* data := r.Context().Value("userlogin").(jwt.MapClaims)
		fmt.Println(data["akun_email"]) */

	})
}

func UserAuth(r *http.Request) bool {
	data := r.Context().Value("userslogin").(jwt.MapClaims)["HasAccessCashier"].(float64)
	arrURIPath := strings.Split(r.URL.Path, "/")
	status := true
	for _, v := range arrURIPath {
		if v == "kasir" && data != 1 {
			status = false
		}
	}
	return status
}
