package clasess

import (
	"errors"
	"log"
)

type Pay interface {
	MakePayment() error
}

type Payment struct {
	ID                int64   `json:"id"`
	GUIDAccommodation int64   `json:"guid_accommodation"`
	Amount            float64 `json:"amount"`
}

func (p *Payment) MakePayment() error {
	//in this way we can check the new payment
	//verifying payment
	if p.Amount < 0 {
		log.Printf("can not withadrow")
		return errors.New("invalid count of money, cant withdraw")
	}

	log.Printf("Success withadrow")
	return nil
}
