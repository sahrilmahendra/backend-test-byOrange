package models

import "gorm.io/gorm"

// struktur data Item
type Items struct {
	gorm.Model
	Item_Name        string `json:"item_name" form:"item_name"`
	Item_Description string `json:"item_description" form:"item_description"`
	Price            int    `json:"price" form:"price"`
	Cost             int    `json:"cost" form:"cost"`
	UsersID          uint
}

// struktur get data item
type GetItem struct {
	ID               uint
	Item_Name        string
	Item_Description string
	Price            int
	Cost             int
	UsersID          uint
}
