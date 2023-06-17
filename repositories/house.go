package repositories

import (
	"housy/models"

	"gorm.io/gorm"
)

type HouseRepository interface {
	FindHouses() ([]models.House, error)
	GetHouse(ID int) (models.House, error)
	CreateHouse(house models.House) (models.House, error)
	UpdateHouse(house models.House, Id int) (models.House, error)
	DeleteHouse(house models.House) (models.House, error)
	// UpdateFullcounter(house models.House) (models.House, error)
	// GetAmenitieshouse(ID int) (models.Amenities, error)
	// GetTransactionCounterQty(tripID int) (int, error)
	// GetCategoryfilm(ID int) (models.Category, error)
}

func RepositoryHouse(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindHouses() ([]models.House, error) {
	var houses []models.House
	err := r.db.Find(&houses).Error

	return houses, err
}

func (r *repository) GetHouse(ID int) (models.House, error) {
	var house models.House
	err := r.db.First(&house, ID).Error

	return house, err
}

func (r *repository) CreateHouse(house models.House) (models.House, error) {
	err := r.db.Create(&house).Error

	return house, err
}
func (r *repository) UpdateHouse(house models.House, Id int) (models.House, error) {
	err := r.db.Model(&house).Updates(&house).Error

	return house, err
}

func (r *repository) DeleteHouse(house models.House) (models.House, error) {
	err := r.db.Delete(&house).Error

	return house, err
}

// func (r *repository) GetAmenitieshouse(Id int) (models.Amenities, error) {
// 	var amenities models.Amenities
// 	err := r.db.First(&amenities, Id).Error
// 	return amenities, err
// 	// err := r.db.Delete(&film).Error

// 	// return cate, err
// }
