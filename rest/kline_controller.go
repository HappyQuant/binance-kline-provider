package rest

import (
	"fmt"
	"github.com/HappyQuant/binance-kline-provider/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GET /api/kline/<symbol>/<interval>/previous?endTime=1503360000000&&limit=1000
func GetPreviousKlines(ctx *gin.Context) {
	symbol := ctx.Param("symbol")
	if len(symbol) > 16 {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("invalid symbol"))
		return
	}

	interval := ctx.Param("interval")
	if len(interval) > 8 {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("invalid interval"))
		return
	}

	endTime, err := strconv.ParseInt(ctx.DefaultQuery("endTime", "2147483647000"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("strconv.ParseInt(endTime): %v", err))
		return
	}

	limit, err := strconv.ParseInt(ctx.DefaultQuery("limit", "1000"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("strconv.ParseInt(limit): %v", err))
	}
	if limit < 0 {
		limit = 1
	} else if limit > 1000 {
		limit = 1000
	}

	var klineVos []*klineVo
	klines, _ := service.GetPreviousKlines(symbol, interval, endTime, int(limit))
	for _, kline := range klines {
		klineVos = append(klineVos, newKlineVo(kline))
	}

	ctx.JSON(http.StatusOK, klineVos)
}

// GET /api/kline/<symbol>/<interval>/next?fromTime=1503360000000&&limit=1000
func GetNextKlines(ctx *gin.Context) {
	symbol := ctx.Param("symbol")
	if len(symbol) > 16 {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("invalid symbol"))
		return
	}

	interval := ctx.Param("interval")
	if len(interval) > 8 {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("invalid interval"))
		return
	}

	fromTime, err := strconv.ParseInt(ctx.DefaultQuery("fromTime", "0"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("strconv.ParseInt(fromTime): %v", err))
		return
	}

	limit, err := strconv.ParseInt(ctx.DefaultQuery("limit", "1000"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("strconv.ParseInt(limit): %v", err))
	}
	if limit < 0 {
		limit = 1
	} else if limit > 1000 {
		limit = 1000
	}

	var klineVos []*klineVo
	klines, _ := service.GetNextKlines(symbol, interval, fromTime, int(limit))
	for _, kline := range klines {
		klineVos = append(klineVos, newKlineVo(kline))
	}

	ctx.JSON(http.StatusOK, klineVos)
}
