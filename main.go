package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main(){
	app := echo.New()
	app.Use(middleware.Recover())
	
	app.GET("/", func(c echo.Context) error {
		t := time.Now()
		f := t.Format("2006-01-02 15:04:05")
		return c.HTML(http.StatusOK, f)
	})

	if err := app.Start(":9998"); err != nil {
		panic(err)
	}
}