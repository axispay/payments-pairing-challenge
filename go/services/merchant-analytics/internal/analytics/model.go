package analytics

import "time"

type Transaction struct {
	ID         interface{} `bson:"_id,omitempty" json:"_id,omitempty"`
	MerchantID string      `bson:"merchantId" json:"merchantId"`
	Amount     float64     `bson:"amount" json:"amount"`
	Status     string      `bson:"status" json:"status"`
	CreatedAt  time.Time   `bson:"createdAt" json:"createdAt"`
}

type DailyTotals struct {
	Total    float64        `json:"total"`
	Count    int            `json:"count"`
	ByStatus map[string]int `json:"byStatus"`
}
