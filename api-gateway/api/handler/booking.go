package handler

import (
	"api-gateway/api/models"
	booking "api-gateway/protos/booking"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBooking godoc
// @Summary Create a new booking
// @Description Create a new booking for a user with specified details
// @Tags bookings
// @Accept json
// @Produce json
// @Param booking body models.Booking true "Booking Data"
// @Success 200 {object} booking.CreateBookingResponse
// @Failure 400 {object} gin.H{"error": "Invalid request body"}
// @Failure 500 {object} gin.H{"error": "Failed to create booking"}
// @Router /api/bookings [post]
func (h *Handler) CreateBooking(c *gin.Context) {
	var req models.Booking
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := h.service.Booking().CreateBooking(context.Background(), &booking.CreateBookingRequest{
		UserID:       req.UserID,
		HotelID:      req.HotelID,
		RoomType:     req.RoomType,
		CheckInDate:  req.CheckInDate,
		CheckOutDate: req.CheckOutDate,
		TotalAmount:  req.TotalAmount,
	})

	if err != nil {
		fmt.Println("Error creating booking:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetBookingDetails godoc
// @Summary Get details of a specific booking
// @Description Retrieve detailed information about a specific booking by its ID
// @Tags bookings
// @Produce json
// @Param id path string true "Booking ID"
// @Success 200 {object} booking.GetBookingDetailsResponse
// @Failure 500 {object} gin.H{"error": "Failed to get booking details"}
// @Router /api/bookings/{id} [get]
func (h *Handler) GetBookingDetails(c *gin.Context) {
	bookingID := c.Param("id")
	resp, err := h.service.Booking().GetBookingDetails(context.Background(), &booking.GetBookingDetailsRequest{BookingID: bookingID})
	if err != nil {
		fmt.Println("Error getting booking details:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get booking details"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateBooking godoc
// @Summary Update an existing booking
// @Description Update the details of an existing booking by its ID
// @Tags bookings
// @Accept json
// @Produce json
// @Param booking body models.UpdateBooking true "Update Booking Data"
// @Success 200 {object} booking.UpdateBookingResponse
// @Failure 400 {object} gin.H{"error": "Invalid request body"}
// @Failure 500 {object} gin.H{"error": "Failed to update booking"}
// @Router /api/bookings [put]
func (h *Handler) UpdateBooking(c *gin.Context) {
	var req models.UpdateBooking
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := h.service.Booking().UpdateBooking(context.Background(), &booking.UpdateBookingRequest{
		BookingID:    req.BookingID,
		CheckInDate:  req.CheckInDate,
		CheckOutDate: req.CheckOutDate,
		TotalAmount:  req.TotalAmount,
		Status:       req.Status,
	})
	if err != nil {
		fmt.Println("Error updating booking:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CancelBooking godoc
// @Summary Cancel a booking
// @Description Cancel an existing booking by its ID
// @Tags bookings
// @Produce json
// @Param id path string true "Booking ID"
// @Success 200 {object} booking.CancelBookingResponse
// @Failure 500 {object} gin.H{"error": "Failed to cancel booking"}
// @Router /api/bookings/{id} [delete]
func (h *Handler) CancelBooking(c *gin.Context) {
	bookingID := c.Param("id")
	resp, err := h.service.Booking().CancelBooking(context.Background(), &booking.CancelBookingRequest{BookingID: bookingID})
	if err != nil {
		fmt.Println("Error canceling booking:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel booking"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ListUserBookings godoc
// @Summary List all bookings for a user
// @Description Retrieve a list of all bookings for a specific user by their ID
// @Tags bookings
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {array} booking.ListUserBookingsResponse
// @Failure 500 {object} gin.H{"error": "Failed to list user bookings"}
// @Router /api/users/{user_id}/bookings [get]
func (h *Handler) ListUserBookings(c *gin.Context) {
	userID := c.Param("user_id")
	resp, err := h.service.Booking().ListUserBookings(context.Background(), &booking.ListUserBookingsRequest{UserID: userID})
	if err != nil {
		fmt.Println("Error listing user bookings:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list user bookings"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
