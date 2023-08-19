package ports

import "net"

type Info struct {
	IP        string    `json:"ip"`
	Continent Continent `json:"continent,omitempty"`
	Country   Country   `json:"country,omitempty"`
	City      City      `json:"city,omitempty"`
	ASN       ASN       `json:"asn,omitempty"`
}

type ShortInfo struct {
	IP           string `json:"ip"`
	Country      string `json:"country"`
	Continent    string `json:"continent"`
	Location     string `json:"location"`
	Organization string `json:"org"`
	TimeZone     string `json:"time_zone"`
}

type IIPInfo interface {
	Info(ip net.IP, continent bool, country bool, city bool, asn bool) (Info, error)
	ShortInfo(ip net.IP) (ShortInfo, error)
}
