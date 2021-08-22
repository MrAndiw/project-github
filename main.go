package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World Home!")
}

func main() {
	fmt.Println("Hello World.")

	e := echo.New()

	e.GET("/", home)

	e.Start(":8080")
}
