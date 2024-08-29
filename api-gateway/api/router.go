package api

import (
	"api-gateway/service"

	"api-gateway/api/handler"

	"github.com/gin-gonic/gin"

	_ "api-gateway/cmd/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // swagger middleware for gin
)

type Options struct {
	Service service.IClients
}

func Router(opt Options) *gin.Engine {
	r := gin.Default()

	h := handler.NewHandler(&handler.HandlerCinfig{
		Service: opt.Service,
	})

	api := r.Group("/api")
	{
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		api.POST("/users", h.RegisterUser)
		api.POST("/users/login", h.LoginUser)

		// Booking service routes
		api.POST("/bookings", h.CreateBooking)
		api.GET("/bookings/:id", h.GetBookingDetails)
		api.PUT("/bookings/:id", h.UpdateBooking)
		api.DELETE("/bookings/:id", h.CancelBooking)
		api.GET("/users/:user_id/bookings", h.ListUserBookings)

		// Hotel service routes
		api.POST("/hotels", h.AddHotel)
		api.GET("/hotels", h.GetHotels)
		api.GET("/hotels/:id", h.GetHotelDetails)
		api.GET("/hotels/:id/rooms/availability", h.CheckRoomAvailability)
	}

	return r
}
