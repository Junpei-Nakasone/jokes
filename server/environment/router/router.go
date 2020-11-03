package router

import (
	"net/http"

	api001 "nuxt-dadjokes/api/api001/handler"
	api002 "nuxt-dadjokes/api/api002/handler"
	api003 "nuxt-dadjokes/api/api003/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewRouter returns router
func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "PONG")
	})

	e.GET("sample", api001.SampleAPI)
	e.GET("fetchJokes", api002.FetchJokes)
	e.POST("addJoke", api003.AddJoke)

	e.Use(middleware.CORS())

	return e
}
