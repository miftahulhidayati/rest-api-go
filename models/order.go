package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	// OrderID      int       `json:"order_id"`
	CustomerName string    
	OrderedAt    time.Time 
	Items        []Item    
}
