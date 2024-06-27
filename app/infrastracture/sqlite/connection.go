package sqlite

import (
	"log"

	"github.com/jmoiron/sqlx"
)

// NewMysqlClient Mysqlデータベースへの接続設定
func NewSqliteClient() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "./db/master.db")
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
