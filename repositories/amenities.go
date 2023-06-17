package repositories

import (
	"housy/models"

	"gorm.io/gorm"
)

type AmenitiesRepository interface {
	FindAmenities() ([]models.Amenities, error)
	GetAmenitie(ID int) (models.Amenities, error)
	CreateAmenitie(amenitie models.Amenities) (models.Amenities, error)
	UpdateAmenitie(amenitie models.Amenities, Id int) (models.Amenities, error)
	DeleteAmenitie(amenitie models.Amenities) (models.Amenities, error)
	// UpdateFullcounter(amenitie models.Amenities) (models.Amenities, error)
	// GetAmenitieshouse(ID int) (models.Amenities, error)
	// GetTransactionCounterQty(tripID int) (int, error)
	// GetCategoryfilm(ID int) (models.Category, error)
}

func RepositoryAmenities(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAmenities() ([]models.Amenities, error) {
	var amenities []models.Amenities
	err := r.db.Find(&amenities).Error

	return amenities, err
}

func (r *repository) GetAmenitie(ID int) (models.Amenities, error) {
	var amenitie models.Amenities
	err := r.db.First(&amenitie, ID).Error

	return amenitie, err
}

func (r *repository) CreateAmenitie(amenitie models.Amenities) (models.Amenities, error) {
	err := r.db.Create(&amenitie).Error

	return amenitie, err
}
func (r *repository) UpdateAmenitie(amenitie models.Amenities, Id int) (models.Amenities, error) {
	err := r.db.Model(&amenitie).Updates(&amenitie).Error

	return amenitie, err
}

func (r *repository) DeleteAmenitie(amenitie models.Amenities) (models.Amenities, error) {
	err := r.db.Delete(&amenitie).Error

	return amenitie, err
}
