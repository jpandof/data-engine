package btc

type Tx struct {
	Id      string       `bson:"_id" json:"i"`
	Time    *int64       `bson:"t" json:"t"`
	Block   *Block       `bson:"k" json:"k"`
	Sellers []*TxPartner `bson:"s" json:"s"`
	Buyers  []*TxPartner `bson:"b" json:"b"`
	Price   *int64       `bson:"p,omitempty" json:"p,omitempty"`
}

func NewTx(id string, block *Block) *Tx {
	return &Tx{Id: id, Block: block}
}
func NewTxWithTime(id string, time *int64, block *Block) *Tx {
	return &Tx{Id: id, Time: time, Block: block}
}

func NewEmptyTx() *Tx {
	return &Tx{}
}

func (t *Tx) GetId() string {
	return t.Id
}

func (t *Tx) GetSellers() []*TxPartner {
	return t.Sellers
}

func (t *Tx) GetBuyers() []*TxPartner {
	return t.Buyers
}

func (t *Tx) SetSellers(sellers []*TxPartner) {
	t.Sellers = sellers
}

func (t *Tx) SetBuyers(buyers []*TxPartner) {
	t.Buyers = buyers
}

func (t *Tx) GetSellerByPosition(position int) *TxPartner {
	for _, seller := range t.Buyers {
		if seller.Position == position {
			return seller
		}
	}
	return nil
}

func (t *Tx) AdjustTx() {
	merged := make(map[string]int64)

	// Merge all TxPartners
	if len(t.Sellers) > 0 {
		for _, seller := range t.Sellers {
			merged[seller.Id] += seller.Quantity
		}
	} else {
		//println("Not found sellers in adjust")
	}

	if len(t.Buyers) > 0 {
		for _, buyer := range t.Buyers {
			merged[buyer.Id] += buyer.Quantity
		}
	} else {
		//println("Not found buyers in adjust")
	}

	// Clear Sellers and Buyers slices
	t.Sellers = nil
	t.Buyers = nil

	// Create new Sellers and Buyers slices based on merged results
	if len(merged) == 0 {
		println("Not found merged in adjust")
		return
	}

	for id, quantity := range merged {
		if quantity < 0 {
			t.Sellers = append(t.Sellers, &TxPartner{Id: id, Quantity: -1 * quantity})
		} else if quantity > 0 {
			t.Buyers = append(t.Buyers, &TxPartner{Id: id, Quantity: quantity})
		}
	}

}
