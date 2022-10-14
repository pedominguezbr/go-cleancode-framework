package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "framework-auto-aprov-go/cmd/api/docs"
	handler "framework-auto-aprov-go/pkg/http/handler"
	middleware "framework-auto-aprov-go/pkg/http/middleware"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(
	userHandler *handler.UserHandler,
	rolHandler *handler.RolHandler,
	auto_Aprov_OmnitokHandler *handler.Auto_Aprov_OmnitokHandler) *ServerHTTP {

	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())
	engine.Use(CORS())

	apiV1 := engine.Group("/api/v1")
	// Swagger docs
	apiV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Request JWT
	apiuth := apiV1.Group("/auth")
	apiuth.POST("/login", userHandler.LoginHandler)
	apiuth.POST("/signUp", userHandler.SignUp)
	apiuth.POST("/refresh-token", userHandler.RefreshToken)

	// Auth middleware users
	apiUsers := apiV1.Group("/users", middleware.AuthorizationMiddleware)
	apiUsers.GET("", userHandler.FindAll)
	apiUsers.GET("/:id", userHandler.FindByID)
	apiUsers.GET("/:id/roles", userHandler.FindByIDWithRole)
	//api.POST("users", userHandler.Save)
	apiUsers.DELETE("/:id", userHandler.Delete)
	apiUsers.DELETE("/:id/roles", userHandler.DeleteRoles)

	// Auth middleware servicio
	apiAuto_Aprov := apiV1.Group("/omnitok", middleware.AuthorizationMiddleware)
	apiAuto_Aprov.GET("", middleware.PaginationMiddleware, auto_Aprov_OmnitokHandler.FindAll)
	apiAuto_Aprov.GET("/:id", auto_Aprov_OmnitokHandler.FindByID)
	apiAuto_Aprov.POST("save", auto_Aprov_OmnitokHandler.Save)
	apiAuto_Aprov.PUT("update/:id", auto_Aprov_OmnitokHandler.Update)
	apiAuto_Aprov.DELETE("/:id", auto_Aprov_OmnitokHandler.Delete)

	// Auth middleware Familia producto
	apiRol := apiV1.Group("/rol", middleware.AuthorizationMiddleware)
	apiRol.GET("", rolHandler.FindAll)
	apiRol.GET("/:id", rolHandler.FindByID)
	apiRol.POST("save", rolHandler.Save)
	apiRol.PUT("update/:id", rolHandler.Update)
	apiRol.DELETE("/:id", rolHandler.Delete)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {

		//fmt.Println(c.Request.Header)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, Origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		//c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT, DELETE")

		// if c.Request.Method == "OPTIONS" {
		// 	c.AbortWithStatus(204)
		// 	return
		// }

		c.Next()
	}
}
