package http

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

func ValidateIP(c *gin.Context) {
	response := make(map[string]any)
	remoteIP := c.Param("ip")
	if remoteIP == "" {
		remoteIP = c.RemoteIP()
	}
	ip := net.ParseIP(remoteIP)
	if ip == nil {
		response["error"] = "invalid ip address"
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}
	c.Set("ip", ip)
	c.Next()
}
