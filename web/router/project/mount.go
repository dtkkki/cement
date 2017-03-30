package project

import (
	"github.com/dtkkki/cement/web/buisness/project"
	"github.com/dtkkki/cement/web/middleware"
	"github.com/labstack/echo"
)

//MountProject ...
func MountProject(group *echo.Group) {
	projects := group.Group("projects", middleware.UserLoginRequired)
	projects.GET("/list", project.HttpCementProjectsListAPI)
}
