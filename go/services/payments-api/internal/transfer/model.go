package transfer

import "time"

type Transfer struct {
	FromAccountID string    `bson:"fromAccountId" json:"fromAccountId"`
	ToAccountID   string    `bson:"toAccountId" json:"toAccountId"`
	Amount        float64   `bson:"amount" json:"amount"`
	CreatedAt     time.Time `bson:"createdAt" json:"createdAt"`
}
type Account struct {
	ID      string  `bson:"_id"`
	Balance float64 `bson:"balance"`
}
