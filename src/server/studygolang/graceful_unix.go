// +build !windows,!plan9

package main

import (
	"log"

	"github.com/facebookgo/grace/gracehttp"
	//"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo"
	//"echo/engine/standard"
)

func gracefulRun(e *echo.Echo) {
	log.Fatal(gracehttp.Serve(e.Server))
}
