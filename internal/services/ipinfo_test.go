package services

import (
	"github.com/red-life/ipinfo/internal/mocks"
	"github.com/red-life/ipinfo/internal/ports"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestIPInfo_GetInfo(t *testing.T) {
	maxMind := mocks.IMaxMind{}
	ip := net.ParseIP("185.143.233.200")
	expectedInfo := ports.Info{
		IP: "185.143.233.200",
		Continent: ports.Continent{
			Code: "AS",
			Name: "Asia",
		},
		Country: ports.Country{
			ISOCode: "IR",
			Name:    "Iran",
		},
		City: ports.City{
			Latitude:  35.698,
			Longitude: 51.4115,
			TimeZone:  "Asia/Tehran",
		},
		ASN: ports.ASN{
			Number:       205585,
			Organization: "Noyan Abr Arvan Co. ( Private Joint Stock)",
		},
	}
	maxMind.EXPECT().GetContinent(ip).Return(expectedInfo.Continent, nil)
	maxMind.EXPECT().GetCountry(ip).Return(expectedInfo.Country, nil)
	maxMind.EXPECT().GetCity(ip).Return(expectedInfo.City, nil)
	maxMind.EXPECT().GetASN(ip).Return(expectedInfo.ASN, nil)
	ipInfo := NewIPInfo(&maxMind)
	_, err := ipInfo.Info(net.ParseIP("2400:6500:ff00::7af8:f305"), true, true, true, true)
	assert.NotNil(t, err, "Expected to return an error on IPv6")
	info, err := ipInfo.Info(ip, true, true, true, true)
	assert.Nil(t, err, "Expected to return nil on getting info")
	assert.Equal(t, expectedInfo, info)
}
