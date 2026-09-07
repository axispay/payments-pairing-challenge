package account

type Account struct {
	ID         string  `bson:"_id" json:"_id"`
	OwnerName  string  `bson:"ownerName" json:"ownerName"`
	Balance    float64 `bson:"balance" json:"balance"`
	MerchantID string  `bson:"merchantId" json:"merchantId"`
}
