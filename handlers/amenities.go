package handlers

import (
	amenitiesdto "housy/dto/amenities"
	dto "housy/dto/result"
	"housy/models"
	"housy/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerAmenities struct {
	AmenitiesRepository repositories.AmenitiesRepository
}

func HandlerAmenities(AmenitiesRepository repositories.AmenitiesRepository) *handlerAmenities {
	return &handlerAmenities{AmenitiesRepository}
}
func (h *handlerAmenities) FindAmenities(c echo.Context) error {
	countrys, err := h.AmenitiesRepository.FindAmenities()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: countrys})
}
func (h *handlerAmenities) GetAmenitie(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	country, err := h.AmenitiesRepository.GetAmenitie(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	// responseCategory := categoriesdto.CategoryResponse{
	// 	ID:   country.ID,
	// 	Name: country.Name,
	// 	Film: []categoriesdto.CategoryFilm{},
	// }

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponsecountry(country)})
}
func (h *handlerAmenities) CreateAmenitie(c echo.Context) error {
	request := new(amenitiesdto.CreateAmenitiesRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	country := models.Amenities{
		Name: request.Name,
		// Films: request.Film,
		// FilmID: request.FilmID,
		// Films: request.Films,
		// Email:    request.Email,
		// Password: request.Password,
	}

	data, err := h.AmenitiesRepository.CreateAmenitie(country)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponsecountry(data)})
}
func (h *handlerAmenities) UpdateAmenitie(c echo.Context) error {
	request := new(amenitiesdto.UpdateAmenitiesRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	amenitie, err := h.AmenitiesRepository.GetAmenitie(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		amenitie.Name = request.Name
	}

	data, err := h.AmenitiesRepository.UpdateAmenitie(amenitie, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
func (h *handlerAmenities) DeleteAmenitie(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	country, err := h.AmenitiesRepository.GetAmenitie(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.AmenitiesRepository.DeleteAmenitie(country)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func convertResponsecountry(u models.Amenities) models.Amenities {
	return models.Amenities{
		ID:   u.ID,
		Name: u.Name,
	}
}
