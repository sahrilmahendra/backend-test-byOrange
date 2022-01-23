package models

import "time"

type PurchaseOrderHeader struct {
	ID             uint      `gorm:"primaryKey"`
	PO_Date        string    `json:"po_date" form:"po_date"`
	PO_Description string    `json:"po_description" form:"po_description"`
	PO_Price       int       `json:"po_price" form:"po_price"`
	PO_Cost        int       `json:"po_cost" form:"po_cost"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdateAt       time.Time `gorm:"autoUpdateTime"`
}

type PurchaseOrderDetail struct {
	ID         uint `gorm:"primaryKey"`
	ItemsID    uint `json:"item_id"`
	ItemQty    int  `json:"item_qty"`
	ItemPrice  int
	ItemCost   int
	POHeaderID uint
}

type PurchaseOrderRequest struct {
	PO_Date             string
	PO_Description      string `json:"po_description"`
	PO_Price            int
	PO_Cost             int
	PurchaseOrderDetail []PurchaseOrderDetail `json:"po_detail"`
}

type PurchaseOrderDetailResponse struct {
	IDDetail  uint
	IDItem    uint
	ItemQty   int
	ItemPrice int
	ItemCost  int
}

// type PurchaseOrderResponse struct {
// 	IDHeader                    uint
// 	PO_Date                     string
// 	PO_Description              string
// 	PO_Price                    int
// 	PO_Cost                     int
// 	PurchaseOrderDetailResponse []PurchaseOrderDetailResponse `json:"PO_Detail"`
// }
