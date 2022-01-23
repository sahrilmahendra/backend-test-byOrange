package controllers

import (
	"byorange/helper"
	"byorange/lib/databases"
	"byorange/middlewares"
	"byorange/models"
	response "byorange/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// controller untuk menambahkan item baru
func CreateItemControllers(c echo.Context) error {
	new_item := models.Items{}
	c.Bind(&new_item)
	err := helper.ValidateItem(new_item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Request Data"))
	}
	login := middlewares.ExtractTokenId(c)
	new_item.UsersID = uint(login)
	_, err = databases.CreateItem(&new_item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}

// controller untuk menampilkan seluruh data item
func GetAllItemControllers(c echo.Context) error {
	items, err := databases.GetAllItems()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	if items == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", items))
}

// controller untuk menampilkan data item by id
func GetItemByIdControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	item, err := databases.GetItemById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	if item == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", item))
}

// controller untuk memperbarui data item by id
func UpdateItemControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	update_item := models.Items{}
	c.Bind(&update_item)
	err = helper.ValidateItem(update_item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Request Data"))
	}
	login := middlewares.ExtractTokenId(c)
	id_user, _ := databases.GetIDUserByIdItem(id)
	if id_user != uint(login) {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	_, err = databases.UpdateItem(id, &update_item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}

// controller untuk menghapus item by id
func DeleteItemControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	item, _ := databases.GetItemById(id)
	if item == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	login := middlewares.ExtractTokenId(c)
	id_user, _ := databases.GetIDUserByIdItem(id)
	if id_user != uint(login) {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	databases.DeleteItem(id)
	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}
