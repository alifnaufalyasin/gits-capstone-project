package entity

import (
	"net/http"
	"src/utils"
	"time"

	"gorm.io/gorm"
)

type ItemOrder struct {
	Id         string          `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdProduk   string          `gorm:"type:varchar(50);not null" json:"id_produk" form:"id_produk"`
	IdOrder    string          `gorm:"type:varchar(50)" json:"id_order,omitempty" form:"id_order,omitempty"`
	Jumlah     int64           `gorm:"not null" json:"jumlah" form:"jumlah"`
	HargaTotal int64           `gorm:"not null" json:"harga_total" form:"harga_total"`
	Catatan    string          `gorm:"type:varchar;not null" json:"catatan" form:"catatan"`
	CreatedAt  time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt  time.Time       `gorm:"type:timestamptz;" json:"updated_at"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (ItemOrder) TableName() string {
	return "item_order"
}

func (ord ItemOrder) ValidateCreate() utils.Error {
	if ord.Jumlah == 0 {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Jumlah tidak boleh 0",
		}
	}
	return utils.Error{}
}
