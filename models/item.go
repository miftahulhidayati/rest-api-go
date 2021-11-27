package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	// ItemID      int    `json:"item_id"`
	ItemCode    string 
	Description string 
	Quantity    int    
	OrderID     int    
}