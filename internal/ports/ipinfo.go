package ports

import "net"

type Info struct {
	Continent Continent `json:"continent"`
	Country   Country   `json:"country"`
	City      City      `json:"city"`
	ASN       ASN       `json:"asn"`
}

type IIPInfo interface {
	GetInfo(ip net.IP) (Info, error)
}
