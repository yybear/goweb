package api

import (
	. "github.com/yybear/goweb/db"
	. "github.com/yybear/goweb/entity"
	"github.com/labstack/echo"
	"log"
	. "github.com/yybear/goweb/helper"
	"net/http"
	"strconv"
)

func init() {
}

func Login(c *echo.Context) {
	r := c.Request

	session := GetSession(r)
	if session != nil {
		session.Invalidate()
	}
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")

	num, _ := Dbmap.SelectInt("select count(*) from users where name=? and password=?", name, password)
	if num == 1 {
		session = NewSession(r, c.Response.ResponseWriter)
		c.JSON(http.StatusOK, "{\"code\":0}")
	} else {
		c.JSON(http.StatusOK, "{\"code\":1}")
	}

}

func Logout(c *echo.Context) {
	r := c.Request
	session := GetSession(r)
	if session != nil {
		session.Invalidate()
	}
}

func CreateUser(c *echo.Context) {
	u := new(User)

	/*c.Request.ParseForm()
	name = c.Request.Form.Get("name") // 这种方式获取必须先要ParseForm 。。。
	log.Printf("name 44 %s", name)*/

	name := c.Request.FormValue("name")
	u.Name = name

	//if c.Bind(u) == nil { 绑定需要request的 headertype 为application/json
	Dbmap.Insert(u)
	c.JSON(http.StatusOK, u)
	//}
}

func GetUsers(c *echo.Context) {
	var users []User
	Dbmap.Select(&users, "select * from users")
	c.JSON(http.StatusOK, users)
}

func GetUser(c *echo.Context) {
	id, _ := strconv.ParseInt(c.P(0), 0, 64)

	user, _ := Dbmap.Get(User{}, id)

	c.JSON(http.StatusOK, user)

	defer log.Println("dddddddddd")
}
