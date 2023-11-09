package clasess

import (
	"errors"
	"log"
)

type Book interface {
	Book() error
}

type Booking struct {
	GUIDAccommodation int64  `json:"guid_accommodation"`
	ID                int64  `json:"id"`
	IDUser            int64  `json:"id_user"`
	StartTime         uint64 `json:"start_time"`
	EndTime           uint64 `json:"end_time"`
	CountGuests       int    `json:"count_guests"`
	CountNights       int    `json:"count_nights"`
	Money             int    `json:"money"`
}

func (b *Booking) Book() error {
	//verifying new booking
	if b.StartTime > b.EndTime {
		return errors.New("invalid booking")
	}

	//proceedPayment
	if err := b.MakePayment(); err != nil {
		return err
	}

	log.Printf("Valid new booking")

	return nil
}

func (b *Booking) MakePayment() error {
	//in this way we can check the new payment
	//verifying payment
	if b.Money < 0 {
		log.Printf("can not withadrow")
		return errors.New("invalid count of money, cant withdraw")
	}

	log.Printf("Success withadrow")
	return nil
}
