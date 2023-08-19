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
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	engine.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code to be 400 because remote IP is 127.0.0.1")
	ipInfoService.EXPECT().Info(net.ParseIP("185.143.233.200"), true, true, true, true).Return(ports.Info{}, nil)
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
	assert.Equal(t, false, exists, "Expected to have no errors")
}

func TestIPInfoHandler_ShortInfo(t *testing.T) {
	engine := gin.New()
	ipInfoService := mocks.IIPInfo{}
	ipInfoHandler := NewIPInfoHandler(&ipInfoService)
	RegisterRoutes(engine, ipInfoHandler)
	ipInfoService.EXPECT().ShortInfo(net.ParseIP("185.143.233.200")).Return(ports.ShortInfo{}, nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/185.143.233.200/short", nil)
	engine.ServeHTTP(w, req)
	response, _ := io.ReadAll(w.Body)
	var jsonResponse map[string]any
	t.Log(string(response))
	err := json.Unmarshal(response, &jsonResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %s", err.Error())
	}
	_, exists := jsonResponse["error"]
	assert.Equal(t, false, exists, "Expected to have no errors")
	_, exists = jsonResponse["org"]
	assert.Equal(t, true, exists, "Invalid response")
}
