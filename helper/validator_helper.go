package helper

import (
	"byorange/models"

	validator "github.com/go-playground/validator/v10"
)

type ValidatorItem struct {
	Item_Name        string `validate:"required"`
	Item_Description string `validate:"required"`
	Price            int    `validate:"required,gt=0"`
	Cost             int    `validate:"required,gt=0"`
}

// function untuk validasi item
func ValidateItem(item models.Items) error {
	v := validator.New()
	validate_item := ValidatorItem{
		Item_Name:        item.Item_Name,
		Item_Description: item.Item_Description,
		Price:            item.Price,
		Cost:             item.Cost,
	}
	err := v.Struct(validate_item)
	if err != nil {
		return err
	}
	return nil
}

type ValidatorPODetail struct {
	Item_ID  uint `validate:"required,gt=0"`
	Item_Qty int  `validate:"required,gt=0"`
}

// function untuk validasi purchase order
func ValidatePO(purchase_order models.PurchaseOrderRequest) error {
	v := validator.New()
	err := v.Var(purchase_order.PO_Description, "required")
	if err != nil {
		return err
	}
	for i := 0; i < len(purchase_order.PurchaseOrderDetail); i++ {
		validate_po_detail := ValidatorPODetail{
			Item_ID:  purchase_order.PurchaseOrderDetail[i].ItemsID,
			Item_Qty: purchase_order.PurchaseOrderDetail[i].ItemQty,
		}

		err := v.Struct(validate_po_detail)
		if err != nil {
			return err
		}
	}
	return nil
}
