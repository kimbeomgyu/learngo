package main

import (
	"os"
	"strings"

	"learngo/webScrapper/scrapper"

	"github.com/labstack/echo"
)

const filename string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}
func handleScrape(c echo.Context) error {
	defer os.Remove(filename)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(filename, filename)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":3000"))
}
