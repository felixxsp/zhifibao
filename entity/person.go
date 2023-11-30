package entity

import "github.com/google/uuid"

type Person struct {
	UUID         uuid.UUID     `json:"uuid" bson:"uuid"`
	Username     string        `json:"username" bson:"username"`
	Password     string        `json:"passworkd" bson:"passworkd"`
	Name         string        `json:"full_name" bson:"full_name"`
	Phone        Phone         `json:"phone" bson:"phone"`
	Balance      float32       `json:"balance" bson:"balance"`
	Transactions []Transaction `json:"transactions" bson:"transactions"`
	Active       bool          `json:"active_status" bson:"active_status"`
	ActiveDevice uuid.UUID     `json:"active_device" bson:"active_device"`
}

type Phone struct {
	CountryCode string `json:"code" bson:"code"`
	Number      string `json:"number" bson:"number"`
}
