package entities

type KlineEvent struct {
	ID             int64  `bun:",pk,autoincrement"`
	Symbol         string `bun:"symbol"`
	StartTime      int64  `bun:"start_time"`
	EndTime        int64  `bun:"end_time"`
	Interval       string `bun:"interval"`
	OpenPrice      string `bun:"open_price"`
	ClosePrice     string `bun:"close_price"`
	HighPrice      string `bun:"high_price"`
	LowPrice       string `bun:"low_price"`
	BaseVolume     string `bun:"base_volume"`
	NumberOfTrades int64  `bun:"number_of_trades"`
	QuoteVolume    string `bun:"quote_volume"`
}
