package interfaces

import "data-engine/entities"

type ProcessedDataPersistence interface {
	AddTransactionPricedByWallets(wallets []*entities.Wallet) error //AddTransactionsByLoteLevelDB(wallets []*entities.Wallet) error
	SaveIdWallets(wallets []*string) error                          //SaveAllWalletsLevelDB(wallets []*string) error

	// No use
	// FetchTx(emptyTx *entities.Tx) (*entities.Tx, error)
	// AddTransactionsLevelDB(wallet *entities.Wallet) error
	// SaveTxBatchLevelDB(wallets *map[string][]byte) error
	// SaveWalletIdDataLevelDB(id *string, data *[]byte) error
	// SaveWalletLevelDB(wallet *entities.Wallet) error
	// SaveWalletsBatchLevelDB(wallets *map[string][]byte) error
	// SaveWalletsLevelDB(wallets []*entities.Wallet) error
	// SaveMapTXLevelDB(wallet *entities.Wallet) error
	// SaveIdWalletLevelDB(idWallet string) error
	// ConvertWalletsToJsonInParallel(wallets []*entities.Wallet) map[string][]byte
	// SaveTx(tx *entities.Tx) error
	// GetWalletByID(walletID string) (*entities.Wallet, error)
	// AddTransactions(wallet *entities.Wallet) error
	// GetAllWalletIDs() ([]string, error)
	// UpdateAnalysisAndAddTransactions(wallet *entities.Wallet) error
}
