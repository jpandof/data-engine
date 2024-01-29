package btc

import "time"

type Blkfile struct {
	Id            string     `bson:"_id"`
	ProcessedTime *time.Time `bson:"processed_time"`
}

func NewBlkfile(id string) *Blkfile {
	return &Blkfile{Id: id}
}

func (bF *Blkfile) GetId() string {
	return bF.Id
}
