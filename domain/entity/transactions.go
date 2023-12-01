package entity

import (
	"github.com/google/uuid"
)

type Transaction struct {
	Person   uuid.UUID `json:"person_uuid" bson:"person_uuid"`
	Receiver uuid.UUID `json:"receiver_uuid" bson:"receiver_uuid"`
	UUID     uuid.UUID `json:"uuid" bson:"uuid"`
	Title    string    `json:"title" bson:"title"`
	Amount   float32   `json:"amount" bson:"amount"`
	Balance  float32   `json:"balance" bson:"balance"`
	Type     bool      `json:"transaction_type" bson:"transaction_type"`
	Time     int64     `json:"time" bson:"time"`
	Notes    string    `json:"notes" bson:"notes"`
	Device   uuid.UUID `json:"device_uuid" bson:"device_uuid"`
}
