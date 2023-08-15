package maxmind

import (
	"github.com/oschwald/maxminddb-golang"
	"github.com/red-life/ip-info/internal/pkg/customerror"
	"github.com/red-life/ip-info/internal/ports"
	"net"
)

func NewMaxMind(countryReader *maxminddb.Reader, cityReader *maxminddb.Reader, asnReader *maxminddb.Reader) ports.IMaxMind {
	return &MaxMind{
		countryReader: countryReader,
		cityReader:    cityReader,
		asnReader:     asnReader,
	}
}

type ContinentRecord struct {
	Continent struct {
		Code  string `maxminddb:"code"`
		Names struct {
			En string `maxminddb:"en"`
		} `maxminddb:"names"`
	} `maxminddb:"continent"`
}

type CountryRecord struct {
	RegisteredCountry struct {
		ISOCode string `maxminddb:"iso_code"`
		Names   struct {
			En string `maxminddb:"en"`
		} `maxminddb:"names"`
	} `maxminddb:"registered_country"`
}

type CityRecord struct {
	Location struct {
		Latitude  float32 `maxminddb:"latitude"`
		Longitude float32 `maxminddb:"longitude"`
		TimeZone  string  `maxminddb:"time_zone"`
	} `maxminddb:"location"`
}
type ASNRecord struct {
	Number       uint   `maxminddb:"autonomous_system_number"`
	Organization string `maxminddb:"autonomous_system_organization"`
}

type MaxMind struct {
	countryReader *maxminddb.Reader
	cityReader    *maxminddb.Reader
	asnReader     *maxminddb.Reader
}

func (m *MaxMind) GetContinent(ip net.IP) (ports.Continent, error) {
	var record ContinentRecord
	err := m.cityReader.Lookup(ip, &record)
	if err != nil {
		return ports.Continent{}, err
	}
	if record.Continent.Code == "" || record.Continent.Names.En == "" {
		return ports.Continent{}, customerror.NotFound
	}
	return ports.Continent{
		Code: record.Continent.Code,
		Name: record.Continent.Names.En,
	}, nil
}

func (m *MaxMind) GetCountry(ip net.IP) (ports.Country, error) {
	var record CountryRecord
	err := m.countryReader.Lookup(ip, &record)
	if err != nil {
		return ports.Country{}, err
	}
	if record.RegisteredCountry.ISOCode == "" || record.RegisteredCountry.Names.En == "" {
		return ports.Country{}, customerror.NotFound
	}
	return ports.Country{
		ISOCode: record.RegisteredCountry.ISOCode,
		Name:    record.RegisteredCountry.Names.En,
	}, nil
}

func (m *MaxMind) GetCity(ip net.IP) (ports.City, error) {
	var record CityRecord
	err := m.cityReader.Lookup(ip, &record)
	if err != nil {
		return ports.City{}, err
	}
	if record.Location.TimeZone == "" || record.Location.Latitude == 0 || record.Location.Longitude == 0 {
		return ports.City{}, customerror.NotFound
	}
	return ports.City{
		Latitude:  record.Location.Latitude,
		Longitude: record.Location.Longitude,
		TimeZone:  record.Location.TimeZone,
	}, nil
}

func (m *MaxMind) GetASN(ip net.IP) (ports.ASN, error) {
	var record ASNRecord
	err := m.asnReader.Lookup(ip, &record)
	if err != nil {
		return ports.ASN{}, err
	}
	if record.Organization == "" || record.Number == 0 {
		return ports.ASN{}, customerror.NotFound
	}
	return ports.ASN{
		Organization: record.Organization,
		Number:       record.Number,
	}, nil
}

func (m *MaxMind) Load(countryReader *maxminddb.Reader, cityReader *maxminddb.Reader, asnReader *maxminddb.Reader) {
	m.countryReader = countryReader
	m.cityReader = cityReader
	m.asnReader = asnReader
}
