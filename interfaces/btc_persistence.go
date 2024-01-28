package interfaces

import "data-engine/entities"

type BTCPersistence interface {
	FetchTx(tx *entities.Tx) (*entities.Tx, error) // FetchTxLevelDB(tx *entities.Tx) (*entities.Tx, error)
	SaveTx(txs []*entities.Tx) error               // SaveTxLevelDB(mapTx map[string][]byte) error
	SaveBlkfile(blkfile *entities.Blkfile) error
	SaveBlock(block *entities.Block) error
	IsBlkProcessed(blkfile *entities.Blkfile) (*bool, error)
	IsBlockProcessed(block *entities.Block) (*bool, error)
}
