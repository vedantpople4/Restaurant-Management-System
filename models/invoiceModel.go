package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID                primitive.ObjectID `bson:"_id`
	Invoice_id        string             `json:"invoice_id"`
	Order_id          string             `json:"order_id"`
	Payement_method   string             `json:"payement_method" validate:"eq=CARD|eq=CASH|eq="`
	Payement_status   *string            `json:"payement_method" validate:"required, eq=PENDING|eq=PAID"`
	Payement_due_date time.Time          `json:"Payement_due_date`
	Created_at        time.Time          `json:"created_at"`
	Updated_at        time.Time          `json:"updated_at"`
}
