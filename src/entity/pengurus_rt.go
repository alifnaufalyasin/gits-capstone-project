package entity

import (
	"net/http"
	"net/mail"
	"src/utils"
	"time"

	"gorm.io/gorm"
)

type PengurusRT struct {
	Id                     string                  `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdRT                   string                  `gorm:"type:varchar(50);not null" json:"id_rt" form:"id_rt"`
	TokenFirebase          string                  `gorm:"type:varchar" json:"token_firebase" form:"token_firebase"`
	ForgetPasswordPengurus *ForgetPasswordPengurus `gorm:"foreignKey:id_pengurus;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"forget_password_pengurus,omitempty" form:"forget_password_pengurus"`
	NoHandphone            string                  `gorm:"type:varchar(20);not null" json:"no_hp" form:"no_hp"`
	KodeRT                 string                  `gorm:"type:varchar(100); not null" json:"kode_rt,omitempty" form:"kode_rt"`
	Gender                 string                  `gorm:"type:varchar(20);not null" json:"gender" form:"gender"`
	Nama                   string                  `gorm:"type:varchar(50);not null" json:"nama" form:"nama"`
	Gambar                 string                  `gorm:"default:default_image" json:"gambar" form:"gambar"`
	Email                  string                  `gorm:"type:varchar(120);not null" json:"email" form:"email"`
	Password               string                  `gorm:"type:varchar(100);not null" json:"password" form:"password"`
	CreatedAt              time.Time               `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt              time.Time               `gorm:"type:timestamptz;" json:"updated_at"`
	DeletedAt              *gorm.DeletedAt         `json:"deleted_at,omitempty"`
}

func (PengurusRT) TableName() string {
	return "pengurus_rt"
}

func (prt PengurusRT) ValidateCreate() utils.Error {
	if prt.Nama == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Nama tidak boleh kosong",
		}
	}
	if _, err := mail.ParseAddress(prt.Email); err != nil {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Email tidak valid",
		}
	}
	if prt.KodeRT == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Kode RT tidak boleh kosong",
		}
	}
	if !utils.CheckStrengthPassword(prt.Password) {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Password panjangnya min. 8 karakter, serta mengandung min. 1 huruf besar, 1 huruf kecil, dan 1 angka!",
		}
	}
	if prt.Gambar == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Gambar tidak boleh kosong",
		}
	}
	if len(prt.NoHandphone) < 10 && len(prt.NoHandphone) > 13 {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Nomor handphone tidak valid (min 10 angka, max 13 angka)",
		}
	}
	if prt.Gender != "laki-laki" && prt.Gender != "perempuan" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "gender tidak valid",
		}
	}
	return utils.Error{}
}

func (prt PengurusRT) ValidateUpdate() utils.Error {
	if prt.Nama == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Nama tidak boleh kosong",
		}
	}
	if _, err := mail.ParseAddress(prt.Email); err != nil {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Email tidak valid",
		}
	}
	if prt.Gambar == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Gambar tidak boleh kosong",
		}
	}
	if len(prt.NoHandphone) < 10 && len(prt.NoHandphone) > 13 {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Nomor handphone tidak valid (min 10 angka, max 13 angka)",
		}
	}
	return utils.Error{}
}
