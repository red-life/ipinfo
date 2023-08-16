package http

import (
	"github.com/gin-gonic/gin"
	"github.com/red-life/ipinfo/internal/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIPInfoHandler_Info(t *testing.T) {
	engine := gin.New()
	ipInfoService := mocks.IIPInfo{}
	ipInfoHandler := NewIPInfoHandler(&ipInfoService)
	RegisterRoutes(engine, ipInfoHandler)
	assert.Equal(t, "/*ip", engine.Routes()[0].Path, "Expected registered handlers length to be 1")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	engine.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code to be OK because remote IP is 127.0.0.1")

}
