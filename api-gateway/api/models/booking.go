package models

type Booking struct {
	UserID       string  `json:"userID"`
	HotelID      string  `json:"hotelID"`
	RoomType     string  `json:"roomType"`
	CheckInDate  string  `json:"checkInDate"`
	CheckOutDate string  `json:"checkOutDate"`
	TotalAmount  float64 `json:"totalAmounts`
}

type UpdateBooking struct {
	BookingID    string  `json:"bookingID"`
	CheckInDate  string  `json:"checkInDate"`
	CheckOutDate string  `json:"checkOutDate"`
	TotalAmount  float64 `json:"totalAmount"`
	Status       string  `json:"status"`
}
