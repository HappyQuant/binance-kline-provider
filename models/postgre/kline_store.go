package postgre

import (
	"fmt"
	"github.com/HappyQuant/binance-kline-provider/models"
)

func KlineTableName(symbol, interval string) string {
	return fmt.Sprintf("t_kline_%s_%s", symbol, interval)
}

func (s *Store) GetPreviousKlines(symbol, interval string, endTime int64, limit int) ([]*models.Kline, error) {
	db := s.db.Table(KlineTableName(symbol, interval)).Where("open_time <= ?", endTime).Order("open_time DESC").Limit(limit)
	var klines []*models.Kline
	err := db.Find(&klines).Error
	return klines, err
}

func (s *Store) GetNextKlines(symbol, interval string, fromTime int64, limit int) ([]*models.Kline, error) {
	db := s.db.Table(KlineTableName(symbol, interval)).Where("open_time >= ?", fromTime).Order("open_time ASC").Limit(limit)
	var klines []*models.Kline
	err := db.Find(&klines).Error
	return klines, err
}
