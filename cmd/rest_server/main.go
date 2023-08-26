package main

import (
	"github.com/HappyQuant/binance-kline-provider/rest"
	_ "net/http/pprof"
)

func main() {
	rest.StartServer()
	select {}
}
