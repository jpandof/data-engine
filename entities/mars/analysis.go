package mars

import "time"

type TypeTrader int

const (
	Scalper TypeTrader = iota
	DayTrader
	Swing
	PositionTrader
	Undefined
)

type Timestamp struct {
	time.Time
}

type Analysis struct {
	CoinsInWallet                    float64    `json:"coins_in_wallet" bson:"coins_in_wallet"`
	CoinsBuy                         float64    `json:"coins_buy" bson:"coins_buy"`
	CoinsSell                        float64    `json:"coins_sell" bson:"coins_sell"`
	WeightedAveragePrice             float64    `json:"weighted_average_price" bson:"weighted_average_price"`
	MoneySpentToBuy                  float64    `json:"money_spent_to_buy" bson:"money_spent_to_buy"`
	MoneyEarnedToSell                float64    `json:"money_earned_to_sell" bson:"money_earned_to_sell"`
	MoneyCapitalMax                  float64    `json:"money_capital_max" bson:"money_capital_max"`
	TxInterleaved                    int64      `json:"tx_interleaved" bson:"tx_interleaved"`
	PercentageInterleaved            float64    `json:"percentage_tx_interleaved" bson:"percentage_tx_interleaved"`
	ProfitabilityByTotal             float64    `json:"profitability_by_total" bson:"profitability_by_total"`
	Profitability                    float64    `json:"profitability" bson:"profitability"`
	ProfitabilityMonth               float64    `json:"profitability_month" bson:"profitability_month"`
	ProfitabilityByTransaction       float64    `json:"profitability_by_transaction" bson:"profitability_by_transaction"`
	PercentageProfitabilityByTotal   float64    `json:"percentage_profitability_by_total" bson:"percentage_profitability_by_total"`
	PercentageTransactionsWithProfit float64    `json:"percentage_transactions_with_profit" bson:"percentage_transactions_with_profit"`
	TypeTrader                       TypeTrader `json:"type_trader" bson:"type_trader"`
	AverageTimeBetweenBuyAndSell     int64      `json:"average_time_between_buy_and_sell" bson:"average_time_between_buy_and_sell"`
	NumberTransactionsXLastTime      float64    `json:"number_transactions_x_last_time" bson:"number_transactions_x_last_time"`
	QuantityTransactionsTotal        int        `json:"quantity_transactions_total" bson:"quantity_transactions_total"`
	QuantityTransactionsBuy          int        `json:"quantity_transactions_buy" bson:"quantity_transactions_buy"`
	QuantityTransactionsSell         int        `json:"quantity_transactions_sell" bson:"quantity_transactions_sell"`
	RatioTransactionsSellOverBuy     float64    `json:"ratio_transactions_sell_over_buy" bson:"ratio_transactions_sell_over_buy"`
	MaxLost                          float64    `json:"max_lost" bson:"max_lost"`
	MaxProfit                        float64    `json:"max_profit" bson:"max_profit"`
	MaxLostPercentage                float64    `json:"max_lost_percentage" bson:"max_lost_percentage"`
	MaxProfitPercentage              float64    `json:"max_profit_percentage" bson:"max_profit_percentage"`
	AverageLost                      float64    `json:"average_lost" bson:"average_lost"`
	AverageProfit                    float64    `json:"average_profit" bson:"average_profit"`
	DateLastBuy                      Timestamp  `json:"date_last_buy" bson:"date_last_buy"`
	DateLastSell                     Timestamp  `json:"date_last_sell" bson:"date_last_sell"`
	DateAnalysis                     Timestamp  `json:"date_analysis" bson:"date_analysis"`
	DateMinTrade                     Timestamp  `json:"date_min_trade" bson:"date_min_trade"`
}
