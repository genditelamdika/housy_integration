package housedto

import (
	"gorm.io/datatypes"
)

type CreateHouseRequest struct {
	Nameproperty string         `json:"nameproperty" form:"nameproperty" gorm:"type: varchar(255)"`
	City         string         `json:"city" form:"city" gorm:"type: varchar(255)"`
	Addres       string         `json:"addres" form:"addres" gorm:"type: varchar(255)"`
	Year         string         `json:"year" form:"year" gorm:"type: varchar(255)"`
	Area         string         `json:"area" form:"area" gorm:"type: varchar(255)"`
	Price        int            `json:"price" form:"price" gorm:"type: int"`
	Tor          int            `json:"tor" form:"tor" gorm:"type: varchar(255)"`
	Amenities    datatypes.JSON `json:"amenities" form:"amenities" gorm:"type: JSON"`
	Description  string         `json:"description" form:"description" gorm:"type: varchar(255)"`
	Bedroom      int            `json:"bedroom" form:"bedroom" gorm:"type: int"`
	Bathroom     int            `json:"bathroom" form:"bathroom" gorm:"type: int"`
	Image        string         `json:"image" form:"image"`
}

type UpdateHouseRequest struct {
	Nameproperty string         `json:"nameproperty" form:"nameproperty" gorm:"type: varchar(255)"`
	City         string         `json:"city" form:"city" gorm:"type: varchar(255)"`
	Addres       string         `json:"addres" form:"addres" gorm:"type: varchar(255)"`
	Price        int            `json:"price" form:"price" gorm:"type: int"`
	Tor          int            `json:"tor" form:"tor" gorm:"type: varchar(255)"`
	Amenities    datatypes.JSON `json:"amenities" form:"amenities" gorm:"type: JSON"`
	Bedroom      int            `json:"bedroom" form:"bedroom" gorm:"type: int"`
	Bathroom     int            `json:"bathroom" form:"bathroom" gorm:"type: int"`
	Image        string         `json:"image" form:"image"`
}
