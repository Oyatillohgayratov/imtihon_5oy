package models

type Room struct {
	RoomType      string  `json:"roomType"`
	PricePerNight float64 `json:"pricePerNight"`
	Availability  bool    `json:"availability"`
}

type Hotel struct {
	HotelID  string  `json:"hotelID"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Rating   float64 `json:"rating"`
	Address  string  `json:"address"`
	Rooms    []Room  `json:"rooms"`
}
