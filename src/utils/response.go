package utils

import (
	"github.com/labstack/echo/v4"
)

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type JSONResponseData struct {
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type JSONResponseDataRT struct {
	Code      int64       `json:"code"`
	GetRTByID interface{} `json:"get_rt_by_id,omitempty"`
	GetAllRT  interface{} `json:"get_all_rt,omitempty"`
	CreateRT  interface{} `json:"create_rt,omitempty"`
	Message   string      `json:"message"`
}

type JSONResponseDataPengurusRT struct {
	Code              int64       `json:"code"`
	GetPengurusRTByID interface{} `json:"get_pengurus_rt_by_id,omitempty"`
	GetPengurusAllRT  interface{} `json:"get_all_pengurus_rt,omitempty"`
	CreatePengurusRT  interface{} `json:"create_pengurus_rt,omitempty"`
	Message           string      `json:"message"`
}

type JSONResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func Response(c echo.Context, res JSONResponse) error {
	return c.JSON(int(res.Code), res)
}

func ResponseData(c echo.Context, res JSONResponseData) error {
	return c.JSON(int(res.Code), res)
}

func ResponseDataRT(c echo.Context, res JSONResponseDataRT) error {
	return c.JSON(int(res.Code), res)
}

func ResponseDataPengurusRT(c echo.Context, res JSONResponseDataPengurusRT) error {
	return c.JSON(int(res.Code), res)
}

func ResponseError(c echo.Context, err Error) error {
	return c.JSON(int(err.Code), err)
}
