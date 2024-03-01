package fs

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
)

func (fs *FileSystem) GetFileRoute(c echo.Context) error {
	path := c.Param("*")

	return c.String(http.StatusOK, path)
}

func (fs *FileSystem) CreateFileRoute(c echo.Context) error {
	parentID, err := strconv.Atoi(c.FormValue("parent_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid parent_id",
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid file",
		})
	}

	name := file.Filename

	path := os.Getenv("UPLOAD_PATH")

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to open file",
		})
	}

	defer src.Close()

	dst, err := os.Create(path + file.Filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create file",
		})
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to copy file",
		})
	}

	newFile := &File{
		Name:     name,
		ParentID: parentID,
		Slug:     slug.Make(name),
	}

	err = fs.CreateFile(newFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create file",
		})
	}

	return c.String(http.StatusOK, "Hello, World!")
}
