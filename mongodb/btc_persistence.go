package mongodb

import (
	"context"
	"github.com/jpandof/data-engine/entities/btc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BTCPersistence struct {
	client             *mongo.Client
	db                 *mongo.Database   // Agregando una referencia a la base de datos
	blockCollection    *mongo.Collection // Agregar un campo para la colección 'blocks'
	txCollection       *mongo.Collection // Agregar un campo para la colección 'blocks'
	blkfilesCollection *mongo.Collection // Agregar un campo para la colección 'blocks'

}

func NewMongoBTCPersistence(uri, dbName string) (*BTCPersistence, error) {
	// Conectar al cliente de MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Obtener la referencia a la base de datos
	db := client.Database(dbName)

	return &BTCPersistence{
		client:             client,
		db:                 db,
		blockCollection:    db.Collection("blocks"),
		txCollection:       db.Collection("txs"),
		blkfilesCollection: db.Collection("blk-files"),
	}, nil
}

func (m *BTCPersistence) FetchTx(tx *btc.Tx) (*btc.Tx, error) {
	filter := bson.M{"_id": tx.Id}
	result := m.txCollection.FindOne(context.Background(), filter)

	if result.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}

	var fetchedTx btc.Tx
	err := result.Decode(&fetchedTx)
	if err != nil {
		return nil, err
	}

	return &fetchedTx, nil
}

func (m *BTCPersistence) SaveTx(txs []*btc.Tx) error {

	// MongoDB permite operaciones de inserción en batch, lo que es más eficiente
	// que insertar cada transacción individualmente
	opts := options.BulkWrite().SetOrdered(false)
	var models []mongo.WriteModel
	for _, tx := range txs {
		// Crear un modelo de upsert para cada transacción
		// Esto actualizará la transacción si existe (basado en el ID) o la insertará si no existe
		model := mongo.NewReplaceOneModel().
			SetFilter(bson.M{"_id": tx.Id}).
			SetReplacement(tx).
			SetUpsert(true)
		models = append(models, model)
	}

	// Realizar la operación de inserción en batch
	if len(models) > 0 {
		_, err := m.txCollection.BulkWrite(context.Background(), models, opts)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *BTCPersistence) SaveBlkfile(blkfile *btc.Blkfile) error {

	filter := bson.M{"_id": blkfile.GetId()}
	update := bson.M{"$setOnInsert": *blkfile}

	_, err := m.blkfilesCollection.UpdateOne(
		context.Background(),
		filter,
		update,
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *BTCPersistence) SaveBlock(block *btc.Block) error {
	filter := bson.M{"_id": block.GetId()}
	update := bson.M{"$setOnInsert": *block}

	_, err := m.blockCollection.UpdateOne(
		context.Background(),
		filter,
		update,
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *BTCPersistence) IsBlkProcessed(blkfile *btc.Blkfile) (*bool, error) {
	filter := bson.M{"_id": blkfile.Id}
	result := m.blkfilesCollection.FindOne(context.Background(), filter)

	if result.Err() == mongo.ErrNoDocuments {
		falseVal := false
		return &falseVal, nil
	}
	if result.Err() != nil {
		return nil, result.Err()
	}

	trueVal := true
	return &trueVal, nil
}

func (m *BTCPersistence) IsBlockProcessed(block *btc.Block) (*bool, error) {

	filter := bson.M{"_id": block.Id}
	result := m.blockCollection.FindOne(context.Background(), filter)

	if result.Err() == mongo.ErrNoDocuments {
		falseVal := false
		return &falseVal, nil
	}
	if result.Err() != nil {
		return nil, result.Err()
	}

	trueVal := true
	return &trueVal, nil
}
