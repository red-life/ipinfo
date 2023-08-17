package ipinfo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/red-life/ipinfo/internal/adapters/http"
)

func NewApp(host string, port uint16, isDev bool, ipInfoHandler *http.IPInfoHandler, engine *gin.Engine) *App {
	return &App{
		host:          host,
		port:          port,
		isDev:         isDev,
		engine:        engine,
		ipInfoHandler: ipInfoHandler,
	}
}

type App struct {
	host          string
	port          uint16
	isDev         bool
	engine        *gin.Engine
	ipInfoHandler *http.IPInfoHandler
}

func (a *App) Run() error {
	http.RegisterRoutes(a.engine, a.ipInfoHandler)
	if a.isDev {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	return a.listenAndServe()
}

func (a *App) listenAndServe() error {
	address := fmt.Sprintf("%s:%d", a.host, a.port)
	return a.engine.Run(address)
}
