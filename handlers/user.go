package handlers

import (
	"context"
	"fmt"
	authdto "housy/dto/auth"
	dto "housy/dto/result"
	usersdto "housy/dto/user"
	"housy/models"
	"housy/pkg/bcrypt"
	"housy/repositories"
	"net/http"
	"os"
	"strconv"

	// "github.com/go-playground/validator/v10"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

func (h *handler) FindUsers(c echo.Context) error {
	users, err := h.UserRepository.FindUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: users})
}

func (h *handler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: user})
}
func (h *handler) UpdateUser(c echo.Context) error {
	// request := new(usersdto.UpdateUserRequest)
	// if err := c.Bind(&request); err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	// }
	// get the datafile here
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "uploads"})

	if err != nil {
		fmt.Println(err.Error())
	}

	request := usersdto.UpdateUserRequest{
		Fullname: c.FormValue("fullname"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
		Phone:    c.FormValue("phone"),
		Gender:   c.FormValue("gender"),
		Address:  c.FormValue("address"),
		Image:    resp.SecureURL,
	}

	id, _ := strconv.Atoi(c.Param("id"))

	profile, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Email != "" {
		profile.Email = request.Email
	}
	if request.Password != "" {
		profile.Password, _ = bcrypt.HashingPassword(request.Password)
	}
	if request.Fullname != "" {
		profile.Fullname = request.Fullname
	}
	if request.Phone != "" {
		profile.Phone = request.Phone
	}
	if request.Address != "" {
		profile.Address = request.Address
	}

	if request.Image != "" {
		profile.Image = request.Image
	}

	data, err := h.UserRepository.UpdateUser(profile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}

func (h *handler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.UserRepository.DeleteUser(user, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}
func (h *handler) ChangePassword(c echo.Context) error {
	request := new(authdto.ChangePasswordRequest)
	if err := c.Bind(request); err != nil {
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		return c.JSON(http.StatusBadRequest, response)
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, err := h.UserRepository.GetUser(int(userId))
	if err != nil {
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		return c.JSON(http.StatusBadRequest, response)
	}

	isValid := bcrypt.CheckPasswordHash(request.OldPassword, user.Password)
	if !isValid {
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "your old password doesn't match!"}
		return c.JSON(http.StatusBadRequest, response)
	}

	newPassword, err := bcrypt.HashingPassword(request.NewPassword)
	if err != nil {
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		return c.JSON(http.StatusInternalServerError, response)
	}

	user.Password = newPassword

	data, err := h.UserRepository.ChangePassword(user)
	if err != nil {
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	return c.JSON(http.StatusOK, response)
}

func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:       u.ID,
		Fullname: u.Fullname,
		Email:    u.Email,
		Password: u.Password,
		Phone:    u.Phone,
		Address:  u.Address,
		// Gender:   u.Gender,
	}
}
