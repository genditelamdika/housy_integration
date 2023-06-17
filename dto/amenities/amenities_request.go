package amenitiesdto

type CreateAmenitiesRequest struct {
	Name string `json:"name" form:"name" gorm:"type: varchar(255)"`
}

type UpdateAmenitiesRequest struct {
	Name string `json:"name" form:"name" gorm:"type: varchar(255)"`
}
