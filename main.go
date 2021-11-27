package main

import (
	"github.com/miftahulhidayati/rest-api-go/controllers"
	"github.com/miftahulhidayati/rest-api-go/database"

	"github.com/gin-gonic/gin"
)

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
	
	router.Run(":8080")
}