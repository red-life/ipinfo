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

func getResponse(handler http.Handler, method string, url string, body io.Reader) (*httptest.ResponseRecorder, []byte, error) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, body)
	handler.ServeHTTP(w, req)
	response, _ := io.ReadAll(w.Body)
	return w, response, nil
}

func toJson(data []byte) (map[string]any, error) {
	var jsonResponse map[string]any
	err := json.Unmarshal(data, &jsonResponse)
	return jsonResponse, err
}

func TestIPInfoHandler_Info(t *testing.T) {
	engine := gin.New()
	ipInfoService := mocks.IIPInfo{}
	ipInfoHandler := NewIPInfoHandler(&ipInfoService)
	RegisterRoutes(engine, ipInfoHandler)
	ipInfoService.EXPECT().Info(net.ParseIP("185.143.233.200"), true, true, false, false).Return(ports.Info{
		Continent: &ports.Continent{},
		Country:   &ports.Country{},
	}, nil)
	_, response, _ := getResponse(engine, "GET", "/185.143.233.200?continent=true&country=1&city=false&asn=0", nil)
	jsonResponse, err := toJson(response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %s", err.Error())
	}
	_, exists := jsonResponse["error"]
	assert.False(t, exists, "Expected to have no errors")
	_, exists = jsonResponse["continent"]
	assert.True(t, exists, "Expected to have continent in response")
	_, exists = jsonResponse["country"]
	assert.True(t, exists, "Expected to have country in response")
	_, exists = jsonResponse["asn"]
	assert.False(t, exists, "Expected not to have asn in response")
	_, exists = jsonResponse["city"]
	assert.False(t, exists, "Expected not to have city in response")
}

func TestIPInfoHandler_ShortInfo(t *testing.T) {
	engine := gin.New()
	ipInfoService := mocks.IIPInfo{}
	ipInfoHandler := NewIPInfoHandler(&ipInfoService)
	RegisterRoutes(engine, ipInfoHandler)
	ipInfoService.EXPECT().ShortInfo(net.ParseIP("185.143.233.200")).Return(ports.ShortInfo{}, nil)
	_, response, _ := getResponse(engine, "GET", "/185.143.233.200/short", nil)
	jsonResponse, err := toJson(response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %s", err.Error())
	}
	_, exists := jsonResponse["error"]
	assert.False(t, exists, "Expected to have no errors")
	_, exists = jsonResponse["org"]
	assert.True(t, exists, "Invalid response")
}
