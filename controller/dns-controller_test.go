package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drone-navigation-system/service"
)

func TestDNSController_Ping(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ping", nil)

	dnsService := service.NewDNSService()
	dnsController := NewDNSController(dnsService)

	handler := http.HandlerFunc(dnsController.Ping)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong staus code: got %v\n want %v", status, http.StatusOK)
	}
}

func TestDNSController_GetLocations(t *testing.T) {

	jsonReq := []byte(`{"x":"23.00","y":"326.0","z":"78.00","vel":"656.0"}`)

	req, _ := http.NewRequest("POST", "/endpoint", bytes.NewBuffer(jsonReq))

	dnsService := service.NewDNSService()
	dnsController := NewDNSController(dnsService)

	handler := http.HandlerFunc(dnsController.GetLocation)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong staus code: got %v\n want %v", status, http.StatusOK)
	}

}

func TestDNSController_Fail_GetLocations(t *testing.T) {

	jsonReq := []byte(`{"x":23,"y":6546,"z":222,"vel":70}`)

	req, _ := http.NewRequest("POST", "/endpoint", bytes.NewBuffer(jsonReq))

	dnsService := service.NewDNSService()
	dnsController := NewDNSController(dnsService)

	handler := http.HandlerFunc(dnsController.GetLocation)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	status := response.Code

	if status != http.StatusBadRequest {
		t.Errorf("Handler returned a wrong staus code: got %v\n want %v", status, http.StatusBadRequest)
	}

}
