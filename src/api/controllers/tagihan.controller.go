package controllers

import (
	"math/rand"
	"net/http"
	"src/api/models"
	"src/entity"
	"src/utils"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

func CreateTagihan(c echo.Context) error {
	s := new(entity.Tagihan)

	if err := c.Bind(s); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)
	if claims.IdRT == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Maaf anda tidak memiliki akses ini",
		})
	}

	s.IdKeluarga = claims.IdKeluarga

	//ini aku comment karena method ValidateCreate() belum ada di entity
	/*
		if err := s.ValidateCreate(); err.Code > 0 {
			return utils.ResponseError(c, err)
		}
	*/

	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	s.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	s.CreatedAt = time.Now()

	Tagihan, err := models.CreateTagihan(c, s)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataPersuratan(c, utils.JSONResponseDataPersuratan{
		Code:             http.StatusCreated,
		CreatePersuratan: Tagihan,
		Message:          "Berhasil",
	})
}

func GetAllTagihan(c echo.Context) error {

	allTagihan, err := models.GetAllTagihan(c, "")
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataPersuratan(c, utils.JSONResponseDataPersuratan{
		Code:             http.StatusOK,
		GetAllPersuratan: allTagihan,
		Message:          "Berhasil",
	})
}

func GetTagihanByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	s, err := models.GetPersuratanByID(c, id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseDataPersuratan(c, utils.JSONResponseDataPersuratan{
		Code:              http.StatusOK,
		GetPersuratanByID: s,
		Message:           "Berhasil",
	})
}

func UpdateTagihanById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	s := new(entity.Tagihan)

	if err := c.Bind(s); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	_, err := models.GetTagihanByID(c, id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	s.UpdatedAt = time.Now()

	_, err = models.UpdateTagihanById(c, id, s)
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

func SoftDeleteTagihanById(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id todak valid",
		})
	}

	_, err := models.GetTagihanByID(c, id)

	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.SoftDeleteTagihanById(c, id)

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
