package interfaces

type PricePersistence interface {
	FetchPriceOrAverage(timestamp int64) (*int64, error)
}
