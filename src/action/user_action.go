package action

import (
	. "db"
	. "entity"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

var users map[int64]User

func init() {
}

func CreateUser(c *echo.Context) {
	u := new(User)

	if c.Bind(u) == nil {
		Dbmap.Insert(u)
		c.JSON(http.StatusOK, u)
	}
}

func GetUsers(c *echo.Context) {
	c.JSON(http.StatusOK, users)
}

func GetUser(c *echo.Context) {
	id, _ := strconv.ParseInt(c.P(0), 0, 64)

	user, _ := Dbmap.Get(User{}, id)

	c.JSON(http.StatusOK, user)
}
