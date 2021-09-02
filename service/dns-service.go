package service

import (
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

	// secID := os.Getenv("SECTOR_ID")
	secID := 5.0
	loc := (details.X * secID) + (details.Y * secID) + (details.Z * secID) + details.Vel
	return models.Response{loc}
}
