package api

import (
	"net/http"

	fs "github.com/joshjms/archive/filesystem"
	"github.com/labstack/echo/v4"
)

func Init(fs *fs.FileSystem) {
	e := echo.New()

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/files/*", func(c echo.Context) error {
		// Get the path from the URL
		path := c.Param("*")

		return c.String(http.StatusOK, path)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}
