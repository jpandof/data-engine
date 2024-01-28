package mongodb

import (
	"context"
	"github.com/jpandof/data-engine/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProcessedDataPersistence struct {
	client                 *mongo.Client
	db                     *mongo.Database
	walletsCollection      *mongo.Collection
	transactionsCollection *mongo.Collection
}

func NewMongoDBProcessedDataPersistence(uri, dbName string) (*ProcessedDataPersistence, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)

	return &ProcessedDataPersistence{
		client:                 client,
		db:                     db,
		walletsCollection:      db.Collection("wallets"),
		transactionsCollection: db.Collection("transactions"),
	}, nil
}

func (p *ProcessedDataPersistence) AddTransactionPricedByWallets(wallets []*entities.Wallet) error {
	ctx := context.Background()

	var transactionDocs []interface{}
	for _, wallet := range wallets {
		for _, tx := range wallet.Transactions {
			transactionDocs = append(transactionDocs, tx)
		}
	}

	if len(transactionDocs) > 0 {
		_, err := p.transactionsCollection.InsertMany(ctx, transactionDocs)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *ProcessedDataPersistence) SaveIdWallets(wallets []*string) error {
	ctx := context.Background()

	var walletDocs []interface{}
	for _, walletID := range wallets {
		walletDocs = append(walletDocs, bson.M{"_id": *walletID})
	}

	if len(walletDocs) > 0 {
		_, err := p.walletsCollection.InsertMany(ctx, walletDocs)
		if err != nil {
			return err
		}
	}

	return nil
}
