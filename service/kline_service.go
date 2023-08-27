package service

import (
	"github.com/HappyQuant/binance-kline-provider/models"
	"github.com/HappyQuant/binance-kline-provider/models/postgre"
)

func GetPreviousKlines(symbol, interval string, endTime int64, limit int) ([]*models.Kline, error) {
	return postgre.SharedStore().GetPreviousKlines(symbol, interval, endTime, limit)
}

func GetNextKlines(symbol, interval string, fromTime int64, limit int) ([]*models.Kline, error) {
	return postgre.SharedStore().GetNextKlines(symbol, interval, fromTime, limit)
}
