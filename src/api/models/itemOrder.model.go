package models

import (
	"errors"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
)

func CreateItemOrder(c echo.Context, i *entity.ItemOrder) (entity.ItemOrder, error) {
	db := db.GetDB(c)

	err := db.Create(&i)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.ItemOrder{}, err.Error
	}

	if err.RowsAffected == 0 {
		return entity.ItemOrder{}, errors.New("gagal menambahkan item order")
	}

	return *i, nil
}

func CreateBatchItemOrder(c echo.Context, items []entity.ItemOrder) ([]entity.ItemOrder, error) {
	db := db.GetDB(c)

	err := db.Create(items)
	if err.Error != nil {
		c.Logger().Error(err)
		return []entity.ItemOrder{}, err.Error
	}

	if err.RowsAffected == 0 {
		return []entity.ItemOrder{}, errors.New("gagal menambahkan item order")
	}

	return items, nil
}

func GetAllItemOrder(c echo.Context) ([]entity.ItemOrder, error) {
	var item []entity.ItemOrder
	db := db.GetDB(c)

	err := db.Find(&item)
	if err.Error != nil {
		c.Logger().Error(err)
		return item, err.Error
	}
	return item, nil
}

func GetItem(c echo.Context, id string) ([]entity.ItemOrder, error) {
	var item []entity.ItemOrder
	db := db.GetDB(c)

	err := db.First(&item, "id = ?", id)
	if err.Error != nil {
		c.Logger().Error(err)
		return item, errors.New("id tidak ditemukan atau tidak valid")
	}
	return item, nil
}

func GetItemByIDOrder(c echo.Context, id_order string) ([]entity.ItemOrder, error) {
	var item []entity.ItemOrder
	db := db.GetDB(c)

	err := db.First(&item, "id_order = ?", id_order)
	if err.Error != nil {
		c.Logger().Error(err)
		return item, errors.New("id tidak ditemukan")
	}
	return item, nil
}

func HardDeleteItemOrder(c echo.Context, id_order string) error {
	item := entity.ItemOrder{
		Id: id_order,
	}
	db := db.GetDB(c)

	err := db.Unscoped().Delete(&item)
	if err.Error != nil {
		c.Logger().Error(err)
		return err.Error
	}
	return nil
}
