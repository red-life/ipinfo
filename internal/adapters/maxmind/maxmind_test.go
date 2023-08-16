package maxmind

import (
	"github.com/oschwald/maxminddb-golang"
	"github.com/red-life/ipinfo/internal/pkg/customerror"
	"github.com/red-life/ipinfo/internal/ports"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestMaxMind_GetASN(t *testing.T) {
	reader, err := maxminddb.Open("tmp/GeoLite2-ASN.mmdb")
	if err != nil {
		t.Fatal(err)
	}
	mm := NewMaxMind(reader, reader, reader)
	testCases := []struct {
		arg1     net.IP
		expected struct {
			asn ports.ASN
			err error
		}
	}{
		{
			arg1: net.ParseIP("1.1.1.1"),
			expected: struct {
				asn ports.ASN
				err error
			}{asn: ports.ASN{Number: 13335, Organization: "CLOUDFLARENET"}, err: nil},
		},
		{
			arg1: net.ParseIP("185.143.234.200"),
			expected: struct {
				asn ports.ASN
				err error
			}{asn: ports.ASN{Number: 205585, Organization: "Noyan Abr Arvan Co. ( Private Joint Stock)"}, err: nil},
		},
		{
			arg1: net.ParseIP("46.143.100.70"),
			expected: struct {
				asn ports.ASN
				err error
			}{asn: ports.ASN{Number: 43754, Organization: "Asiatech Data Transmission company"}, err: nil},
		},
		{
			arg1: net.ParseIP("8.8.8.8"),
			expected: struct {
				asn ports.ASN
				err error
			}{asn: ports.ASN{Number: 15169, Organization: "GOOGLE"}, err: nil},
		},
	}
	for _, testCase := range testCases {
		result, err := mm.GetASN(testCase.arg1)
		assert.Equal(t, testCase.expected.err, err, "%s: Error expected %s but got %s", testCase.arg1, testCase.expected.err, err)
		assert.Equal(t, testCase.expected.asn, result, "%s: ASN expected %v but got %v", testCase.arg1, testCase.expected.asn, result)
	}
}

func TestMaxMind_GetCity(t *testing.T) {
	reader, err := maxminddb.Open("tmp/GeoLite2-City.mmdb")
	if err != nil {
		t.Fatal(err)
	}
	mm := NewMaxMind(reader, reader, reader)
	testCases := []struct {
		arg1     net.IP
		expected struct {
			city ports.City
			err  error
		}
	}{
		{
			arg1: net.ParseIP("1.1.1.1"),
			expected: struct {
				city ports.City
				err  error
			}{city: ports.City{}, err: nil},
		},
		{
			arg1: net.ParseIP("185.143.234.200"),
			expected: struct {
				city ports.City
				err  error
			}{city: ports.City{Latitude: 35.698, Longitude: 51.4115, TimeZone: "Asia/Tehran"}, err: customerror.NotFoundErr},
		},
		{
			arg1: net.ParseIP("46.143.100.70"),
			expected: struct {
				city ports.City
				err  error
			}{city: ports.City{Longitude: 52.6668, Latitude: 36.544, TimeZone: "Asia/Tehran"}, err: nil},
		},
		{
			arg1: net.ParseIP("8.8.8.8"),
			expected: struct {
				city ports.City
				err  error
			}{city: ports.City{Longitude: -97.822, Latitude: 37.751, TimeZone: "America/Chicago"}, err: nil},
		},
	}
	for _, testCase := range testCases {
		result, err := mm.GetCity(testCase.arg1)
		assert.Equal(t, testCase.expected.err, err, "%s: Error expected %s but got %s", testCase.arg1, testCase.expected.err, err)
		assert.Equal(t, testCase.expected.city, result, "%s: City expected %v but got %v", testCase.arg1, testCase.expected.city, result)
	}
}

func TestMaxMind_GetContinent(t *testing.T) {
	reader, err := maxminddb.Open("tmp/GeoLite2-City.mmdb")
	if err != nil {
		t.Fatal(err)
	}
	mm := NewMaxMind(reader, reader, reader)
	testCases := []struct {
		arg1     net.IP
		expected struct {
			continent ports.Continent
			err       error
		}
	}{
		{
			arg1: net.ParseIP("1.1.1.1"),
			expected: struct {
				continent ports.Continent
				err       error
			}{continent: ports.Continent{}, err: customerror.NotFoundErr},
		},
		{
			arg1: net.ParseIP("185.143.234.200"),
			expected: struct {
				continent ports.Continent
				err       error
			}{continent: ports.Continent{Code: "AS", Name: "Asia"}, err: nil},
		},
		{
			arg1: net.ParseIP("46.143.100.70"),
			expected: struct {
				continent ports.Continent
				err       error
			}{continent: ports.Continent{Code: "AS", Name: "Asia"}, err: nil},
		},
		{
			arg1: net.ParseIP("8.8.8.8"),
			expected: struct {
				continent ports.Continent
				err       error
			}{continent: ports.Continent{Code: "NA", Name: "North America"}, err: nil},
		},
	}
	for _, testCase := range testCases {
		result, err := mm.GetContinent(testCase.arg1)
		assert.Equal(t, testCase.expected.err, err, "%s: Error expected %s but got %s", testCase.arg1, testCase.expected.err, err)
		assert.Equal(t, testCase.expected.continent, result, "%s: Continent expected %v but got %v", testCase.arg1, testCase.expected.continent, result)
	}
}

func TestMaxMind_GetCountry(t *testing.T) {
	reader, err := maxminddb.Open("tmp/GeoLite2-Country.mmdb")
	if err != nil {
		t.Fatal(err)
	}
	mm := NewMaxMind(reader, reader, reader)
	testCases := []struct {
		arg1     net.IP
		expected struct {
			country ports.Country
			err     error
		}
	}{
		{
			arg1: net.ParseIP("1.1.1.1"),
			expected: struct {
				country ports.Country
				err     error
			}{country: ports.Country{ISOCode: "AU", Name: "Australia"}, err: nil},
		},
		{
			arg1: net.ParseIP("185.143.234.200"),
			expected: struct {
				country ports.Country
				err     error
			}{country: ports.Country{ISOCode: "IR", Name: "Iran"}, err: nil},
		},
		{
			arg1: net.ParseIP("46.143.100.70"),
			expected: struct {
				country ports.Country
				err     error
			}{country: ports.Country{ISOCode: "IR", Name: "Iran"}, err: nil},
		},
		{
			arg1: net.ParseIP("8.8.8.8"),
			expected: struct {
				country ports.Country
				err     error
			}{country: ports.Country{ISOCode: "US", Name: "United States"}, err: nil},
		},
	}
	for _, testCase := range testCases {
		result, err := mm.GetCountry(testCase.arg1)
		assert.Equal(t, testCase.expected.err, err, "%s: Error expected %s but got %s", testCase.arg1, testCase.expected.err, err)
		assert.Equal(t, testCase.expected.country, result, "%s: Country expected %v but got %v", testCase.arg1, testCase.expected.country, result)
	}
}
