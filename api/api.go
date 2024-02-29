package api

import (
	"net/http"

	"github.com/joshjms/archive/filesystem"
	"github.com/labstack/echo/v4"
)

func Init(fs *filesystem.FileSystem) {

	e := echo.New()

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}
