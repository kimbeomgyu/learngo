package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/myProject/learngo/checker"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", handleHome)
	e.GET("/profile/:user", handleProfile)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}

func handleHome(c echo.Context) error {
	return c.File("index.html")
}

func handleProfile(c echo.Context) error {
	url := "https://api.github.com/users/" + c.Param("user")
	res, err := http.Get(url)
	checker.CheckErr(err)
	checker.CheckStatusCode(res)

	defer res.Body.Close()

	body, bodyErr := ioutil.ReadAll(res.Body)
	checker.CheckErr(bodyErr)
	fmt.Println(string(body))
	return c.String(res.StatusCode, string(body))
}
