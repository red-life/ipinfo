package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/red-life/ipinfo/internal/mocks"
	"github.com/red-life/ipinfo/internal/ports"
	"github.com/stretchr/testify/assert"
	"io"
	"net"
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

	ipInfoService.EXPECT().GetInfo(net.ParseIP("185.143.233.200")).Return(ports.Info{}, nil)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/185.143.233.200", nil)
	engine.ServeHTTP(w, req)
	response, _ := io.ReadAll(w.Body)
	var jsonResponse map[string]any
	err := json.Unmarshal(response, &jsonResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %s", err.Error())
	}
	_, exists := jsonResponse["error"]
	assert.Equal(t, false, exists, "Expected to has no errors")
}
