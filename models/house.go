package models

import "gorm.io/datatypes"

type House struct {
	ID           int    `json:"id"`
	Nameproperty string `json:"nameproperty" form:"nameproperty" gorm:"type: varchar(255)"`
	City         string `json:"city" form:"city" gorm:"type: varchar(255)"`
	Addres       string `json:"addres" form:"addres" gorm:"type: varchar(255)"`
	Year         string `json:"year" form:"year" gorm:"type: varchar(255)"`
	Area         string `json:"area" form:"area" gorm:"type: varchar(255)"`
	Description  string `json:"description" form:"description" gorm:"type: varchar(255)"`
	Status       string `json:"status" form:"status" gorm:"type: varchar(255)"`
	Price        int    `json:"price" form:"price" gorm:"type: int"`

	Tor int `json:"tor" form:"tor" gorm:"type: varchar(255)"`
	// AmenitiesID  []int       `json:"amenitiesID"`
	Amenities datatypes.JSON `json:"amenities" form:"amenities" gorm:"type: JSON"`
	Bedroom   int            `json:"bedroom" form:"bedroom" gorm:"type: int"`
	Bathroom  int            `json:"bathroom" form:"bathroom" gorm:"type: int"`
	Image     string         `json:"image" form:"image"`
}
