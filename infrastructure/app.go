package infrastructure

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Application interface {
	Register(method AppMethod, path string, handler echo.HandlerFunc)
	Run(addr string) error
}

type App struct {
	server *echo.Echo
}

type AppMethod string

const (
	AppMethodGet    AppMethod = "GET"
	AppMethodPost   AppMethod = "POST"
	AppMethodPut    AppMethod = "PUT"
	AppMethodDelete AppMethod = "DELETE"
)

func NewApp() *App {
	e := echo.New()

	app := &App{server: e}

	app.initMiddleware()
	return app
}

func (a *App) Register(method AppMethod, path string, handler echo.HandlerFunc) {

	switch method {
	case AppMethodGet:
		a.server.GET(path, handler)
	case AppMethodPost:
		a.server.POST(path, handler)
	case AppMethodPut:
		a.server.PUT(path, handler)
	case AppMethodDelete:
		a.server.DELETE(path, handler)
	default:
		panic("unsupported method")
	}
}

func (a *App) Run(addr string) error {
	return a.server.Start(addr)
}

func HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
