package services

import (
	"apirest/src/models"
	// "apirest/src/database"
)




func GetAllUsers() {//([]models.User, error) 
	//var users []models.User
//	result := database.DB.Find(&users)
//	return users, result.Error
}

func GetUserByID(id string){ //(models.User, error) {
//	var user models.User
	//result := database.DB.First(&user, id)
	//return user, result.Error
}

func CreateUser(user models.User){ //(models.User, error) {
	//result := database.DB.Create(&user)
	//return user, result.Error


}

func UpdateUser(id string, updatedUser models.User){ //(models.User, error) {
	//var user models.User
//	result := database.DB.First(&user, id)
	/*
	if result.Error != nil {
		return user, result.Error
	}
	result = database.DB.Model(&user).Updates(updatedUser)
	return user, result.Error
	*/
}

func DeleteUser(id string){ // error {
	//result := database.DB.Delete(&models.User{}, id)
	//return result.Error
}

