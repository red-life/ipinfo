package http

import (
	"github.com/gin-gonic/gin"
	"github.com/red-life/ipinfo/internal/pkg/customerror"
	"github.com/red-life/ipinfo/internal/ports"
	"github.com/red-life/ipinfo/internal/utils"
	"net"
	"net/http"
)

func RegisterRoutes(engine *gin.Engine, ipInfoHandler *IPInfoHandler) {
	engine.GET("/", ValidateIP, ipInfoHandler.Info)
	engine.GET("/:ip", ValidateIP, ipInfoHandler.Info)
	engine.GET("/short", ValidateIP, ipInfoHandler.ShortInfo)
	engine.GET("/:ip/short", ValidateIP, ipInfoHandler.ShortInfo)
}

func NewIPInfoHandler(ipInfoService ports.IIPInfo) *IPInfoHandler {
	return &IPInfoHandler{
		ipInfoService: ipInfoService,
	}
}

type IPInfoHandler struct {
	ipInfoService ports.IIPInfo
}

func (I *IPInfoHandler) Info(c *gin.Context) {
	response := make(map[string]any)
	ip, exists := c.Get("ip")
	if !exists {
		response["error"] = "ip not found"
		c.JSON(http.StatusNotFound, response)
		return
	}
	continent := utils.StringToBool(c.DefaultQuery("continent", "true"))
	country := utils.StringToBool(c.DefaultQuery("country", "true"))
	city := utils.StringToBool(c.DefaultQuery("city", "true"))
	asn := utils.StringToBool(c.DefaultQuery("asn", "true"))
	info, err := I.ipInfoService.Info(ip.(net.IP), continent, country, city, asn)
	if err != nil {
		response["error"] = err.Error()
		c.JSON(customerror.ErrorToStatusCode(err), response)
		return
	}
	c.JSON(http.StatusOK, &info)
	return
}

func (I *IPInfoHandler) ShortInfo(c *gin.Context) {
	var response map[string]any
	ip, exists := c.Get("ip")
	if !exists {
		response["error"] = "ip not found"
		c.JSON(http.StatusNotFound, response)
		return
	}
	info, err := I.ipInfoService.ShortInfo(ip.(net.IP))
	if err != nil {
		response["error"] = err.Error()
		c.JSON(customerror.ErrorToStatusCode(err), response)
		return
	}
	c.JSON(http.StatusOK, info)
	return
}
