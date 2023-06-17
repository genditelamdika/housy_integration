package amenitiesdto

type AmenitiesResponse struct {
	Nameproperty string `json:"name" form:"name" gorm:"type: varchar(255)"`
}
