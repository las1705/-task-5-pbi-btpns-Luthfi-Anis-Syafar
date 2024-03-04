package routes

import (
	"UserPhoto-API/config/app_config"
	"UserPhoto-API/controllers/auth_controller"
	"UserPhoto-API/controllers/photo_controller"
	"UserPhoto-API/controllers/user_controller"
	"UserPhoto-API/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app.Group("")

	route.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)

	// AUTH ROUTE LOGIN

	// ROUTE USER
	userRoute := route.Group("users")
	userRoute.POST("/login", auth_controller.Login)
	userRoute.POST("/register", user_controller.Register)
	userRoute.PUT("/update/:userId", middleware.AuthMiddleware, user_controller.UpdateUser)
	userRoute.DELETE("/delete/:userId", middleware.AuthMiddleware, user_controller.DeleteUser)

	// ROUTE FILE
	photoRoute := route.Group("photos")
	photoRoute.POST("/", middleware.AuthMiddleware, photo_controller.HandlerUploadPhoto)
	photoRoute.PUT("/:photoId", middleware.AuthMiddleware, photo_controller.HandlerEditPhoto)
	photoRoute.DELETE("/:photoId", middleware.AuthMiddleware, photo_controller.HandlerRemovePhoto)
	photoRoute.GET("/", middleware.AuthMiddleware, photo_controller.GetPhoto)

}
