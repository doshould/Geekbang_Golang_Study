package router

import (
	"github.com/LyricTian/gin-admin/v6/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterAPI register api group router
func (a *Router) RegisterAPI(app *gin.Engine) {
	// 权限系统内部接口
	g := app.Group("/api")

	// 请求频率限制
	g.Use(middleware.RateLimiterMiddleware())

	// 跳过token检查
	g.Use(middleware.UserAuthMiddleware(a.Auth,
		middleware.AllowPathPrefixSkipper("/api/pub/login"),
	))

	// 跳过鉴权
	g.Use(middleware.CasbinMiddleware(a.CasbinEnforcer,
		middleware.AllowPathPrefixSkipper("/api/pub"),
		middleware.AllowPathPrefixSkipper("/api/external"),
	))

	pub := g.Group("/pub")
	{
		gLogin := pub.Group("login")
		{
			//gLogin.GET("captchaid", a.LoginAPI.GetCaptcha)
			//gLogin.GET("captcha", a.LoginAPI.ResCaptcha)
			gLogin.POST("", a.LoginAPI.Login)
		}

		//pub.POST("/refreshToken", a.LoginAPI.RefreshToken)
		//pub.POST("/logout", a.LoginAPI.Logout)

		//gCurrent := pub.Group("current")
		//{
		//	gCurrent.PUT("password", a.LoginAPI.UpdatePassword)
		//	gCurrent.GET("user", a.LoginAPI.GetUserInfo)
		//	gCurrent.GET("menutree", a.LoginAPI.QueryUserMenuTree)
		//}

		gUser := pub.Group("users")
		{
			gUser.POST("", a.AccountAPI.Create)
			//gUser.GET("", a.UserAPI.Query)
			//gUser.GET(":id", a.UserAPI.Get)
			//gUser.PUT(":id", a.UserAPI.Update)
			//gUser.DELETE(":id", a.UserAPI.Delete)
			//gUser.PATCH(":id/enable", a.UserAPI.Enable)
			//gUser.PATCH(":id/disable", a.UserAPI.Disable)
		}
	}

	external := g.Group("/external")
	{
		gLogin := external.Group("login")
		{
			gLogin.GET("authenticate", a.LoginAPI.Authenticate)
		}
	}

	//gDemo := pub.Group("demos")
	//{
	//	gDemo.GET("", a.DemoAPI.Query)
	//	gDemo.GET(":id", a.DemoAPI.Get)
	//	gDemo.POST("", a.DemoAPI.Create)
	//	gDemo.PUT(":id", a.DemoAPI.Update)
	//	gDemo.DELETE(":id", a.DemoAPI.Delete)
	//	gDemo.PATCH(":id/enable", a.DemoAPI.Enable)
	//	gDemo.PATCH(":id/disable", a.DemoAPI.Disable)
	//}

	//gMenu := pub.Group("menus")
	//{
	//	gMenu.GET("", a.MenuAPI.Query)
	//	gMenu.GET(":id", a.MenuAPI.Get)
	//	gMenu.POST("", a.MenuAPI.Create)
	//	gMenu.PUT(":id", a.MenuAPI.Update)
	//	gMenu.DELETE(":id", a.MenuAPI.Delete)
	//	gMenu.PATCH(":id/enable", a.MenuAPI.Enable)
	//	gMenu.PATCH(":id/disable", a.MenuAPI.Disable)
	//}
	//pub.GET("/menus.tree", a.MenuAPI.QueryTree)

	//gRole := pub.Group("roles")
	//{
	//	gRole.GET("", a.RoleAPI.Query)
	//	gRole.GET(":id", a.RoleAPI.Get)
	//	gRole.POST("", a.RoleAPI.Create)
	//	gRole.PUT(":id", a.RoleAPI.Update)
	//	gRole.DELETE(":id", a.RoleAPI.Delete)
	//	gRole.PATCH(":id/enable", a.RoleAPI.Enable)
	//	gRole.PATCH(":id/disable", a.RoleAPI.Disable)
	//}
	//pub.GET("/roles.select", a.RoleAPI.QuerySelect)
}
