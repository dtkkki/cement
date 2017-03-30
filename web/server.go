package web

import (
	"github.com/labstack/echo"
)

// app is a app.
var app = echo.New()

// App returns the app.
func App() *echo.Echo {
	return app
}

func init() {
	// Init routes
	route(app)
	// Init middlewares
}
