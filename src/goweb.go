package main

import (
	. "action"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/rs/cors"
	"github.com/thoas/stats"
)

func main() {
	e := echo.New()
	e.Use(mw.Logger)
	e.Use(cors.Default().Handler)

	// https://github.com/thoas/stats
	s := stats.New()
	e.Use(s.Handler)
	// Route
	e.Get("/stats", func(c *echo.Context) {
		c.JSON(200, s.Data())
	})

	// Serve index file
	e.Index("public/index.html")

	// Serve static files
	e.Static("/js", "public/js")
	e.Static("/css", "public/css")

	e.Post("/users", CreateUser)
	e.Get("/users", GetUsers)
	e.Get("/users/:id", GetUser)

	// Start server
	e.Run(":8080")
}
