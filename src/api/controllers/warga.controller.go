package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"src/api/models"
	"src/entity"
	"src/utils"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

func CreateWarga(c echo.Context) error {
	// Pertama inisiasi variable dulu
	w := new(entity.Warga)

	// kemudian ini buat dapetin request body dari mobile
	if err := c.Bind(w); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	w.Gambar = "default_image"

	if strings.HasPrefix(w.NoHandphone, "62") {
		w.NoHandphone = fmt.Sprintf("0%s", strings.SplitAfter(w.NoHandphone, "62")[1])
	}
	if strings.HasPrefix(w.NoHandphone, "+62") {
		w.NoHandphone = fmt.Sprintf("0%s", strings.SplitAfter(w.NoHandphone, "+62")[1])
	}

	// terus ini ada validasi buat ngecek inputan dari reqeust body udah sesuai apa belum
	if err := w.ValidateCreate(); err.Code > 0 {
		c.Logger().Error(err)
		return utils.ResponseError(c, err)
	}

	k, err := models.GetKeluargaByKode(c, w.KodeKeluarga)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	w.IdKeluarga = k.Id

	cek, err := models.GetWargaByEmail(c, w.Email)
	if err != nil && err.Error() != "email tidak ditemukan atau tidak valid" {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	if cek.Id != "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Email sudah terdaftar",
		})
	}

	//Ini buat generate ID
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	w.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	pass := w.Password
	// Ini buat masukin isi dari created_at nya
	w.CreatedAt = time.Now()
	w.Password = utils.HashPassword(w.Password, w.Id)

	// Ini fungsi dari models buat create data ke database
	_, err = models.CreateWarga(c, w)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	ents := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)

	_, err = models.CreateKeranjang(c, &entity.Keranjang{
		Id:      ulid.MustNew(ulid.Timestamp(time.Now()), ents).String(),
		IdWarga: w.Id,
	})
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// Return datanya
	return loginWarga(c, pass, k.IdRT, w)
}

func GetAllWarga(c echo.Context) error {
	allWarga, err := models.GetAllWarga(c, c.QueryParam("id_keluarga"), c.QueryParam("nama"))
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataWarga(c, utils.JSONResponseDataWarga{
		Code:        http.StatusOK,
		GetAllWarga: allWarga,
		Message:     "Berhasil",
	})
}

func GetWargaByID(c echo.Context) error {
	var id string
	paramid := c.Param("id")
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)

	if paramid != "" {
		id = paramid
	} else if claims.UserId != "" {
		id = claims.UserId
	} else {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	w, err := models.GetWargaByID(c, id)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseDataWarga(c, utils.JSONResponseDataWarga{
		Code:         http.StatusOK,
		GetWargaByID: w,
		Message:      "Berhasil",
	})
}

func UpdateWargaById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	w := new(entity.Warga)

	if err := c.Bind(w); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := w.ValidateUpdate(); err.Code > 0 {
		c.Logger().Error(err)
		return utils.ResponseError(c, err)
	}

	_, err := models.GetWargaByID(c, id)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	if w.Password != "" {
		w.Password = utils.HashPassword(w.Password, w.Email)
	}
	w.UpdatedAt = time.Now()

	_, err = models.UpdateWargaById(c, id, w)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.Response(c, utils.JSONResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
	})
}

func SoftDeleteWargaById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	_, err := models.GetWargaByID(c, id)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.SoftDeleteWargaById(c, id)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.Response(c, utils.JSONResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
	})
}

