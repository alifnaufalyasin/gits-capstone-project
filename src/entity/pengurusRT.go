package entity

import (
	"time"

	"gorm.io/gorm"
)

type PengurusRT struct {
	Id        string          `gorm:"type:varchar(50);primaryKey" json:"id"`
	IdRT      string          `gorm:"type:varchar(50);not null" json:"id_rt"`
	Nama      string          `gorm:"type:varchar(50);not null" json:"nama"`
	Email     string          `gorm:"type:varchar(120);not null" json:"email"`
	Password  string          `gorm:"type:varchar(100);not null" json:"password"`
	CreatedAt time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt time.Time       `gorm:"type:timestamptz;" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (PengurusRT) TableName() string {
	return "pengurus_rt"
}