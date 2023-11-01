package main

import (
	"example/go-api-with-db-course/database"
	_ "example/go-api-with-db-course/docs"
	"example/go-api-with-db-course/internal/api"

	"github.com/gin-gonic/gin"
)

// @title			GO API Course
// @version		1.0
// @description	GO API using Gin framework.
// @host			localhost:8000
// @BasePath		/
func main() {
	database.ConnectDB()
	r := gin.Default()

	api.SetupRoutes(r)

	r.Run(":8000") // listen and serve on 0.0.0.0:8000 (for windows "localhost:8000")
}
