package services

import (
	"github.com/red-life/ipinfo/internal/pkg/customerror"
	"github.com/red-life/ipinfo/internal/ports"
	"net"
)

func NewIPInfo(maxMind ports.IMaxMind) ports.IIPInfo {
	return &IPInfo{
		mm: maxMind,
	}
}

type IPInfo struct {
	mm ports.IMaxMind
}

func (I *IPInfo) GetInfo(ip net.IP) (ports.Info, error) {
	if ip.To4() == nil {
		return ports.Info{}, customerror.NotSupportedErr
	}
	continent, err := I.mm.GetContinent(ip)
	if err != nil {
		return ports.Info{}, customerror.MaxMindToCustom(err)
	}
	country, err := I.mm.GetCountry(ip)
	if err != nil {
		return ports.Info{}, customerror.MaxMindToCustom(err)
	}
	city, err := I.mm.GetCity(ip)
	if err != nil {
		return ports.Info{}, customerror.MaxMindToCustom(err)
	}
	asn, err := I.mm.GetASN(ip)
	if err != nil {
		return ports.Info{}, customerror.MaxMindToCustom(err)
	}
	return ports.Info{
		Continent: continent,
		Country:   country,
		City:      city,
		ASN:       asn,
	}, nil
}
