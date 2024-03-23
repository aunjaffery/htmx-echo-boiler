package main

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/aunjaffery/htmx-echo-boiler/views/home"
	"github.com/aunjaffery/htmx-echo-boiler/views/misc"
	"github.com/labstack/echo/v4"
)

func RenderX(c echo.Context, comp templ.Component) error {
	if c.Request().Header.Get("Hx-Request") == "true" {
		err := comp.Render(c.Request().Context(), c.Response())
		if err != nil {
			return c.String(500, "Failed to render webpage.")

		}
	} else {
		err := misc.Base(comp, true).Render(c.Request().Context(), c.Response())
		if err != nil {
			return c.String(500, "Failed to render webpage.")
		}
	}
	return nil
}
func main() {
	fmt.Println("daisy ui")

	e := echo.New()

	e.Static("/static", "./public")

	e.GET("/", func(c echo.Context) error {
		comp := home.Home()
		return RenderX(c, comp)
	})

	e.Logger.Fatal(e.Start(":6969"))
}
