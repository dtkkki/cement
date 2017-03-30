package web

import (
	"github.com/dtkkki/cement/web/handlers"
	"github.com/labstack/echo"
)

func route(app *echo.Echo) {
	hooks := app.Group("/webhook")
	hooks.POST("/github", handlers.GithubHookServer)
	hooks.POST("/gitlab", handlers.GitlabHookServer)
}
