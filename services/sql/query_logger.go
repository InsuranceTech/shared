package sql

import (
	"context"
	"github.com/InsuranceTech/shared/log"
	"github.com/go-pg/pg/v10"
)

type QueryLogger struct {
}

var (
	dblog = log.CreateTag("DB Query")
)

func (d QueryLogger) BeforeQuery(c context.Context, e *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d QueryLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	cmd, _ := q.FormattedQuery()
	dblog.Log(string(cmd), q.Result)
	return nil
}
