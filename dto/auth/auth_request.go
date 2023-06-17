package authdto

type AuthRequest struct {
	Email    string `json:"email" gorm:"type: varchar(255)" validate:"required"`
	Password string `json:"password" gorm:"type: varchar(255)" validate:"required"`
	Fullname string `json:"fullname" gorm:"type: varchar(255)" validate:"required"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Gender   string `json:"gender" gorm:"type: varchar(255)"`
	Address  string `json:"address" gorm:"type: varchar(255)"`
	// Subcribe bool   `json:"subcribe" gorm:"type: text"`
	UserId int `json:"-" gorm:"type: int"`
}

type LoginRequest struct {
	Email    string `json:"email" gorm:"type: varchar(255)" validate:"required"`
	Password string `json:"password" gorm:"type: varchar(255)" validate:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" gorm:"type: varchar(255)" validate:"required"`
	NewPassword string `json:"new_password" gorm:"type: varchar(255)" validate:"required"`
}
