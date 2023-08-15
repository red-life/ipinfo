package ports

import (
	"github.com/oschwald/maxminddb-golang"
	"net"
)

type Continent struct {
	Code string
	Name string
}

type Country struct {
	ISOCode string
	Name    string
}

type City struct {
	Latitude  float32
	Longitude float32
	TimeZone  string
}

type ASN struct {
	Number       uint
	Organization string
}

type IMaxMind interface {
	GetContinent(ip net.IP) (Continent, error)
	GetCountry(ip net.IP) (Country, error)
	GetCity(ip net.IP) (City, error)
	GetASN(ip net.IP) (ASN, error)
	Load(countryReader *maxminddb.Reader, cityReader *maxminddb.Reader, asnReader *maxminddb.Reader)
}
