package controllers

import (
	"byorange/helper"
	"byorange/lib/databases"
	"byorange/models"
	response "byorange/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

// controller untuk menambahkan data purchase order
func CreatePurchaseOrderControllers(c echo.Context) error {
	po_req := models.PurchaseOrderRequest{}
	c.Bind(&po_req)

	err := helper.ValidatePO(po_req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Data Request"))
	}

	price, cost, err := databases.CreatePurchaseOrder(&po_req)
	if price == 0 || cost == 0 || err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}

// controller untuk menampilkan seluruh data purchase order
func GetAllPurchaseOrderControllers(c echo.Context) error {
	purchase_order, err := databases.GetAllOrderDetail()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	if purchase_order == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", purchase_order))
}
