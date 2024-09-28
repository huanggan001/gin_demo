package router

import (
	"gin_demo/common/lib"
	"gin_demo/controller"
	"gin_demo/docs"
	"gin_demo/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {

	docs.SwaggerInfo.Title = lib.ConfBase.Swagger.Title
	docs.SwaggerInfo.Description = lib.ConfBase.Swagger.Desc
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = lib.ConfBase.Swagger.Host
	docs.SwaggerInfo.BasePath = lib.ConfBase.Swagger.BasePath

	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//demo
	testRouter := router.Group("test")
	testRouter.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept"},
	}), middleware.TranslationMiddleware())
	{
		controller.TestRegister(testRouter)
	}
	//demo
	//adminLoginRouter := router.Group("/admin_login")
	//store, err := sessions.NewRedisStore(10, "tcp", lib.GetStringConf("base.session.redis_server"), lib.GetStringConf("base.session.redis_password"), []byte("secret"))
	//if err != nil {
	//	log.Fatalf("sessions.NewRedisStore err : %v", err)
	//}
	//adminLoginRouter.Use(
	//	sessions.Sessions("mysession", store),
	//	middleware.RecoveryMiddleware(),
	//	middleware.RequestLog(),
	//	middleware.TranslationMiddleware())
	//{
	//	controller.AdminLoginRegister(adminLoginRouter)
	//}

	return router
}
