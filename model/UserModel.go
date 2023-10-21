package model

import "time"

type User struct {
	ID       uint
	Name     string    `gorm:"type:varchar(600)"`
	Email    string    `gorm:"type:varchar(250);unique;not null"`
	Password string    `gorm:"type:varchar(600)"`
	Active   int       `gorm:"type:tinyint(10);default:1"`
	Created  time.Time `gorm:"autoCreateTime"`
	Updated  time.Time `gorm:"autoUpdateTime"`
}
