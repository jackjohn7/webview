package main

import (
	//"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//AllowOrigins: []string{"http://localhost:3000", "http://localhost:5173"},
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	t := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	app.Renderer = t

	app.GET("/", func (c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", time.Now())
	})

	app.GET("/upload", func(c echo.Context) error {
		return c.Render(http.StatusOK, "upload_video_form.html", "\"localhost:3000/upload_video\"")
	})
	//w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")

	app.Start(":5173")
}
