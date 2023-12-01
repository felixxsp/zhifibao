package entity

import "github.com/google/uuid"

type Trc_req_one struct {
	Device      uuid.UUID `json:"device_uuid" bson:"device_uuid"`
	PersonID    uuid.UUID `json:"person_uuid" bson:"person_uuid"`
	Transaction uuid.UUID `json:"transaction_uuid" bson:"transaction_uuid"`
}

type Trc_req_multi struct {
	Device      uuid.UUID `json:"device_uuid" bson:"device_uuid"`
	PersonID    uuid.UUID `json:"person_uuid" bson:"person_uuid"`
	FilterStart int64     `json:"filter_start" bson:"filter_start"`
	FilterEnd   int64     `json:"filter_end" bson:"filter_end"`
}

type Login_req struct {
	Device   uuid.UUID `json:"device_uuid" bson:"device_uuid"`
	PersonID uuid.UUID `json:"person_uuid" bson:"person_uuid"`
	Password string    `json:"password" bson:"password"`
}
