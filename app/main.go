package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

var PORT = os.Getenv("PORTO")

// Quick response object
type GenericResponse struct {
	Message string `json:"message"`;
}

// Example handler
func exampleRoute(c echo.Context) error {
	// Make a quick response object
	resp := &GenericResponse {
		Message: "todo bien",
	}

	return c.JSON(http.StatusOK, resp)
}

func main()  {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Grabs the port to run on from the env var
	if len(PORT) == 0 {
		fmt.Println("A port value is not set! Defaulting to 3000")
		PORT = ":3000"
	} else {
		PORT = ":" + PORT
	}

	// Routes
	e.GET("/", exampleRoute)

	e.Logger.Fatal(e.Start(PORT))
}
