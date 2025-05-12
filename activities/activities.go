package activities

import (
	"binanceTemporal/entities"
	"binanceTemporal/sqlite"
	"context"

	"github.com/adshao/go-binance/v2"
	_ "github.com/mattn/go-sqlite3"
)

var db sqlite.Sqlite // global pointer

func SetDB(database sqlite.Sqlite) {
	db = database
}

func SaveToDB(event binance.WsKlineEvent) error {
	events := entities.KlineEvent{
		Symbol:         event.Symbol,
		StartTime:      event.Kline.StartTime,
		EndTime:        event.Kline.EndTime,
		Interval:       event.Kline.Interval,
		OpenPrice:      event.Kline.Open,
		ClosePrice:     event.Kline.Close,
		HighPrice:      event.Kline.High,
		LowPrice:       event.Kline.Low,
		BaseVolume:     event.Kline.Volume,
		NumberOfTrades: event.Kline.TradeNum,
		QuoteVolume:    event.Kline.QuoteVolume,
	}
	return db.AddKlineEvent(context.TODO(), &events)
}
