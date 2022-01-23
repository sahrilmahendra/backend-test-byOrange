package databases

import (
	"byorange/config"
	"byorange/middlewares"
	"byorange/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var user models.Users

// function database untuk menampilkan seluruh data users
func GetAllUsers() (interface{}, error) {
	var users []models.GetUser
	query := config.DB.Table("users").Select("*").Where("users.deleted_at IS NULL").Find(&users)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return users, nil
}

// function database untuk menampilkan user by id
func GetUserById(id int) (interface{}, error) {
	var get_user_by_id models.GetUser
	query := config.DB.Table("users").Select("*").Where("users.deleted_at IS NULL AND users.id = ?", id).Find(&get_user_by_id)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return get_user_by_id, nil
}

// function database untuk menambahkan user baru (registrasi)
func CreateUser(user *models.Users) (interface{}, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// function database untuk menghapus user by id
func DeleteUser(id int) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// function database untuk memperbarui data user by id
func UpdateUser(id int, user *models.Users) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Updates(&user).Error; err != nil {
		return nil, err
	}
	config.DB.First(&user, id)
	return user, nil
}

// function login database untuk mendapatkan token
func LoginUser(plan_pass string, user *models.Users) (interface{}, error) {
	var get_login models.GetLoginUser
	err := config.DB.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return nil, err
	}

	// cek plan password dengan hash password
	match := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plan_pass))
	if match != nil {
		return nil, match
	}
	user.Token, err = middlewares.CreateToken(int(user.ID)) // generate token
	if err != nil {
		return nil, err
	}
	// restruktur respons
	get_login.ID = user.ID
	get_login.Token = user.Token
	user.LastLogin = time.Now()
	if err = config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return get_login, nil
}
