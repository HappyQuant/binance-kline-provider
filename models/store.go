package models

type Store interface {
	BeginTx() (Store, error)
	Rollback() error
	CommitTx() error

	GetPreviousKlines(symbol, interval string, endTime int64, limit int) ([]*Kline, error)
	GetNextKlines(symbol, interval string, fromTime int64, limit int) ([]*Kline, error)
}
