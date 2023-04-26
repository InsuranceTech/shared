package sql

import (
	"context"
	"fmt"
	"github.com/InsuranceTech/shared/config"
	"github.com/InsuranceTech/shared/services/sql/model"
	"github.com/go-pg/pg/v10"
	"strconv"
)

var (
	cfg *config.Config
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

	err := pg.Model(&symbols).
		Relation("BaseCoin").
		Relation("QuoteCoin").
		Select()

	if err != nil {
		return nil, err
	}

	return symbols, nil
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
					panic(err.Error())
				}
			}
			return nil
		},
	}
	con = pg.Connect(options)
	return con
}
