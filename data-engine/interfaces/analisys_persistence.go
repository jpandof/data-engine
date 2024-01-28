package interfaces

import "data-engine/entities"

type AnalysisPersistence interface {
	FetchAllIdWallet() ([]string, error) // GetAllWalletIDsLevelDB(id string) ([]string, error)
	FetchWalletLevelDB(walletID *string) (*entities.Wallet, error)
	UpdateAnalysis(wallet *entities.Wallet) error
}
