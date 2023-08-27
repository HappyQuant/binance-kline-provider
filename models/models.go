package models

import (
	"github.com/shopspring/decimal"
)

type Kline struct {
	OpenTime            int64 `gorm:"column:open_time;primary_key"`
	CloseTime           int64
	OpenPrice           decimal.Decimal `sql:"type:decimal(32,16);"`
	ClosePrice          decimal.Decimal `sql:"type:decimal(32,16);"`
	HighPrice           decimal.Decimal `sql:"type:decimal(32,16);"`
	LowPrice            decimal.Decimal `sql:"type:decimal(32,16);"`
	BaseVolume          decimal.Decimal `sql:"type:decimal(32,16);"`
	QuoteVolume         decimal.Decimal `sql:"type:decimal(32,16);"`
	TradesCount         int32
	TakerBuyBaseVolume  decimal.Decimal `sql:"type:decimal(32,16);"`
	TakerBuyQuoteVolume decimal.Decimal `sql:"type:decimal(32,16);"`
	reserved            string
}
