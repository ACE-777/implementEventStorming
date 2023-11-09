package clasess

import (
	"errors"
	"log"
)

type Review interface {
	MakeReview() error
}

type Reviewing struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Review  string `json:"review"`
}

func (r *Reviewing) MakeReview() error {
	//in this way we can check the new review
	//verifying review
	if len(r.Review) < 1 {
		log.Printf("invalid review")
		return errors.New("invalid review")
	}

	log.Printf("Success review!")

	return nil
}
