package models

import (
	"time"

	"go.monogdb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id"`
	Order_Data time.Time          `json:"order_date" validate="required"`
	Created_at time.Time          `json:"created_at"`
	updated_at time.Time          `json:"updated_at"`
	Order_id   string             `json:"order_id"`
	Table_id   *string            `json:"table_id" validate:"required"`
}
