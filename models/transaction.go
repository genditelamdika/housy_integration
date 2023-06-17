package models

type Transaction struct {
	ID      int    `json:"id"`
	Chekin  string `json:"chekin" form:"chekin" gorm:"type: varchar(255)"`
	Chekout string `json:"chekout" form:"chekout" gorm:"type: varchar(255)"`
	HouseID int    `json:"houseid"`
	House   House  `json:"house"`
	UserID  int    `json:"userid"`
	User    User   `json:"user"`
	Total   int    `json:"total" form:"total"`
	Status  string `json:"status" form:"status" gorm:"type: varchar(255)"`
}
