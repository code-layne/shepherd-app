package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates map[string]*template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates[name].Execute(w, data)
}

func main() {
	e := echo.New()

	// Register templates
	templates := make(map[string]*template.Template)
	templates["home"] = template.Must(template.ParseFiles("html/layout.html", "html/index.html"))
	templates["algebra2"] = template.Must(template.ParseFiles("html/layout.html", "html/algebra2.html"))
	templates["algebra2::unit0"] = template.Must(template.ParseFiles("html/layout.html", "html/algebra2::unit0.html"))
	templates["algebra2::unit0::lesson01"] = template.Must(template.ParseFiles("html/layout.html", "html/algebra2::unit0::lesson01.html"))
	templates["algebra2::unit0::lesson02"] = template.Must(template.ParseFiles("html/layout.html", "html/algebra2::unit0::lesson02.html"))
	templates["algebra2::unit0::lesson03"] = template.Must(template.ParseFiles("html/layout.html", "html/algebra2::unit0::lesson03.html"))
	renderer := &TemplateRenderer{
		templates: templates,
	}
	e.Renderer = renderer

	e.Static("/assets", "static")
	e.Static("/i", "images")

	e.GET("/home", Home)
	e.GET("/algebra2", Algebra2)
	e.GET("/algebra2/unit0", Algebra2Unit0)
	e.GET("/algebra2/unit0/lesson01", Algebra2Unit0Lesson01)
	e.GET("/algebra2/unit0/lesson02", Algebra2Unit0Lesson02)
	e.GET("/algebra2/unit0/lesson03", Algebra2Unit0Lesson03)
	e.GET("/algebra2/unit0/lesson04", Algebra2Unit0Lesson04)
	e.GET("/algebra2/unit0/lesson05", Algebra2Unit0Lesson05)
	e.GET("/algebra2/unit0/lesson06", Algebra2Unit0Lesson06)

	e.Logger.Fatal(e.Start(":9000"))
}

func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home", nil)
}

func Algebra2(c echo.Context) error {
	return c.Render(http.StatusOK, "algebra2", nil)
}

func Algebra2Unit0(c echo.Context) error {
	return c.Render(http.StatusOK, "algebra2::unit0", nil)
}

func Algebra2Unit0Lesson01(c echo.Context) error {
	return c.Render(http.StatusOK, "algebra2::unit0::lesson01", nil)
}

func Algebra2Unit0Lesson02(c echo.Context) error {
	return c.Render(http.StatusOK, "algebra2::unit0::lesson02", nil)
}

func Algebra2Unit0Lesson03(c echo.Context) error {
	return c.Render(http.StatusOK, "algebra2::unit0::lesson03", nil)
}

func Algebra2Unit0Lesson04(c echo.Context) error {
	return c.Render(http.StatusOK, "algebra2::unit0::lesson04", nil)
}

func Algebra2Unit0Lesson05(c echo.Context) error {
	return c.Render(http.StatusOK, "algebra2::unit0::lesson05", nil)
}

func Algebra2Unit0Lesson06(c echo.Context) error {
	return c.Render(http.StatusOK, "algebra2::unit0::lesson06", nil)
}
