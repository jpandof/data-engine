package mongodb

import (
	"context"
	"github.com/jpandof/data-engine/entities/mars"
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
		walletsCollection:      db.Collection("wallets-id"),
		transactionsCollection: db.Collection("transactions"),
	}, nil
}

func (p *ProcessedDataPersistence) AddTransactionPricedByWallets(wallets []*mars.Wallet) error {

	ctx := context.Background()

	var bulkOps []mongo.WriteModel
	for _, wallet := range wallets {
		for _, tx := range wallet.Transactions {
			// Aquí asumo que tienes un campo único, como un ID, en tus transacciones.
			filter := bson.M{"_id": tx.Id}

			// En este ejemplo, 'tx' es el documento que quieres insertar o actualizar.
			update := mongo.NewUpdateOneModel().
				SetFilter(filter).
				SetUpdate(bson.M{"$set": tx}).
				SetUpsert(true) // Esto es clave: crea el documento si no existe.

			bulkOps = append(bulkOps, update)
		}
	}

	if len(bulkOps) > 0 {
		opts := options.BulkWrite().SetOrdered(false)
		_, err := p.transactionsCollection.BulkWrite(ctx, bulkOps, opts)
		if err != nil {
			return err
		}
	}

	return nil

}

/*func (p *ProcessedDataPersistence) SaveTransactions(wallets []*mars.Wallet) error {
	ctx := context.Background()

	// Prepare bulk write options
	opts := options.BulkWrite().SetOrdered(false) // Set to 'false' to allow operations to be processed in parallel

	var operations []mongo.WriteModel
	for _, wallet := range wallets {
		for _, tx := range wallet.Transactions {
			// Asigna un filtro para buscar la transacción basada en un identificador único, como _id
			filter := bson.M{"_id": tx.Id}
			// Asigna la operación con upsert=true. Esto actualizará el documento si existe, o lo insertará si no.
			upsertModel := mongo.NewUpdateOneModel().
				SetFilter(filter).
				SetUpdate(bson.M{"$set": tx}).
				SetUpsert(true)
			operations = append(operations, upsertModel)
		}
	}

	// Ejecuta todas las operaciones como un BulkWrite
	if len(operations) > 0 {
		_, err := p.transactionsCollection.BulkWrite(ctx, operations, opts)
		if err != nil {
			return err
		}
	}

	return nil
}*/

func (p *ProcessedDataPersistence) SaveIdWallets(wallets []*string) error {
	ctx := context.Background()
	opts := options.BulkWrite().SetOrdered(false)

	var operations []mongo.WriteModel
	for _, walletID := range wallets {
		// Cada operación es un modelo de actualización con upsert verdadero.
		upsertModel := mongo.NewUpdateOneModel().
			SetFilter(bson.M{"_id": *walletID}). // Filtrar por el ID de la wallet.
			SetUpdate(bson.D{{"$set", bson.M{"_id": *walletID}}}).
			SetUpsert(true) // Configurar upsert como verdadero.
		operations = append(operations, upsertModel)
	}

	// Realizar las operaciones en un BulkWrite para manejar todos los upserts.
	if len(operations) > 0 {
		_, err := p.walletsCollection.BulkWrite(ctx, operations, opts)
		if err != nil {
			return err
		}
	}

	return nil
}
