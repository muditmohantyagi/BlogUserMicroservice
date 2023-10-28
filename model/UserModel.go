package model

import (
	"errors"
	"time"

	"blog.com/config"
	"gorm.io/gorm"
)

type User struct {
	ID       uint
	Name     string `gorm:"type:varchar(600)"`
	Email    string `gorm:"type:varchar(250);unique;not null"`
	Password string `gorm:"type:varchar(600)"`
	Active   int    `gorm:"type:tinyint(10);default:1"`
	JwtToken string
	Created  time.Time `gorm:"autoCreateTime"`
	Updated  time.Time `gorm:"autoUpdateTime"`
}

var db = config.GoConnect().Debug()

func UserByID(UserId uint) (*User, error) {
	var user *User = new(User)
	if result := db.Where("id=? AND active=?", UserId, 1).Take(&user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			return nil, nil
		}
		return nil, result.Error
	}
	return user, nil
}

func UpdateToken(UserId uint, JwtToken string) (bool, error) {
	user_data, err := UserByID(UserId)
	if err != nil {
		return false, err
	}
	if user_data != nil {
		user_data.JwtToken = JwtToken
		if result := db.Model(&User{}).Where("id=?", UserId).Updates(&user_data); result.Error != nil {
			return false, result.Error
		} else if result.RowsAffected == 0 {
			err := errors.New("record not updated")

			return false, err
		}
		return true, nil
	} else {
		return false, errors.New("User not found")
	}
}
