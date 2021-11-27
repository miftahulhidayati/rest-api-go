package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miftahulhidayati/rest-api-go/models"
	"gorm.io/gorm"
)


func (conn *DBConn) CreateOrder(c *gin.Context) {
	var order models.Order
	var result gin.H

	err := c.ShouldBindJSON(&order)
	if err != nil{
		result = gin.H{
			"result": "insert failed",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	tx := conn.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
		tx.Rollback()
		}
	}()

	err = conn.DB.Create(&order).Error
	if err != nil{
		tx.Rollback()
    	result = gin.H{
			"result": "insert failed",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	result = gin.H{
		"result": order,
	}
	c.JSON(http.StatusOK, result)
}

func (conn *DBConn) GetOrder(c *gin.Context) {
	var (
		order models.Order
		result gin.H
	)
	id := c.Param("id")
	err := conn.DB.Where("id = ?", id).Preload("Items").First(&order).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": order,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (conn *DBConn) GetOrders(c *gin.Context) {
	var (
		orders []models.Order
		result  gin.H
	)

	err := conn.DB.Preload("Items").Find(&orders).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
		c.JSON(http.StatusBadRequest, result)
		return
	} 		
	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (conn *DBConn) UpdateOrder(c *gin.Context) {
	
	id := c.Query("id")
	var order models.Order
	var result gin.H
	err := conn.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	err = c.ShouldBindJSON(&order)
	if err != nil{
		result = gin.H{
			"result": "update failed",
		}
		
	}
	tx := conn.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
		tx.Rollback()
		}
	}()

	// err = conn.DB.Create(&order).Error
	err = conn.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order).Error
	if err != nil{
		tx.Rollback()
    	result = gin.H{
			"result": "update failed",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	result = gin.H{
		"result": order,
	}
	c.JSON(http.StatusOK, result)
}

func (conn *DBConn) DeleteOrder(c *gin.Context) {
	var (
		order models.Order
		result gin.H
	)
	id := c.Param("id")
	err := conn.DB.Where("id = ?", id).Preload("Items").First(&order).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	tx := conn.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
		tx.Rollback()
		}
	}()

	err = conn.DB.Delete(&order, id).Error
	if err != nil{
	
		tx.Rollback()
		result = gin.H{
			"result": "delete failed",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	err = conn.DB.Where("order_id = ?", id).Delete(&order.Items).Error
	if err != nil {
		tx.Rollback()
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
