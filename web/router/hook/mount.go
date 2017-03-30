package hook

import (
	hook "github.com/dtkkki/cement/web/buisness/hook"
	"github.com/labstack/echo"
)

func MountHook(group *echo.Group) {
	hooks := group.Group("/webhook")
	hooks.POST("/github", hook.GithubHookServer)
	hooks.POST("/gitlab", hook.GitlabHookServer)
}
