package sql

import (
	"context"
	"fmt"
	"github.com/InsuranceTech/shared/config"
	"github.com/InsuranceTech/shared/log"
	"github.com/InsuranceTech/shared/services/sql/model"
	"github.com/go-pg/pg/v10"
	"strconv"
)

var (
	cfg  *config.Config
	_log = log.CreateTag("Sql")
)

func SetConfig(cf *config.Config) {
	cfg = cf
}

func Ping() (bool, error) {
	conn := NewDBConn()
	defer conn.Close()
	err := conn.Ping(context.Background())
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetAllSymbols() ([]*model.Symbol, error) {
	conn := NewDBConn()
	defer conn.Close()
	var symbols = make([]*model.Symbol, 0)

	err := conn.Model(&symbols).
		Relation("BaseCoin").
		Relation("QuoteCoin").
		Select()

	if err != nil {
		_log.Error("GetAllSymbols", err)
		return nil, err
	}

	return symbols, nil
}

func GetAllIndicators() ([]*model.Indicator, error) {
	conn := NewDBConn()
	defer conn.Close()
	var indicators = make([]*model.Indicator, 0)

	err := conn.Model(&indicators).
		Relation("Outs").
		Select()

	if err != nil {
		_log.Error("GetAllIndicators", err)
		return nil, err
	}

	return indicators, nil
}

func NewDBConn() (con *pg.DB) {
	address := fmt.Sprintf("%s:%s", cfg.Postgresql.HOST, strconv.Itoa(cfg.Postgresql.PORT))
	options := &pg.Options{
		User:     cfg.Postgresql.USER,
		Password: cfg.Postgresql.PASS,
		Addr:     address,
		Database: cfg.Postgresql.DEFAULT_DB,
		OnConnect: func(ctx context.Context, conn *pg.Conn) error {
			if cfg.Postgresql.SCHEMA != "" {
				_, err := conn.Exec("set search_path=?", cfg.Postgresql.SCHEMA)
				if err != nil {
					_log.Fatal("NewDBConn.OnConnect", err)
				}
			}
			return nil
		},
	}
	con = pg.Connect(options)
	return con
}
