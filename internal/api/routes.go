package api

import (
	"example/go-api-with-db-course/internal/api/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", handlers.Home)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/facts", handlers.ListFacts)
	r.GET("/facts/:id", handlers.GetFactByID)
	r.POST("/facts", handlers.CreateFacts)
}
