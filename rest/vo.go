package rest

import (
	"github.com/HappyQuant/binance-kline-provider/models"
)

type klineVo struct {
	OpenTime   int64  `json:"open_time"`
	OpenPrice  string `json:"open_price"`
	ClosePrice string `json:"close_price"`
	HighPrice  string `json:"high_price"`
	LowPrice   string `json:"low_price"`
}

func newKlineVo(kline *models.Kline) *klineVo {
	return &klineVo{
		OpenTime:   kline.OpenTime,
		OpenPrice:  kline.OpenPrice.String(),
		ClosePrice: kline.ClosePrice.String(),
		HighPrice:  kline.HighPrice.String(),
		LowPrice:   kline.LowPrice.String(),
	}
}
