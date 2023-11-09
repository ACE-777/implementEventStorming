package server

import (
	"log"
	"net/http"
)

func RunServer() {
	http.HandleFunc("/coloring/home", home)
	http.HandleFunc("/coloring/list", list)
	http.HandleFunc("/coloring/book", book)
	http.HandleFunc("/coloring/payment", payment)
	http.HandleFunc("/coloring/review", review)

	err := http.ListenAndServe("localhost:6060", nil)
	if err != nil {
		log.Printf("error in startup server %v", err)
	}

}
