package web

import (
	"github.com/dtkkki/cement/web/router/hook"
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
	mountAPI(app)
	// Init middlewares
}

func mountAPI(e *echo.Echo) {
	group := e.Group("/apiv1")
	hook.MountHook(group)
}
