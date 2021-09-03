package service

import (
	"os"
	"testing"

	"github.com/drone-navigation-system/models"
	"github.com/stretchr/testify/assert"
)

func TestDNSService_GetLocation(t *testing.T) {

	os.Setenv("SECTOR_ID", "5.0")

	req := models.Request{X: "534.0", Y: "77.00", Z: "898.0", Vel: "66.6"}
	resp := models.Response{Location: 7611.6}

	testService := NewDNSService()

	result := testService.GetLocation(req)

	assert.Equal(t, resp, result)

}
