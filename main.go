package main

import (
	"encoding/json"
	"fmt"
	"github.com/oschwald/maxminddb-golang"
	"net"
)

func country() {
	reader, _ := maxminddb.Open("GeoLite2-Country.mmdb")
	var record struct {
		Country struct {
			ISOCode string `maxminddb:"iso_code"`
			Names   struct {
				En string `maxminddb:"en"`
			} `maxminddb:"names"`
		} `maxminddb:"registered_country"`
	}
	var recordAny any
	reader.Lookup(net.ParseIP("185.143.233.200"), &record)
	reader.Lookup(net.ParseIP("185.143.233.200"), &recordAny)
	fmt.Printf("%+v\n", record)
	js, _ := json.Marshal(&recordAny)
	fmt.Printf("%+v\n", string(js))
}

func city() {
	fmt.Println("city")
	reader, _ := maxminddb.Open("GeoLite2-City.mmdb")
	var _ struct {
		Country struct {
			ISOCode string `maxminddb:"iso_code"`
			Names   struct {
				En string `maxminddb:"en"`
			} `maxminddb:"names"`
		} `maxminddb:"registered_country"`
	}
	//reader.Lookup(net.ParseIP("185.143.233.200"), &record)
	//fmt.Printf("%+v\n", record)
	var recordAny any
	reader.Lookup(net.ParseIP("8.8.8.8"), &recordAny)
	js, _ := json.Marshal(&recordAny)
	fmt.Printf("%+v\n", string(js))
}

func main() {
	city()
}
