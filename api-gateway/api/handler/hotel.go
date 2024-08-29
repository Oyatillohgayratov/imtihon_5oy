package handler

import (
	"api-gateway/api/models"
	hotel "api-gateway/protos/hotel"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddHotel godoc
// @Summary Add a new hotel
// @Description Add a new hotel with its details and rooms
// @Tags hotels
// @Accept json
// @Produce json
// @Param hotel body models.Hotel true "Hotel Data"
// @Success 200 {object} hotel.Hotel
// @Failure 400 {object} gin.H{"error": "Invalid request body"}
// @Failure 500 {object} gin.H{"error": "Failed to add hotel"}
// @Router /api/hotels [post]
func (h *Handler) AddHotel(c *gin.Context) {
	var req models.Hotel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	var hotels []*hotel.Room
	for _, room := range req.Rooms {
		hotels = append(hotels, &hotel.Room{
			RoomType:      room.RoomType,
			PricePerNight: room.PricePerNight,
			Availability:  room.Availability,
		})
	}

	resp, err := h.service.Hotel().AddHotel(context.Background(), &hotel.Hotel{
		HotelID:  req.HotelID,
		Name:     req.Name,
		Location: req.Location,
		Rating:   req.Rating,
		Address:  req.Address,
		Rooms:    hotels,
	})
	if err != nil {
		fmt.Println("Error adding hotel:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add hotel"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetHotels godoc
// @Summary Get a list of all hotels
// @Description Retrieve a list of all hotels
// @Tags hotels
// @Produce json
// @Success 200 {array} hotel.Hotel
// @Failure 500 {object} gin.H{"error": "Failed to get hotels"}
// @Router /api/hotels [get]
func (h *Handler) GetHotels(c *gin.Context) {
	resp, err := h.service.Hotel().GetHotels(context.Background(), &hotel.GetHotelsRequest{})
	if err != nil {
		fmt.Println("Error getting hotels:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get hotels"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetHotelDetails godoc
// @Summary Get details of a specific hotel
// @Description Retrieve details of a hotel by its ID
// @Tags hotels
// @Produce json
// @Param id path string true "Hotel ID"
// @Success 200 {object} hotel.Hotel
// @Failure 400 {object} gin.H{"error": "Invalid hotel ID"}
// @Failure 500 {object} gin.H{"error": "Failed to get hotel details"}
// @Router /api/hotels/{id} [get]
func (h *Handler) GetHotelDetails(c *gin.Context) {
	hotelID := c.Param("id")
	resp, err := h.service.Hotel().GetHotelDetails(context.Background(), &hotel.GetHotelDetailsRequest{HotelID: hotelID})
	if err != nil {
		fmt.Println("Error getting hotel details:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get hotel details"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CheckRoomAvailability godoc
// @Summary Check availability of rooms in a hotel
// @Description Check the availability of rooms in a hotel by its ID
// @Tags hotels
// @Produce json
// @Param id path string true "Hotel ID"
// @Success 200 {array} hotel.Room
// @Failure 400 {object} gin.H{"error": "Invalid hotel ID"}
// @Failure 500 {object} gin.H{"error": "Failed to check room availability"}
// @Router /api/hotels/{id}/rooms/availability [get]
func (h *Handler) CheckRoomAvailability(c *gin.Context) {
	hotelID := c.Param("id")
	resp, err := h.service.Hotel().CheckRoomAvailability(context.Background(), &hotel.CheckRoomAvailabilityRequest{HotelID: hotelID})
	if err != nil {
		fmt.Println("Error checking room availability:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check room availability"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
