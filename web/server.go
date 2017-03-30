package web

import (
	"github.com/dtkkki/cement/web/router"
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
	router.MountAPI(app)
	// Init middlewares
}
