package conf

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type Config struct {
	DataSource DataSourceConfig `json:"dataSource"`
	RestServer RestServerConfig `json:"restServer"`
}

type DataSourceConfig struct {
	DriverName string `json:"driverName"`
	Addr       string `json:"addr"`
	Database   string `json:"database"`
	User       string `json:"user"`
	Password   string `json:"password"`
}

type RestServerConfig struct {
	Addr string `json:"addr"`
}

var config Config
var configOnce sync.Once

func GetConfig() *Config {
	configOnce.Do(func() {
		bytes, err := ioutil.ReadFile("conf.json")
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &config)
		if err != nil {
			panic(err)
		}
	})
	return &config
}
