package ports

import (
	"github.com/oschwald/maxminddb-golang"
	"net"
)

type Continent struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Country struct {
	ISOCode string `json:"iso_code"`
	Name    string `json:"name"`
}

type City struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	TimeZone  string  `json:"time_zone"`
}

type ASN struct {
	Number       uint   `json:"number"`
	Organization string `json:"organization"`
}

type IMaxMind interface {
	GetContinent(ip net.IP) (Continent, error)
	GetCountry(ip net.IP) (Country, error)
	GetCity(ip net.IP) (City, error)
	GetASN(ip net.IP) (ASN, error)
	Load(countryReader *maxminddb.Reader, cityReader *maxminddb.Reader, asnReader *maxminddb.Reader)
}
