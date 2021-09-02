package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/drone-navigation-system/models"
	"github.com/drone-navigation-system/service"
)

type DNSController interface {
	Swagger(w http.ResponseWriter, r *http.Request)
	Ping(w http.ResponseWriter, r *http.Request)
	GetLocation(w http.ResponseWriter, req *http.Request)
}

type dnsController struct {
	Srv service.DNSService
}

//NewDNSController constructor function
func NewDNSController(srv service.DNSService) DNSController {
	s := &dnsController{srv}
	return s
}

//Swagger serves up the Swagger definition file
func (s *dnsController) Swagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "swagger/swagger.yaml")
}

//Ping returns a response if the drone is alive
func (s *dnsController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

//GetLocation initiates the GetLocation handler
func (s *dnsController) GetLocation(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(req.Body)
	var e models.Request
	if err := decoder.Decode(&e); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(s.Srv.GetLocation(e))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}
