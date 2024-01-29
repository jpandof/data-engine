package interfaces

import (
	"github.com/jpandof/data-engine/entities/btc"
)

type BTCPersistence interface {
	FetchTx(tx *btc.Tx) (*btc.Tx, error)
	SaveTx(txs []*btc.Tx) error
	SaveBlkfile(blkfile *btc.Blkfile) error
	SaveBlock(block *btc.Block) error
	IsBlkProcessed(blkfile *btc.Blkfile) (*bool, error)
	IsBlockProcessed(block *btc.Block) (*bool, error)
}
