package mw

import (
	"github.com/labstack/echo"
	"log"
)

func SessionAuth(h echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c *echo.Context) {
		log.Println("session auth mw")
		r := c.Request

		if "/logout" == r.RequestURI || "/login" == r.RequestURI {

		} else {
			session := GetSession(r)
			if session == nil {
				log.Println("need login")
				c.Response.WriteHeader(401)
				return
			}
		}
		h(c)
	})
}
