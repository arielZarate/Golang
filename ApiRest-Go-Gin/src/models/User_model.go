package models

import "gorm.io/gorm"



type User struct {
gorm.Model

ID uint                                            `json:"id" gorm:"primaryKey"`

Name string                                         `json:"name"`

LastName string                                     `json:"last_name"`

Email string       									`json:"email"`


}

//en MONGODB USARLO BSON
//Password		string`bson:"password" json:"password,omitempty"` 
type Login struct {
gorm.Model
Usersame string `json:"username"`
Password string `json:"password"`

}
   
