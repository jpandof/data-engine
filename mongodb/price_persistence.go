package mongodb

import (
	"context"
	"fmt"
	price2 "github.com/jpandof/data-engine/entities/price"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoPricePersistence struct {
	db              *mongo.Database
	priceCollection *mongo.Collection
}

func NewMongoPricePersistence(uri, dbName string) (*MongoPricePersistence, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)

	return &MongoPricePersistence{
		db:              db,
		priceCollection: db.Collection("prices"),
	}, nil
}

func (m MongoPricePersistence) FetchPriceOrAverage(timestamp int64) (*int64, error) {
	t := time.Unix(timestamp, 0)
	roundedTimestamp := t.Truncate(time.Minute).Unix()

	price, err := m.GetPrice(roundedTimestamp)
	if err != nil {
		// return nil, err
	}

	if price != nil {
		return price.Close, nil
	}

	return m.getAveragePrice(roundedTimestamp)

}

func (m MongoPricePersistence) GetPrice(timestamp int64) (*price2.Price, error) {
	var price *price2.Price

	filter := bson.M{"_id": timestamp}

	err := m.priceCollection.FindOne(context.Background(), filter).Decode(&price)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no price found for timestamp: %d", timestamp)
		}
		return nil, err
	}

	return price, nil
}

func (m MongoPricePersistence) getAveragePrice(timestamp int64) (*int64, error) {
	opts := options.FindOne().SetSort(bson.D{{"_id", 1}})
	var nextPriceMongo, prevPriceMongo *price2.Price

	result := m.priceCollection.FindOne(context.Background(), bson.M{"_id": bson.M{"$gt": timestamp}}, opts)
	errNext := result.Decode(&nextPriceMongo)
	opts.SetSort(bson.D{{"_id", -1}})
	errPrev := m.priceCollection.FindOne(context.Background(), bson.M{"_id": bson.M{"$lt": timestamp}}, opts).Decode(&prevPriceMongo)

	if errNext != nil && errPrev != nil {
		return nil, errNext
	}

	// Si se encuentra solo uno de los dos precios, devuelve ese
	if errNext != nil {
		return prevPriceMongo.Close, nil
	}
	if errPrev != nil {
		return nextPriceMongo.Close, nil
	}

	// Si se encuentran ambos, calcula y devuelve la media
	averagePrice := int64((*prevPriceMongo.Close + *nextPriceMongo.Close) / 2)
	return &averagePrice, nil
}

func (m MongoPricePersistence) toInt64(price *price2.Price) (*int64, error) {
	floatValue := *price.Close

	// Convertir el valor float64 a int64
	intValue := int64(floatValue)

	// Retornar un puntero al valor int64
	return &intValue, nil
}
