package models

import "gorm.io/gorm"

type User struct {
	ID       int    `gorm:"column:id;type:int unsigned PRIMARY KEY AUTO_INCREMENT"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"column:password;type:varchar(255)"`
	Email    string `json:"email" gorm:"column:email;type:varchar(255)"`
	Status   int    `json:"status" gorm:"column:status;type:tinyint(1)"`
	Balance  int    `json:"balance" gorm:"column:balance;type:int(5)"`
	gorm.Model
}

func (u User) TableName() string { return "users" }
