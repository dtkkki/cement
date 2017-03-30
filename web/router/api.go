package router

import (
	"github.com/dtkkki/cement/web/router/hook"
	"github.com/dtkkki/cement/web/router/log"
	_ "github.com/dtkkki/cement/web/router/user"
	"github.com/labstack/echo"
)

func MountAPI(e *echo.Echo) {
	group := e.Group("/apiv1")
	hook.MountHook(group)
	log.MountLog(group)
}