package mars

type TransactionPriced struct {
	Id            *string `json:"_id" bson:"_id"`
	TypeOperation *string `json:"t" bson:"t"`
	Date          *int64  `json:"d" bson:"d"`
	Quantity      *int64  `json:"q" bson:"q"`
	UnitPrice     *int64  `json:"u" bson:"u"`
	IDWallet      *string `json:"w,omitempty" bson:"w,omitempty"`
}

func (tx TransactionPriced) GetUnitPrice() float64 {
	return float64(*tx.UnitPrice)
}
func (tx TransactionPriced) GetTotalPrice() float64 {
	return tx.GetUnitPrice() * tx.GetQuantity()
}
func (tx TransactionPriced) GetQuantity() float64 {
	return float64(*tx.Quantity) / 10e8
}
