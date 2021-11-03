package controllers

import (
	"math/rand"
	"net/http"
	"src/api/models"
	"src/entity"
	"src/utils"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

func CreateKeluarga(c echo.Context) error {
	// Pertama inisiasi variable dulu
	k := new(entity.Keluarga)

	// kemudian ini buat dapetin request body dari mobile
	if err := c.Bind(k); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// terus ini ada validasi buat ngecek inputan dari reqeust body udah sesuai apa belum
	if err := k.ValidateCreate(); err.Code > 0 {
		return utils.ResponseError(c, err)
	}

	//Ini buat generate ID
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	k.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	// Ini buat masukin isi dari created_at nya
	k.CreatedAt = time.Now()

	// Ini fungsi dari models buat create data ke database
	keluarga, err := models.CreateKeluarga(c, k)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// Return datanya
	return utils.ResponseDataKeluarga(c, utils.JSONResponseDataKeluarga{
		Code:           http.StatusCreated,
		CreateKeluarga: keluarga,
		Message:        "Berhasil",
	})
}

func GetAllKeluarga(c echo.Context) error {
	allKeluarga, err := models.GetAllKeluarga(c, c.QueryParam("nama"))
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataKeluarga(c, utils.JSONResponseDataKeluarga{
		Code:           http.StatusOK,
		GetAllKeluarga: allKeluarga,
		Message:        "Berhasil",
	})
}

func GetAllKeluargaWithWarga(c echo.Context) error {
	allKeluarga, err := models.GetAllKeluargaWithEntity(c, c.QueryParam("nama"), "Warga")
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataKeluarga(c, utils.JSONResponseDataKeluarga{
		Code:                http.StatusOK,
		GetAllKeluargaWarga: allKeluarga,
		Message:             "Berhasil",
	})
}

func GetKeluargaByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	k, err := models.GetKeluargaByID(c, id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseDataKeluarga(c, utils.JSONResponseDataKeluarga{
		Code:            http.StatusOK,
		GetKeluargaByID: k,
		Message:         "Berhasil",
	})
}

func UpdateKeluargaById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	k := new(entity.Keluarga)

	if err := c.Bind(k); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	_, err := models.GetKeluargaByID(c, id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	k.UpdatedAt = time.Now()
	_, err = models.UpdateKeluargaById(c, id, k)
	if err != nil {
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

func SoftDeleteKeluargaById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	_, err := models.GetKeluargaByID(c, id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.SoftDeleteKeluargaById(c, id)
	if err != nil {
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
