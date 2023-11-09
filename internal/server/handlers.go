package server

import (
	"encoding/json"
	"fmt"
	"implementEventStorming/internal/clasess"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(fmt.Sprintf("method must be PUT")))
		return
	}

	_, err := w.Write([]byte("welcome to airbnb tools!"))
	if err != nil {
		log.Printf("error in loading home page %v", err)
	}

	// here we need to save new user by incoming parameters from login to home page
	//vasua and vasiliev- just the simplest example
	user := NewUser("vasua", "vasuliev", 1, 0)
	user.saveToDB()

	w.WriteHeader(http.StatusOK)
}

func list(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(fmt.Sprintf("method must be PUT")))
		return
	}

	decode := json.NewDecoder(r.Body)
	decode.DisallowUnknownFields()

	var user clasess.Listing
	if err := decode.Decode(&user); err != nil {
		//request declined
		log.Printf("Request declined via invalid list parameters %v", err)
		_, err = w.Write([]byte("Request declined via invalid user parameters"))
		if err != nil {
			log.Printf("error in loading list page %v", err)
		}

		w.WriteHeader(http.StatusBadGateway)
	}

	if err := user.CreateList(); err != nil {
		//request declined
		sendBadNotification(err.Error(), w)
		log.Printf("User can't create list with invalid parametres %v", err)

		return
	}

	//request approved
	if err := sendGoodNotification(w); err != nil {
		log.Printf("Can't send good notify to user")
	}

	log.Printf("User create new list %v", user.IDUser)

	w.WriteHeader(http.StatusOK)

	return
}

func book(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(fmt.Sprintf("method must be PUT")))
		return
	}

	decode := json.NewDecoder(r.Body)
	decode.DisallowUnknownFields()

	var user clasess.Booking
	if err := decode.Decode(&user); err != nil {
		//request declined
		log.Printf("Request declined via invalid booking parameters %v", err)
		_, err = w.Write([]byte("Request declined via invalid user parameters"))
		if err != nil {
			log.Printf("error in loading list page %v", err)
		}

		w.WriteHeader(http.StatusBadGateway)
	}

	if err := user.Book(); err != nil {
		//request declined
		sendBadNotification(err.Error(), w)
		log.Printf("User can't create booking with invalid parametres %v", err)

		return
	}

	//request approved
	if err := sendGoodNotification(w); err != nil {
		log.Printf("Can't send good notify to user")
	}

	log.Printf("User make new booking %v", user.IDUser)
	w.WriteHeader(http.StatusOK)

	return
}

func payment(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(fmt.Sprintf("method must be PUT")))
		return
	}

	decode := json.NewDecoder(r.Body)
	decode.DisallowUnknownFields()

	var user clasess.Payment
	if err := decode.Decode(&user); err != nil {
		//request declined
		log.Printf("Request declined via invalid payment parameters %v", err)
		_, err = w.Write([]byte("Request declined via invalid user parameters"))
		if err != nil {
			log.Printf("error in loading list page %v", err)
		}

		w.WriteHeader(http.StatusBadGateway)
	}

	if err := user.MakePayment(); err != nil {
		//request declined
		sendBadNotification(err.Error(), w)
		log.Printf("User can't create payment with invalid parametres %v", err)

		return
	}

	//request approved
	if err := sendGoodNotification(w); err != nil {
		log.Printf("Can't send good notify to user")
	}

	log.Printf("User make new payment %v", user.ID)
	w.WriteHeader(http.StatusOK)

	return
}

func review(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(fmt.Sprintf("method must be PUT")))
		return
	}

	decode := json.NewDecoder(r.Body)
	decode.DisallowUnknownFields()

	var user clasess.Reviewing
	if err := decode.Decode(&user); err != nil {
		//request declined
		log.Printf("Request declined via invalid review parameters %v", err)
		_, err = w.Write([]byte("Request declined via invalid user parameters"))
		if err != nil {
			log.Printf("error in loading list page %v", err)
		}

		w.WriteHeader(http.StatusBadGateway)
	}

	if err := user.MakeReview(); err != nil {
		//request declined
		sendBadNotification(err.Error(), w)
		log.Printf("User can't create review with invalid parametres %v", err)

		return
	}

	//request approved
	if err := sendGoodNotification(w); err != nil {
		log.Printf("Can't send good notify to user")
	}

	log.Printf("User make new review %v", user.ID)
	w.WriteHeader(http.StatusOK)

	return
}
