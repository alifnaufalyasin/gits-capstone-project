package entity

import (
	"net/http"
	"src/utils"
	"time"

	"gorm.io/gorm"
)

type Produk struct {
	Id            string          `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdKeluarga    string          `gorm:"type:varchar(50);not null" json:"id_keluarga" form:"id_keluarga"`
	IdRT          string          `gorm:"type:varchar(50);not null" json:"id_rt" form:"id_rt"`
	Nama          string          `gorm:"type:varchar(50);not null" json:"nama" form:"nama"`
	ItemOrder     []ItemOrder     `gorm:"foreignKey:id_produk;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"item_order,omitempty" form:"item_order"`
	ItemKeranjang []ItemKeranjang `gorm:"foreignKey:id_produk;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"item_Keranjang,omitempty" form:"item_Keranjang"`
	Detail        string          `gorm:"not null" json:"detail" form:"detail"`
	Gambar        string          `gorm:"not null" json:"gambar" form:"gambar"`
	Tersedia      string          `gorm:"default:true" json:"tersedia" form:"tersedia"`
	Harga         int64           `gorm:"not null" json:"harga" form:"harga"`
	CreatedAt     time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt     time.Time       `gorm:"type:timestamptz;" json:"updated_at"`
	DeletedAt     *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Produk) TableName() string {
	return "produk"
}

func (p Produk) ValidateCreate() utils.Error {
	if p.Nama == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Nama tidak boleh kosong",
		}
	}
	if p.Detail == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Detail produk tidak boleh kosong",
		}
	}
	if p.Gambar == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Gambar produk tidak boleh kosong",
		}
	}
	if p.Harga == 0 {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Harga tidak boleh kosong",
		}
	}
	if p.Tersedia != "true" && p.Tersedia != "false" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Format tersedia salah",
		}
	}
	return utils.Error{}
}
