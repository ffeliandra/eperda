package model

import "time"

type User struct {
	ID       int32     `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama     string    `json:"nama" gorm:"type:varchar(255);not null"`
	Username string    `json:"username" gorm:"type:varchar(250);not null;unique"`
	Password string    `json:"-" gorm:"type:varchar(100);not null"`
	Created  time.Time `json:"created" gorm:"type:datetime;autoCreateTime:milli"`
	Modified time.Time `json:"modified" gorm:"type:datetime;autoUpdateTime:milli"`
}

type AllUsersWithTotal struct {
	Users []User `json:"users"`
	Total int64  `json:"total"`
}
