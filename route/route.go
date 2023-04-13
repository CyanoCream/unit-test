package route

import (
	"github.com/gin-gonic/gin"
	"sesi_8/handler"

	_ "sesi_8/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Employee API
// @version 1.0
// @description This is a sample service for managing employees
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.basic BasicAuth
// @host localhost:8080
// @BasePath /
func RegisterApi(r *gin.Engine, server handler.HttpServer) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // API for get swagger UI

	// middleware
	//r.Use(middleware.BasicAuth)

	api := r.Group("/book") // prefix
	{
		//api.POST("", server.CreateEmployee)
		//api.GET("/:id", server.GetEmployeeByID)
		//api.PUT("/:id", server.UpdateEmployee)
		//api.DELETE("/:id", server.DeleteEmployee)
		//api.GET("", server.GetBooks)
		api.POST("", server.CreateBook)
		api.GET("/:id", server.GetBookByID)
		//api.GET("/:id", server.GetBookByID())
		api.PUT("/:id", server.UpdateBook)
		api.DELETE("/:id", server.DeleteBook)
	}

}
