package price

type Price struct {
	ID        int    `bson:"_id" json:"_id"`
	Date      string `bson:"date" json:"date"`
	Symbol    string `bson:"symbol" json:"symbol"`
	Open      *int64 `bson:"open" json:"open"`
	High      *int64 `bson:"high" json:"high"`
	Low       *int64 `bson:"low" json:"low"`
	Close     *int64 `bson:"close" json:"close"`
	VolumeBTC *int64 `bson:"Volume BTC" json:"Volume BTC"`
	VolumeUSD *int64 `bson:"Volume USD" json:"Volume USD"`
}
