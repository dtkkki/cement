package log

import (
	log "github.com/dtkkki/cement/web/buisness/log"
	"github.com/labstack/echo"
)

func MountLog(group *echo.Group) {
	logger := group.Group("/log")
	logger.GET("/:id", log.Pineline)
}
