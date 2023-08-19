package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/oschwald/maxminddb-golang"
	"github.com/red-life/ipinfo/cmd/ipinfo"
	"github.com/red-life/ipinfo/internal/adapters/http"
	"github.com/red-life/ipinfo/internal/adapters/maxmind"
	"github.com/red-life/ipinfo/internal/services"
	"github.com/red-life/ipinfo/internal/utils"
	"os"
	"path"
	"strconv"
)

const (
	countryFileName = "GeoLite2-Country.mmdb"
	cityFileName    = "GeoLite2-City.mmdb"
	asnFileName     = "GeoLite2-ASN.mmdb"
)

func getMMDBReaders(mmdbFilesPath string) (country, city, asn *maxminddb.Reader, err error) {
	countryMMDBPath := path.Join(mmdbFilesPath, countryFileName)
	cityMMDBPath := path.Join(mmdbFilesPath, cityFileName)
	asnMMDBPath := path.Join(mmdbFilesPath, asnFileName)
	country, err = maxminddb.Open(countryMMDBPath)
	city, err = maxminddb.Open(cityMMDBPath)
	asn, err = maxminddb.Open(asnMMDBPath)
	return
}

func getEnv() (host string, port uint16, mmdbFiles string, isDev bool, err error) {
	host = os.Getenv("host")
	intPort, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		return "", 0, "", false, err
	}
	port = uint16(intPort)
	mmdbFiles = os.Getenv("mmdb_files")
	isDev = utils.StringToBool(os.Getenv("is_dev"))
	return
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	err := godotenv.Load(".env")
	checkError(err)
	host, port, mmdbFiles, isDev, err := getEnv()
	countryReader, cityReader, asnReader, err := getMMDBReaders(mmdbFiles)
	checkError(err)
	maxMind := maxmind.NewMaxMind(countryReader, cityReader, asnReader)
	ipInfoService := services.NewIPInfo(maxMind)
	ipInfoHandler := http.NewIPInfoHandler(ipInfoService)
	engine := gin.Default()
	ipInfoApp := ipinfo.NewApp(host, port, isDev, ipInfoHandler, engine)
	err = ipInfoApp.Run()
	checkError(err)
}
