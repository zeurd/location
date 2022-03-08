package models

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
	Timestamp int64
}

type LocationHistory struct {
	OrderId string     `json:"order_id"`
	History []Location `json:"history"`
}

func NewLocationHistory(orderId string) *LocationHistory {
	return &LocationHistory{
		orderId,
		make([]Location, 0, 10),
	}
}