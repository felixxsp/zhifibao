package entity

import "github.com/google/uuid"

type Person struct {
	UUID         uuid.UUID `json:"uuid" bson:"uuid"`
	Username     string    `json:"username" bson:"username"`
	Password     string    `json:"password" bson:"password"`
	Name         string    `json:"full_name" bson:"full_name"`
	Phone        Phone     `json:"phone" bson:"phone"`
	Balance      float32   `json:"balance" bson:"balance"`
	Active       bool      `json:"active_status" bson:"active_status"`
	Friends      []Person
	ActiveDevice uuid.UUID `json:"active_device" bson:"active_device"`
}

type Phone struct {
	CountryCode string `json:"code" bson:"code"`
	Number      string `json:"number" bson:"number"`
}
