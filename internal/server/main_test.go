package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckHomeHandlerSuccessfulMethod(t *testing.T) {
	requestBody := []byte(`{"id":1,"name":"Vasya","surname":"Vasilev","experience":1}`)
	req, err := http.NewRequest(http.MethodPut, "/coloring/home", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	home(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedResponseBody := "welcome to airbnb tools!"

	if w.Body.String() != expectedResponseBody {
		t.Errorf("Expected response body '%s', got '%s'", expectedResponseBody, w.Body.String())
	}
}

func TestCheckListHandlerSuccessfulMethod(t *testing.T) {
	requestBody := []byte(`{"id_user":1,"accommodation_type":"apartments","available_capacity":6,"address":"MKN",
	"country":"Russia","city":"Saint-Petersburg","guid":1,"price_per_night":100.45}`)
	req, err := http.NewRequest(http.MethodPut, "/coloring/list", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	list(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedResponseBody := "You operation was succeed!"

	if w.Body.String() != expectedResponseBody {
		t.Errorf("Expected response body '%s', got '%s'", expectedResponseBody, w.Body.String())
	}
}

func TestCheckBookHandlerSuccessfulMethod(t *testing.T) {
	requestBody := []byte(`{"guid_accommodation":1,"id":1,"id_user":1,"start_time":1000,
	"end_time":2000,"count_guests":4,"count_nights":1,"money":1000}`)
	req, err := http.NewRequest(http.MethodPut, "/coloring/book", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	book(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedResponseBody := "You operation was succeed!"

	if w.Body.String() != expectedResponseBody {
		t.Errorf("Expected response body '%s', got '%s'", expectedResponseBody, w.Body.String())
	}
}

func TestCheckPaymentHandlerSuccessfulMethod(t *testing.T) {
	requestBody := []byte(`{"id":1,"guid_accommodation":1,"amount":1000.5}`)
	req, err := http.NewRequest(http.MethodPut, "/coloring/payment", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	payment(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedResponseBody := "You operation was succeed!"

	if w.Body.String() != expectedResponseBody {
		t.Errorf("Expected response body '%s', got '%s'", expectedResponseBody, w.Body.String())
	}
}

func TestCheckReviewHandlerSuccessfulMethod(t *testing.T) {
	requestBody := []byte(`{"id":1,"name":"Misha","surname":"Diagilev","review":"good apparts!"}`)
	req, err := http.NewRequest(http.MethodPut, "/coloring/review", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	review(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedResponseBody := "You operation was succeed!"

	if w.Body.String() != expectedResponseBody {
		t.Errorf("Expected response body '%s', got '%s'", expectedResponseBody, w.Body.String())
	}
}
