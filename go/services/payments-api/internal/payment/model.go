package payment

import "time"

type Payment struct {
	ID        interface{} `bson:"_id,omitempty" json:"_id,omitempty"`
	AccountID string      `bson:"accountId" json:"accountId"`
	Amount    float64     `bson:"amount" json:"amount"`
	Currency  string      `bson:"currency" json:"currency"`
	Status    string      `bson:"status" json:"status"`
	CreatedAt time.Time   `bson:"createdAt" json:"createdAt"`
}
