package service

import (
	"log"
	"os"
	"strconv"

	"github.com/drone-navigation-system/models"
)

type DNSService interface {
	GetLocation(details models.Request) models.Response
}

type dnsService struct {
}

//NewDNSService constructor method
func NewDNSService() DNSService {
	return &dnsService{}
}

//GetLocation calls GetLocation function
func (s *dnsService) GetLocation(details models.Request) models.Response {

	secID := convertStringToFloat64(os.Getenv("SECTOR_ID"))
	loc := (convertStringToFloat64(details.X) * secID) + (convertStringToFloat64(details.Y) * secID) + (convertStringToFloat64(details.Z) * secID) + convertStringToFloat64(details.Vel)
	return models.Response{loc}
}

func convertStringToFloat64(str string) float64 {

	s, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Println(err.Error())
	}
	return s
}
