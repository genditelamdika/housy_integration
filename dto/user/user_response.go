package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	// Gender   string `json:"gender" form:"gender" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
	Role    string `json:"role"  gorm:"type: varchar(255)"`
	Image   string `json:"image" form:"image" gorm:"type: varchar(255)"`
	// Subcribe bool   `json:"subcribe" form:"subcribe"`
}

type ProfileResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Address  string `json:"address"`
}
