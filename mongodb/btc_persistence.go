package mongodb

import (
	"context"
	"data-engine/entities"
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

func (m *BTCPersistence) FetchTx(tx *entities.Tx) (*entities.Tx, error) {
	filter := bson.M{"_id": tx.Id}
	result := m.txCollection.FindOne(context.Background(), filter)

	var fetchedTx entities.Tx
	err := result.Decode(&fetchedTx)
	if err != nil {
		return nil, err
	}

	return &fetchedTx, nil
}

func (m *BTCPersistence) SaveTx(txs []*entities.Tx) error {

	// MongoDB permite operaciones de inserción en batch, lo que es más eficiente
	// que insertar cada transacción individualmente
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
		_, err := m.txCollection.BulkWrite(context.Background(), models)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *BTCPersistence) SaveBlkfile(blkfile *entities.Blkfile) error {
	_, err := m.blkfilesCollection.InsertOne(context.Background(), blkfile)
	return err
}

func (m *BTCPersistence) SaveBlock(block *entities.Block) error {
	_, err := m.blockCollection.InsertOne(context.Background(), block)
	return err
}

func (m *BTCPersistence) IsBlkProcessed(blkfile *entities.Blkfile) (*bool, error) {
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

func (m *BTCPersistence) IsBlockProcessed(block *entities.Block) (*bool, error) {

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
