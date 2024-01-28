package entities

import "time"

type Blkfile struct {
	Id            string     `bson:"_id"`
	ProcessedTime *time.Time `bson:"processed_time"`
}

type Block struct {
	Id string `bson:"_id" json:"i"`
}

type Tx struct {
	Id      string       `bson:"_id" json:"i"`
	Time    *int64       `bson:"t" json:"t"`
	Block   *Block       `bson:"k" json:"k"`
	Sellers []*TxPartner `bson:"s" json:"s"`
	Buyers  []*TxPartner `bson:"b" json:"b"`
}

type TxPartner struct {
	Id       string `bson:"i" json:"i"`
	Position int    `bson:"p" json:"p"`
	Quantity int64  `bson:"q" json:"q"`
}
