package db

import (
	"math/rand"
	"src/entity"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

func SeedAduan(db *gorm.DB, listIdRT, listIdWarga []string) []string {
	entropy1 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id1 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy1).String()
	data1 := entity.Aduan{
		Id:        Id1,
		IdWarga:   listIdWarga[0],
		IdRT:      listIdRT[0],
		Judul:     "Aduan 1",
		Gambar:    "default_image",
		Deskripsi: "Deskripsi Aduan 1",
		CreatedBy: "Agustina",
		CreatedAt: time.Now(),
	}

	db.Create(&data1)

	// Data 2
	entropy2 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id2 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy2).String()
	data2 := entity.Aduan{
		Id:        Id2,
		IdWarga:   listIdWarga[1],
		IdRT:      listIdRT[0],
		Judul:     "Aduan 2",
		Gambar:    "default_image",
		Deskripsi: "Deskripsi Aduan 2",
		CreatedBy: "Ronals",
		CreatedAt: time.Now(),
	}

	db.Create(&data2)

	return []string{Id1, Id2}
}
