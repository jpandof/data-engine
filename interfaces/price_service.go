package interfaces

type PriceService interface {
	FetchPrice(timestamp int64) (int64, error)
}
