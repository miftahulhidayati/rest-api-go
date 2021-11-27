package main

import (
	"github.com/miftahulhidayati/rest-api-go/controllers"
	"github.com/miftahulhidayati/rest-api-go/database"

	"github.com/gin-gonic/gin"
	docs "github.com/miftahulhidayati/rest-api-go/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Order
// @version 1.0
// @description This is a sample API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	router := gin.Default()

	db := database.InitMysqlDB()

	DBConn := &controllers.DBConn{DB: db}
	//Read All
	router.GET("/orders", DBConn.GetOrders)
	//Read One
	router.GET("/orders/:id", DBConn.GetOrder)
	//Create
	router.POST("/orders", DBConn.CreateOrder)
	//Update
	router.PUT("/orders/:id", DBConn.UpdateOrder)
	//Delete
	router.DELETE("/orders/:id", DBConn.DeleteOrder)
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	router.Run(":8080")
}