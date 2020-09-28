package repository

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

// PGConnect function will create a connection to the database
func PGConnect(connectionURL string) *pg.DB {
	opt, err := pg.ParseURL(connectionURL)
	if err != nil {
		panic(err)
	}

	return pg.Connect(opt)
}

// DBLogger ..
type DBLogger struct{}

// BeforeQuery ..
func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

// AfterQuery ..
func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.Query)
	fmt.Println(q.Result)
	return nil
}