func LoginWarga(c echo.Context) error {
	w := new(entity.Warga)

	if err := c.Bind(w); err != nil {
		c.Logger().Error(err)
		return utils.ResponseErrorLogin(c, utils.ErrorLogin{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	keluarga, err := models.GetKeluargaByEmail(c, w.Email)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseErrorLogin(c, utils.ErrorLogin{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	var warga entity.Warga
	if len(keluarga.Warga) == 0 {
		return utils.ResponseErrorLogin(c, utils.ErrorLogin{
			Code:    http.StatusInternalServerError,
			Message: "Akun tidak ditemukan",
		})
	} else {
		warga = keluarga.Warga[0]
	}

	isValid := utils.CheckPassword(w.Password, warga.Id, warga.Password)
	if !isValid {
		return utils.ResponseErrorLogin(c, utils.ErrorLogin{
			Code:    http.StatusBadRequest,
			Message: "Password yang anda masukkan salah",
		})
	}
	if w.TokenFirebase != "" {
		_, err := models.UpdateWargaById(c, warga.Id, &entity.Warga{TokenFirebase: w.TokenFirebase})
		if err != nil {
			c.Logger().Error(err)
			return utils.ResponseErrorLogin(c, utils.ErrorLogin{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
	}

	token, err := utils.GenerateTokenWarga(c, warga.Nama, warga.Email, warga.Id, warga.IdKeluarga, keluarga.IdRT, utils.ExpiredHour)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseErrorLogin(c, utils.ErrorLogin{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return utils.ResponseLogin(c, utils.JSONResponseLogin{
		Code:    http.StatusOK,
		Token:   token,
		Message: "Berhasil",
	})

}

func loginWarga(c echo.Context, pass, id_rt string, w *entity.Warga) error {
	isValid := utils.CheckPassword(pass, w.Id, w.Password)
	if !isValid {
		return utils.ResponseErrorLogin(c, utils.ErrorLogin{
			Code:    http.StatusBadRequest,
			Token:   "",
			Message: "Password yang anda masukkan salah",
		})
	}

	if w.TokenFirebase != "" {
		_, err := models.UpdateWargaById(c, w.Id, &entity.Warga{TokenFirebase: w.TokenFirebase})
		if err != nil {
			c.Logger().Error(err)
			return utils.ResponseErrorLogin(c, utils.ErrorLogin{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
	}

	token, err := utils.GenerateTokenWarga(c, w.Nama, w.Email, w.Id, w.IdKeluarga, id_rt, utils.ExpiredHour)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseErrorLogin(c, utils.ErrorLogin{
			Code:    http.StatusBadRequest,
			Token:   "",
			Message: err.Error(),
		})
	}
	reqNotif := utils.RequestSendNotificationToken{
		To: w.TokenFirebase,
		Notification: utils.Notification{
			Title: fmt.Sprintf("Hi %s, Selamat datang di Sma-RT", w.Nama),
			Body:  "Selamat menggunakan aplikasi kami",
		},
	}
	err = utils.SendNotificationToken(c, reqNotif)
	if err != nil {
		c.Logger().Error("Notif Error: ", err)
	}

	return utils.ResponseLogin(c, utils.JSONResponseLogin{
		Code:    http.StatusOK,
		Token:   token,
		Message: "Berhasil",
	})

}

type ForgetPasswordRequest struct {
	Email string `json:"email" form:"email"`
}

func ForgetPasswordWarga(c echo.Context) error {
	fp := new(ForgetPasswordRequest)

	if err := c.Bind(fp); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	warga, err := models.GetWargaByEmail(c, fp.Email)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	fpId := ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
	kode := models.GenerateKodeForgetPasswordWarga(c, 6)
	forgetPass := entity.ForgetPasswordWarga{
		Id:        fpId,
		IdWarga:   warga.Id,
		Kode:      kode,
		CreatedAt: time.Now(),
	}

	fpw, err := models.CreateForgetPasswordWarga(c, &forgetPass)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	warga.ForgetPasswordWarga = &fpw
	_, err = models.UpdateWargaById(c, warga.Id, &warga)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	err = utils.SendEmail(c, fp.Email, "Kode Reset Password", fmt.Sprintf("Berikut ini adalah kode Verifikasi untuk reset password akun warga anda <br><br> Kode: <b>%s</b> <br><br> abaikan jika anda tidak sedang mereset password", kode))
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.Response(c, utils.JSONResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
	})
}

type ResetPasswordRequest struct {
	Kode     string `json:"kode" form:"kode"`
	Password string `json:"password" form:"password"`
}

func ResetPasswordWargaByKode(c echo.Context) error {
	rp := new(ResetPasswordRequest)

	if err := c.Bind(rp); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if !utils.CheckStrengthPassword(rp.Password) {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Password panjangnya min. 8 karakter, serta mengandung min. 1 huruf besar, 1 huruf kecil, dan 1 angka!",
		})
	}

	w, err := models.GetWargaByForgetPasswordKode(c, rp.Kode)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	w.ForgetPasswordWarga = &entity.ForgetPasswordWarga{}
	w.Password = utils.HashPassword(rp.Password, w.Id)

	_, err = models.UpdateWargaById(c, w.Id, &w)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.DeleteForgetPasswordWarga(c, rp.Kode)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.Response(c, utils.JSONResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
	})
}

type ChangePasswordRequest struct {
	Password    string `json:"password" form:"password"`
	NewPaswword string `json:"new_password" form:"new_password"`
}

func GantiPasswordWarga(c echo.Context) error {
	cp := new(ChangePasswordRequest)

	if err := c.Bind(cp); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if !utils.CheckStrengthPassword(cp.NewPaswword) {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Password panjangnya min. 8 karakter, serta mengandung min. 1 huruf besar, 1 huruf kecil, dan 1 angka!",
		})
	}

	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)

	w, _ := models.GetWargaByID(c, claims.UserId)

	isValid := utils.CheckPassword(cp.Password, claims.UserId, w.Password)
	if !isValid {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Password yang anda masukkan salah",
		})
	}

	_, err := models.UpdateWargaById(c, claims.UserId, &entity.Warga{Password: utils.HashPassword(cp.NewPaswword, w.Id)})
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.Response(c, utils.JSONResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
	})
}
