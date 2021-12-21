package entity

import (
	"github.com/shopspring/decimal"
)

// Parking is an individual parking spot belongs to a ParkingGroup
// Each parking spot can have customized pricing, but in general, use ParkingGroup's pricing model
type Parking struct {
	ID                 int          `json:"id"`
	Status             string       `json:"status"` //might be paid, available, reserved, out of order
	StartTime          string       `json:"start"`
	EndTime            string       `json:"end"`
	Customized         bool         `json:"customRules"`
	CustomParkingPrice ParkingPrice `json:"customPrice"`
}

// ParkingGroup is set of parking spots with defined pricing model
// Each ParkingGroup belongs to a Municipality
type ParkingGroup struct {
	GroupLetter   string       `json:"parkingGroupLetter"`
	ParkingPlaces []Parking    `json:"parkingPlaces"`
	Price         ParkingPrice `json:"price"`
}

// ParkingPrice is a pricing model for a group or individual parking spots
// Each ParkingPrice has range of time for free and paid parking,
// limit for how long the place might be taken
// rate that based on dollars per hour
type ParkingPrice struct {
	PaidParkingStart       string          `json:"paidParkingStart"`
	PaidParkingEnd         string          `json:"paidParkingEnd"`
	ProhibitedParkingStart string          `json:"prohibitedParkingStart"`
	ProhibitedParkingEnd   string          `json:"prohibitedParkingEnd"`
	Limit                  int             `json:"limit"`
	Rate                   decimal.Decimal `json:"rate"`
}
