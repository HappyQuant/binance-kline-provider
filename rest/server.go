package rest

import (
	"github.com/HappyQuant/binance-kline-provider/conf"
	"github.com/HappyQuant/binance-kline-provider/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type HttpServer struct {
	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{
		addr: addr,
	}
}

func (server *HttpServer) Start() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()
	r.Use(cors.Default())

	err := r.Run(server.addr)
	if err != nil {
		panic(err)
	}
}

func StartServer() {
	config := conf.GetConfig()
	httpServer := NewHttpServer(config.RestServer.Addr)
	go httpServer.Start()

	logger.Info("rest server ok")
}
