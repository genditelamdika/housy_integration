package handlers

import (
	"context"
	"fmt"
	housedto "housy/dto/house"
	dto "housy/dto/result"
	"housy/models"
	"housy/repositories"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
)

var path_file = "http://localhost:5000/uploads/"

type handlerHouse struct {
	HouseRepository repositories.HouseRepository
}

func HandlerHouse(HouseRepository repositories.HouseRepository) *handlerHouse {
	return &handlerHouse{HouseRepository}
}
func (h *handlerHouse) FindHouses(c echo.Context) error {
	houses, err := h.HouseRepository.FindHouses()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	// for i, p := range houses {
	// 	houses[i].Image = path_file + p.Image
	// }

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: houses})
}
func (h *handlerHouse) GetHouse(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	house, err := h.HouseRepository.GetHouse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	// house.Image = path_file + house.Image

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: house})
}
func (h *handlerHouse) CreateHouse(c echo.Context) error {
	// get the datafile here
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	// amenitiesid, _ := strconv.Atoi(c.FormValue("amenitiesid"))
	tor, _ := strconv.Atoi(c.FormValue("tor"))
	price, _ := strconv.Atoi(c.FormValue("price"))
	bedroom, _ := strconv.Atoi(c.FormValue("bedroom"))
	bathroom, _ := strconv.Atoi(c.FormValue("bathroom"))

	request := housedto.CreateHouseRequest{
		Nameproperty: c.FormValue("nameproperty"),
		Amenities:    datatypes.JSON(c.FormValue("amenities")),
		City:         c.FormValue("city"),
		Addres:       c.FormValue("addres"),
		Year:         c.FormValue("year"),
		Area:         c.FormValue("area"),
		Description:  c.FormValue("description"),
		Price:        price,
		Tor:          tor,
		Bedroom:      bedroom,
		Bathroom:     bathroom,
		Image:        dataFile,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

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

	house := models.House{
		Nameproperty: request.Nameproperty,
		// AmenitiesID:  request.AmenitiesID,
		Amenities:   request.Amenities,
		City:        request.City,
		Addres:      request.Addres,
		Year:        request.Year,
		Area:        request.Area,
		Description: request.Description,
		Price:       request.Price,
		Tor:         request.Tor,
		Bedroom:     request.Bedroom,
		Bathroom:    request.Bathroom,
		Image:       resp.SecureURL,
	}

	data, err := h.HouseRepository.CreateHouse(house)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	house, _ = h.HouseRepository.GetHouse(house.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
func (h *handlerHouse) UpdateHouse(c echo.Context) error {
	// request := new(tripdto.UpdateTripRequest)
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

	// amenitiesid, _ := strconv.Atoi(c.FormValue("amenitiesid"))
	tor, _ := strconv.Atoi(c.FormValue("tor"))
	price, _ := strconv.Atoi(c.FormValue("price"))
	bedroom, _ := strconv.Atoi(c.FormValue("bedroom"))
	bathroom, _ := strconv.Atoi(c.FormValue("bathroom"))

	request := housedto.UpdateHouseRequest{
		Nameproperty: c.FormValue("nameproperty"),
		// AmenitiesID:  amenitiesid,
		City:     c.FormValue("city"),
		Addres:   c.FormValue("addres"),
		Price:    price,
		Tor:      tor,
		Bedroom:  bedroom,
		Bathroom: bathroom,
		Image:    resp.SecureURL,
	}

	id, _ := strconv.Atoi(c.Param("id"))

	house, err := h.HouseRepository.GetHouse(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Nameproperty != "" {
		house.Nameproperty = request.Nameproperty
	}
	// if request.AmenitiesID != 0 {
	// 	house.AmenitiesID = request.AmenitiesID
	// }
	if request.City != "" {
		house.City = request.City
	}
	if request.Addres != "" {
		house.Addres = request.Addres
	}

	if request.Price != 0 {
		house.Price = request.Price
	}
	if request.Tor != 0 {
		house.Tor = request.Tor
	}
	if request.Bedroom != 0 {
		house.Bedroom = request.Bedroom
	}
	if request.Bathroom != 0 {
		house.Bathroom = request.Bathroom
	}
	if request.Image != "" {
		house.Image = request.Image
	}

	// dataAmenities, _ := h.HouseRepository.GetAmenitieshouse(house.AmenitiesID)

	data, err := h.HouseRepository.UpdateHouse(house, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	// data.Amenities = dataAmenities

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponsehouse(data)})
}
func (h *handlerHouse) DeleteHouse(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	house, err := h.HouseRepository.GetHouse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.HouseRepository.DeleteHouse(house)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponsehouse(data)})
}

// func (h *handlerTrip) UpdateFullcounter(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	trip, err := h.HouseRepository.GetTrip(int(id))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
// 	}
// 	 // Mengambil counterqty dari tabel transaction

// 	trip.Fullcounter = trip.Fullcounter + counterQty

// 	data, err := h.HouseRepository.UpdateFullcounter(trip)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
// 	}

//		return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
//	}

func convertResponsehouse(u models.House) models.House {
	return models.House{
		ID:           u.ID,
		Nameproperty: u.Nameproperty,
		// CountryID:      u.CountryID,
		Amenities: u.Amenities,
		City:      u.City,
		Addres:    u.Addres,
		Year:      u.Year,
		Area:      u.Area,
		Price:     u.Price,
		Tor:       u.Tor,
		Bedroom:   u.Bedroom,
		Bathroom:  u.Bathroom,
		Image:     u.Image,
	}
}
