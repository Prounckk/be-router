package model

import (
	"be-router/entity"
	"errors"
	"github.com/shopspring/decimal"
	"strconv"
)

const (
	Booked     string = "booked"
	Available         = "available"
	OutOfOrder        = "out of order"
	Paid              = "paid"
)

const (
	ErrorCode string = "3434"
)

type ParkingGroup struct {
	ParkingGroup entity.ParkingGroup
}

var ParkingGroupP = entity.ParkingGroup{
	GroupLetter: "P",
	ParkingPlaces: []entity.Parking{
		{
			ID:        231,
			Status:    Paid,
			StartTime: "7:00AM",
			EndTime:   "10:50PM",
		},
		{
			ID:     002,
			Status: Available,
		},
		{
			ID:         003,
			Status:     Available,
			Customized: true,
			CustomParkingPrice: entity.ParkingPrice{
				PaidParkingStart:       "8:00AM",
				PaidParkingEnd:         "8:00PM",
				ProhibitedParkingStart: "5:00AM",
				ProhibitedParkingEnd:   "7:59AM",
				Limit:                  1,
				Rate:                   decimal.New(12.00, 0),
			},
		},
	},
	Price: entity.ParkingPrice{
		PaidParkingStart:       "7:00AM",
		PaidParkingEnd:         "6:00PM",
		ProhibitedParkingStart: "1:00AM",
		ProhibitedParkingEnd:   "3:00AM",
		Limit:                  2,
		Rate:                   decimal.New(14, 0),
	},
}

var ParkingGroupB = entity.ParkingGroup{
	GroupLetter: "B",
	ParkingPlaces: []entity.Parking{
		{
			ID:     233,
			Status: OutOfOrder,
		},
		{
			ID:        3444334,
			Status:    Booked,
			StartTime: "6:00PM",
			EndTime:   "10:00PM",
		},
		{
			ID:         003,
			Status:     Available,
			Customized: true,
			CustomParkingPrice: entity.ParkingPrice{
				PaidParkingStart:       "8:00AM",
				PaidParkingEnd:         "8:00PM",
				ProhibitedParkingStart: "2:00AM",
				ProhibitedParkingEnd:   "4:00AM",
				Limit:                  1,
				Rate:                   decimal.New(12.00, 0),
			},
		},
	},
	Price: entity.ParkingPrice{
		PaidParkingStart:       "7:00AM",
		PaidParkingEnd:         "6:00PM",
		ProhibitedParkingStart: "1:00AM",
		ProhibitedParkingEnd:   "3:00AM",
		Limit:                  2,
		Rate:                   decimal.New(11.00, 0),
	},
}

// let's imagine this is our city(yeah, small, I know) with all ParkingGroups
var cityParkingGroups = []entity.ParkingGroup{
	ParkingGroupB,
	ParkingGroupP,
}

func GetParkingByName(s string) (entity.Parking, error) {
	// fail first, if the string is empty
	if s == "" {
		return entity.Parking{}, errors.New("empty parking name provided, code: " + ErrorCode)
	}
	// create a new instance of the parking group to interact with
	pg := ParkingGroup{}
	p := entity.Parking{}
	ParkingGroup := pg.findParkingByGroupLetter(s[0:1])
	spotId, err := strconv.Atoi(s[1:])
	if err != nil {
		return entity.Parking{}, err
	}
	for _, spot := range ParkingGroup.ParkingPlaces {
		if spot.ID == spotId {
			p = spot
			if p.Status == OutOfOrder {
				return p, nil
			}
			if p.Customized == false {
				p.CustomParkingPrice = ParkingGroup.Price
			}
			return p, nil
		}
	}

	return entity.Parking{}, errors.New("Can not find the parking with provided ID: " + s + ", code: " + ErrorCode)
}

// findParkingByGroupLetter search parking by first provided letter
func (P *ParkingGroup) findParkingByGroupLetter(s string) entity.ParkingGroup {

	for _, group := range cityParkingGroups {
		if group.GroupLetter == s {
			return group
		}
	}
	return entity.ParkingGroup{}
}

// findParkingByStatus search parking by specific status Ex: booked or available
// return an empty slice if wrong status or no available parking spots for this status
func (P *ParkingGroup) findParkingByStatus(s string) []entity.Parking {

	var available []entity.Parking
	for _, group := range cityParkingGroups {
		for _, spots := range group.ParkingPlaces {
			if spots.Status == s {
				available = append(available, spots)
			}
		}
	}
	return available
}

func (P *ParkingGroup) findParkingPriceByLetter(s string) entity.ParkingPrice {

	p := ParkingGroup{}
	ParkingGroup := p.findParkingByGroupLetter(s)
	return ParkingGroup.Price

}
