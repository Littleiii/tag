package model

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var MyDB *sqlx.DB
func intDB() error  {
	MyDB =  sqlx.MustOpen("mysql","root:root@tcp(localhost:3306)/test?parseTime=True&loc=Local&multiStatements=true&charset=utf8mb4")
	ctx,cancel := context.WithTimeout(context.Background(),time.Second*3)
	defer cancel()
	MyDB.PingContext(ctx)
	return MyDB.PingContext(ctx)
}
