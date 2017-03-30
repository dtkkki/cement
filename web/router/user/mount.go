package user

import (
	"github.com/dtkkki/cement/web/buisness/user"
	"github.com/labstack/echo"
)

//MountUser just mount user api
func MountUser(group *echo.Group) {
	users := group.Group("/user")
	_ = users

	admin := group.Group("/admin")
	admin.GET("/user/list", user.HTTPCementUserListAPI)
	admin.GET("/events/list", user.HTTPCementAdminEventsAPI)

	events := group.Group("/event")
	events.POST("/user/project", user.HTTPCementUserProjectAPI)

	_ = admin
}
