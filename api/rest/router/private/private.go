package private

import (
	"github.com/HiBang15/signle-sign-on/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func SetRouter(router *gin.RouterGroup) {
	log.Print("Start init private router .....")
	// Define all private router in here
	//middleware auth
	//authMiddleware, err := auth.New(constant.SIGNING_ALGORITHM, constant.SECRET_KEY, constant.TOKEN_TIMEOUT)
	//if err != nil {
	//	log.Fatalf("has error when init auth middleware: %v", err)
	//}
	////group Auth
	//auth := router.Group("auth")
	//auth.Use(authMiddleware.MiddlewareFunc())
	//auth.GET("me", controller.AuthMe)
	//auth.POST("update")
	//auth.POST("refresh-token", controller.RefreshToken)
	//auth.POST("logout", controller.Logout)
	//auth.POST("logout-all", controller.LogoutAllDevice)


	//group User
	user := router.Group("user")
	{
		user.GET("/create", controllers.CreateUserAccount)
		//user.POST("/", controllers.CreateUserAccount)
		////user.GET("get-list-account")
		//user.PUT("/:id", controllers.UpdateUserAccount)
		//user.DELETE("/:id", controllers.DeleteUserAccount)
		//user.GET("/all", controllers.GetAllUser)
	}
	//
	////category
	//category := router.Group("category", authMiddleware.MiddlewareFunc())
	//{
	//	category.GET("/all", controller.GetListCategory)
	//}
	////product
	//product := router.Group("product", authMiddleware.MiddlewareFunc())
	//{
	//	product.GET("/detail/:uuid", controller.GetProduct)
	//	product.GET("/all", controller.GetListProduct)
	//	product.POST("/", controller.CreateProduct)
	//	//user.GET("get-list-account")
	//	user.PUT("/", controller.UpdateProduct)
	//	user.DELETE("/:uuid", controller.DeleteProduct)
	//	//user.DELETE("/:id", controllers.DeleteUserAccount)
	//	//user.GET("/all", controllers.GetAllUser)
	//}
	//
	////cart
	//cart := router.Group("cart", authMiddleware.MiddlewareFunc())
	//{
	//	cart.POST("/handler", controller.HandlerCart)
	//	cart.GET("/detail/user_id", controller.GetCart)
	//}
	//
	////common
	//common := router.Group("common", authMiddleware.MiddlewareFunc())
	//{
	//	common.POST("/upload", controller.Upload)
	//	common.POST("/upload-multi", controller.UploadMulti)
	//}

	log.Print("Finish init private router ....")
}