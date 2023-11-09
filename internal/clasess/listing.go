package clasess

import (
	"errors"
	"log"
)

type Listi interface {
	CreateList() error
}

type Listing struct {
	IDUser            int64   `json:"id_user"`
	AccommodationType string  `json:"accommodation_type"`
	AvailableCapacity int     `json:"available_capacity"`
	Address           string  `json:"address"`
	Country           string  `json:"country"`
	City              string  `json:"city"`
	GUID              int64   `json:"guid"`
	PricePerNight     float64 `json:"price_per_night"`
}

func (l *Listing) CreateList() error {
	//in this way we can check the new list
	//verifying ticket
	if l.PricePerNight < 0 {
		return errors.New("invalid list")
	}

	log.Printf("Valid new list")

	return nil
}
