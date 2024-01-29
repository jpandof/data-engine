package mars

type Wallet struct {
	ID           string              `json:"_id" bson:"_id"`
	Transactions []TransactionPriced `json:"Transactions" bson:"Transactions"`
	Analysis     *Analysis           `json:"Analysis,omitempty" bson:"Analysis,omitempty"`
}
