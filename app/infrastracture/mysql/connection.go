package mysql

import (
	"database/sql"

	"github.com/volatiletech/sqlboiler/boil"
)

// NewMysqlClient Mysqlデータベースへの接続設定
func NewMysqlClient() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/db?parseTime=true")
	if err != nil {
		return nil, err
	}

	// sqlboiler用の設定
	boil.SetDB(db)

	return db, nil
}
