package interfaces

import (
	"github.com/jpandof/data-engine/entities/mars"
)

type AnalysisPersistence interface {
	FetchAllIdWallet() ([]string, error)
	FetchWallet(walletID *string) (*mars.Wallet, error)
	UpdateAnalysis(wallet *mars.Wallet) error
}
