package services

import (
	"fmt"
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

func (I *IPInfo) Info(ip net.IP, continent bool, country bool, city bool, asn bool) (ports.Info, error) {
	if ip.To4() == nil {
		return ports.Info{}, customerror.NotSupportedErr
	}
	var info ports.Info
	info.IP = ip.String()
	if continent {
		continentInfo, err := I.mm.GetContinent(ip)
		if err != nil {
			return ports.Info{}, customerror.MaxMindToCustom(err)
		}
		info.Continent = &continentInfo
	}
	if country {
		countryInfo, err := I.mm.GetCountry(ip)
		if err != nil {
			return ports.Info{}, customerror.MaxMindToCustom(err)
		}
		info.Country = &countryInfo
	}
	if city {
		cityInfo, err := I.mm.GetCity(ip)
		if err != nil {
			return ports.Info{}, customerror.MaxMindToCustom(err)
		}
		info.City = &cityInfo
	}
	if asn {
		asnInfo, err := I.mm.GetASN(ip)
		if err != nil {
			return ports.Info{}, customerror.MaxMindToCustom(err)
		}
		info.ASN = &asnInfo
	}
	return info, nil
}

func (I *IPInfo) ShortInfo(ip net.IP) (ports.ShortInfo, error) {
	info, err := I.Info(ip, true, true, true, true)
	if err != nil {
		return ports.ShortInfo{}, err
	}
	return ports.ShortInfo{
		IP:           ip.String(),
		Continent:    fmt.Sprintf("%s (%s)", info.Continent.Name, info.Continent.Code),
		Country:      fmt.Sprintf("%s (%s)", info.Country.Name, info.Country.ISOCode),
		Location:     fmt.Sprintf("%f, %f", info.City.Latitude, info.City.Longitude),
		TimeZone:     info.City.TimeZone,
		Organization: fmt.Sprintf("AS%d %s", info.ASN.Number, info.ASN.Organization),
	}, nil
}
