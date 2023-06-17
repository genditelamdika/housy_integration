package usersdto

type CreateUserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" form:"password" validate:"required"`
	Role     string `json:"role" form:"role"`
	Gender   string `json:"gender" form:"password" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Subcribe bool   `json:"subcribe" form:"password" `
}

type UpdateUserRequest struct {
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	// Role     string `json:"role" form:"role"`
	Gender  string `json:"gender" form:"gender" `
	Address string `json:"address" form:"address" validate:"required"`
	Image   string `json:"image" form:"image" gorm:"type: varchar(255)"`
}
