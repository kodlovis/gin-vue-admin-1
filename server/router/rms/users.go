package rms

import (
	"gin-vue-admin/api/v1/rms"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitUsersRouter(Router *gin.RouterGroup) {
	UsersRouter := Router.Group("Users").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		UsersRouter.POST("createUsers", rms.CreateUsers)   // 新建Users
		UsersRouter.DELETE("deleteUsers", rms.DeleteUsers) // 删除Users
		UsersRouter.DELETE("deleteUsersByIds", rms.DeleteUsersByIds) // 批量删除Users
		UsersRouter.PUT("updateUsers", rms.UpdateUsers)    // 更新Users
		UsersRouter.GET("findUsers", rms.FindUsers)        // 根据ID获取Users
		UsersRouter.GET("getUsersList", rms.GetUsersList)  // 获取Users列表
	}
}
