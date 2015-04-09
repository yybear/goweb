package action

import (
	. "db"
	. "entity"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)

var users []User

func init() {
}

func CreateUser(c *echo.Context) {
	u := new(User)

	/*c.Request.ParseForm()
	name = c.Request.Form.Get("name") // 这种方式获取必须先要ParseForm 。。。
	log.Printf("name 44 %s", name)*/

	name := c.Request.FormValue("name")
	u.Name = name
	log.Printf("name 44 %s", name)

	//if c.Bind(u) == nil { 绑定需要request的 headertype 为application/json
	Dbmap.Insert(u)
	c.JSON(http.StatusOK, u)
	//}
}

func GetUsers(c *echo.Context) {
	Dbmap.Select(&users, "select * from users")
	c.JSON(http.StatusOK, users)
}

func GetUser(c *echo.Context) {
	id, _ := strconv.ParseInt(c.P(0), 0, 64)

	user, _ := Dbmap.Get(User{}, id)

	c.JSON(http.StatusOK, user)
}
