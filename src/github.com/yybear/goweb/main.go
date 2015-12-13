package main

import (
	. "github.com/yybear/goweb/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rs/cors"
	"github.com/thoas/stats"
	"log"
	. "github.com/yybear/goweb/helper"
	"reflect"
	"github.com/golang/glog"
)

func main() {
	log.Println("starting ...")

	e := echo.New()
	e.Use(middleware.Logger)
	e.Use(cors.Default().Handler)

	// https://github.com/thoas/stats
	s := stats.New()
	e.Use(s.Handler)
	e.Use(SessionAuth)
	// Route
	e.Get("/stats", func(c *echo.Context) {
		c.JSON(200, s.Data())
	})

	// Serve index file
	e.Index("public/index.html")

	// Serve static files
	e.Static("/js", "public/js")
	e.Static("/css", "public/css")

	e.Post("/login", Login)
	e.Get("/logout", Logout)

	e.Post("/users", CreateUser)
	e.Get("/users", GetUsers)
	e.Get("/users/:id", GetUser)

	// Start server
	e.Run(":8080")

	//t := testFunc(tfunc)
	//t.test(1, "dd")
}

/*type ti interface {
	test(int, string)
}

type testFunc func(i int)

func (f testFunc) test(i int, s string) {
	f(i)
}

func tfunc(i int) {
	log.Printf("%d , %s", i, "fuck")
}*/
