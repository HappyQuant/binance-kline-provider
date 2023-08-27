package postgre

import (
	"fmt"
	"github.com/HappyQuant/binance-kline-provider/conf"
	"github.com/HappyQuant/binance-kline-provider/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
	"sync"
)

var gdb *gorm.DB
var store models.Store
var storeOnce sync.Once

type Store struct {
	db *gorm.DB
}

func SharedStore() models.Store {
	storeOnce.Do(func() {
		err := initDb()
		if err != nil {
			panic(err)
		}
		store = NewStore(gdb)
	})
	return store
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func initDb() error {
	cfg := conf.GetConfig()

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		strings.Split(cfg.DataSource.Addr, ":")[0], cfg.DataSource.User, cfg.DataSource.Password,
		cfg.DataSource.Database, strings.Split(cfg.DataSource.Addr, ":")[1])

	var err error
	gdb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) BeginTx() (models.Store, error) {
	db := s.db.Begin()
	if db.Error != nil {
		return nil, db.Error
	}
	return NewStore(db), nil
}

func (s *Store) Rollback() error {
	return s.db.Rollback().Error
}

func (s *Store) CommitTx() error {
	return s.db.Commit().Error
}
