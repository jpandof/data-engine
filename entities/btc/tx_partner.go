package btc

type TxPartner struct {
	Id       string `bson:"i" json:"i"`
	Position int    `bson:"p" json:"p"`
	Quantity int64  `bson:"q" json:"q"`
}

func NewTxPartner(id string, position int, quantity int64) *TxPartner {
	return &TxPartner{id, position, quantity}
}

func (tP *TxPartner) GetQuantity() int64 {
	return tP.Quantity
}
func (tP *TxPartner) SetQuantity(quantity int64) {
	tP.Quantity = quantity
}
