package databases

import (
	"byorange/config"
	"byorange/models"
)

// function database untuk membuat item baru
func CreateItem(item *models.Items) (interface{}, error) {
	if err := config.DB.Create(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

// function database untuk menampilkan seluruh data item
func GetAllItems() (interface{}, error) {
	var get_items []models.GetItem
	query := config.DB.Table("items").Select("*").Where("items.deleted_at IS NULL").Find(&get_items)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_items, nil
}

// function database untuk menampilkan data item by id
func GetItemById(id int) (interface{}, error) {
	var get_item models.GetItem
	query := config.DB.Table("items").Select("*").Where("items.deleted_at IS NULL AND items.id = ?", id).Find(&get_item)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_item, nil
}

// function database untuk menampilkan id user by id item
func GetIDUserByIdItem(id int) (uint, error) {
	var get_item models.GetItem
	query := config.DB.Table("items").Select("*").Where("items.deleted_at IS NULL AND items.id = ?", id).Find(&get_item)
	if query.Error != nil || query.RowsAffected == 0 {
		return 0, query.Error
	}
	return get_item.UsersID, nil
}

// function database untuk menampilkan price & cost by id item
func GetPriceCostByIdItem(id int) (int, int, error) {
	var get_item models.GetItem
	query := config.DB.Table("items").Select("*").Where("items.deleted_at IS NULL AND items.id = ?", id).Find(&get_item)
	if query.Error != nil || query.RowsAffected == 0 {
		return 0, 0, query.Error
	}
	return get_item.Price, get_item.Cost, nil
}

// function database untuk menampilkan item by user id
func GetItemByUserId(id int) (interface{}, error) {
	var get_item models.GetItem
	query := config.DB.Table("items").Select("*").Where("items.deleted_at IS NULL AND items.users_id = ?", id).Find(&get_item)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_item, nil
}

// function database untuk memperbarui data item by id
func UpdateItem(id int, update_item *models.Items) (interface{}, error) {
	var item models.Items
	query_select := config.DB.Find(&item, id)
	if query_select.Error != nil || query_select.RowsAffected == 0 {
		return 0, query_select.Error
	}
	query_update := config.DB.Model(&item).Updates(update_item)
	if query_update.Error != nil {
		return nil, query_update.Error
	}
	return item, nil
}

// function database untuk menghapus item by id
func DeleteItem(id int) (interface{}, error) {
	var item models.Items
	check_item := config.DB.Find(&item, id).RowsAffected

	err := config.DB.Delete(&item).Error
	if err != nil || check_item > 0 {
		return nil, err
	}
	return item, nil
}
