package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQuotesHandler_Success(t *testing.T) {
	fetcher = func() ([]Quote, error) {
		return []Quote{
			{Quote: "Test Quote", Author: "Test Author"},
		}, nil
	}

	req, err := http.NewRequest("GET", "/quotes", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(QuotesHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var quotes []Quote
	err = json.Unmarshal(rr.Body.Bytes(), &quotes)
	if err != nil {
		t.Fatal(err)
	}

	if len(quotes) != 1 || quotes[0].Quote != "Test Quote" || quotes[0].Author != "Test Author" {
		t.Errorf("Handler returned unexpected body: got %v", rr.Body.String())
	}
}

func TestQuotesHandler_FetchError(t *testing.T) {
	fetcher = func() ([]Quote, error) {
		return nil, fmt.Errorf("mock error")
	}

	req, err := http.NewRequest("GET", "/quotes", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(QuotesHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}
