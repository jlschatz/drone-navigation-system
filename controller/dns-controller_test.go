package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drone-navigation-system/service"
)

func TestDNSController_Ping(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ping", nil)

	dnsService := service.NewDNSService()
	dnsController := NewDNSController(dnsService)

	//Controller method needs to the Handler
	handler := http.HandlerFunc(dnsController.Ping)

	//Record HTTP response
	response := httptest.NewRecorder()

	//Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add assertions on the HTTP status code and response
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong staus code: got %v\n want %v", status, http.StatusOK)
	}
}
