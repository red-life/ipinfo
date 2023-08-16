package http

import (
	"github.com/gin-gonic/gin"
	"github.com/red-life/ipinfo/internal/pkg/customerror"
	"github.com/red-life/ipinfo/internal/ports"
	"net"
	"net/http"
)

func RegisterRoutes(engine *gin.Engine, ipInfoHandler *IPInfoHandler) {
	engine.GET("/*ip", ipInfoHandler.Info)
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
	remoteIP := c.Param("ip")
	if remoteIP == "" {
		remoteIP = c.RemoteIP()
	}
	ip := net.ParseIP(remoteIP)
	if ip == nil {
		response["error"] = "invalid ip address"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	info, err := I.ipInfoService.GetInfo(ip)
	if err != nil {
		response["error"] = err.Error()
		c.JSON(customerror.ErrorToStatusCode(err), response)
		return
	}
	c.JSON(http.StatusOK, info)
	return
}
