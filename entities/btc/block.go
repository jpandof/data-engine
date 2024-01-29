package btc

type Block struct {
	Id string `bson:"_id" json:"i"`
}

func NewBlock(id string) *Block {
	return &Block{id}
}

func (b *Block) GetId() string {
	return b.Id
}
