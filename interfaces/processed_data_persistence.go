package interfaces

import (
	"github.com/jpandof/data-engine/entities/mars"
)

type ProcessedDataPersistence interface {
	AddTransactionPricedByWallets(wallets []*mars.Wallet) error
	SaveIdWallets(wallets []*string) error
}
