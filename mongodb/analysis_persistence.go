package mongodb

import (
	"context"
	"data-engine/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type AnalysisPersistence struct {
	db                *mongo.Database
	walletsCollection *mongo.Collection
}

func NewMongoDBAnalysisPersistence(uri, dbName string) (*AnalysisPersistence, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)

	return &AnalysisPersistence{
		db:                db,
		walletsCollection: db.Collection("wallets"),
	}, nil
}

func (p *AnalysisPersistence) FetchAllIdWallet() ([]string, error) {
	ctx := context.Background()
	var walletIDs []string

	cursor, err := p.walletsCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var wallet entities.Wallet
		if err := cursor.Decode(&wallet); err != nil {
			log.Fatal(err)
		}
		walletIDs = append(walletIDs, wallet.ID)
	}

	return walletIDs, nil
}

func (p *AnalysisPersistence) FetchWalletLevelDB(walletID *string) (*entities.Wallet, error) {
	ctx := context.Background()
	var wallet entities.Wallet

	err := p.walletsCollection.FindOne(ctx, bson.M{"_id": walletID}).Decode(&wallet)
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (p *AnalysisPersistence) UpdateAnalysis(wallet *entities.Wallet) error {
	ctx := context.Background()

	opts := options.Update().SetUpsert(true)

	_, err := p.walletsCollection.UpdateOne(
		ctx,
		bson.M{"_id": wallet.ID},
		bson.M{"$set": wallet},
		opts, // Pasa las opciones a UpdateOne
	)

	return err
}
