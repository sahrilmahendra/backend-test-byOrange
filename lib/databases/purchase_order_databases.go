package databases

import (
	"byorange/config"
	"byorange/models"
	"fmt"
	"time"
)

// function database untuk menambahkan data purchase order
func CreatePurchaseOrder(purchase_order *models.PurchaseOrderRequest) (int, int, error) {
	// header
	var (
		price_tot, cost_tot int
	)
	for i := 0; i < len(purchase_order.PurchaseOrderDetail); i++ {
		id := purchase_order.PurchaseOrderDetail[i].ItemsID
		price_item, cost_item, err := GetPriceCostByIdItem(int(id))
		if err != nil || price_item == 0 || cost_item == 0 {
			return price_item, cost_item, err
		}
		price_tot_per_item := price_item * purchase_order.PurchaseOrderDetail[i].ItemQty
		cost_tot_per_item := cost_item * purchase_order.PurchaseOrderDetail[i].ItemQty
		price_tot += price_tot_per_item
		cost_tot += cost_tot_per_item
	}
	po_header := models.PurchaseOrderHeader{}
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	po_header.PO_Date = formatted
	po_header.PO_Description = purchase_order.PO_Description
	po_header.PO_Price = price_tot
	po_header.PO_Cost = cost_tot
	if err := config.DB.Create(&po_header).Error; err != nil {
		return price_tot, cost_tot, err
	}

	// detail
	po_detail := models.PurchaseOrderDetail{}
	for i := 0; i < len(purchase_order.PurchaseOrderDetail); i++ {
		po_detail = purchase_order.PurchaseOrderDetail[i]
		po_detail.POHeaderID = po_header.ID
		id := purchase_order.PurchaseOrderDetail[i].ItemsID
		price_item, cost_item, err := GetPriceCostByIdItem(int(id))
		if err != nil {
			return price_item, cost_item, err
		}
		price_tot_per_item := price_item * purchase_order.PurchaseOrderDetail[i].ItemQty
		cost_tot_per_item := cost_item * purchase_order.PurchaseOrderDetail[i].ItemQty
		po_detail.ItemPrice = price_tot_per_item
		po_detail.ItemCost = cost_tot_per_item
		if err := config.DB.Create(&po_detail).Error; err != nil {
			config.DB.Where("id = ?", po_header.ID).Delete(&po_header)
			return price_item, cost_item, err
		}
	}
	return po_detail.ItemPrice, po_detail.ItemCost, nil
}

// function database untuk menampilkan seluruh data purchase order header
func GetAllOrderHeader() (interface{}, error) {
	order_header := []models.PurchaseOrderHeader{}
	err := config.DB.Table("purchase_order_headers").Select("*").Find(&order_header)
	if err != nil || order_header == nil {
		return nil, err.Error
	}
	return order_header, nil
}

// function database untuk menampilkan seluruh data purchase order header
func GetAllOrderDetail() (interface{}, error) {
	order_detail := []models.PurchaseOrderDetailResponse{}
	err := config.DB.Table("purchase_order_details").Select("*").Find(&order_detail)
	if err != nil || order_detail == nil {
		return nil, err.Error
	}
	return order_detail, nil
}
