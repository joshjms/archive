package controllers

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gosimple/slug"
	fs "github.com/joshjms/archive/filesystem"
	"github.com/labstack/echo/v4"
)

type FileRequest struct {
	Name     string `json:"name"`
	ParentID int    `json:"parent_id"`
}

type FileResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DirectoryRequest struct {
	Name     string `json:"name"`
	ParentID int    `json:"parent_id"`
}

func GetFile(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func CreateFile(c echo.Context, f *fs.FileSystem) error {
	name := c.FormValue("name")
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

	newFile := &fs.File{
		Name:     name,
		ParentID: parentID,
		Slug:     slug.Make(name),
	}

	err = f.CreateFile(newFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create file",
		})
	}

	return c.String(http.StatusOK, "Hello, World!")
}

func GetDirectory(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func CreateDirectory(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
