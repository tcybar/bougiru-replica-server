package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tcybar/guild-battle-records-server/app/gen/guildapi"
	"github.com/tcybar/guild-battle-records-server/app/handler"
	"github.com/tcybar/guild-battle-records-server/app/infrastracture/mysql"
	"github.com/tcybar/guild-battle-records-server/app/infrastracture/sqlite"
	"github.com/tcybar/guild-battle-records-server/app/usecase"

	_ "github.com/go-sql-driver/mysql"
)

type Weapon struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func main() {
	// Sqlite接続[マスタデータ用]
	sqliteClient, err := sqlite.NewSqliteClient()
	if err != nil {
		log.Fatal(err)
	}
	defer sqliteClient.Close()

	// Mysql接続[ユーザーデータ用]
	mysqlClient, err := mysql.NewMysqlClient()
	if err != nil {
		log.Fatal(err)
	}
	defer mysqlClient.Close()

	// MySqlリポジトリ定義
	userCharacterRepo := mysql.NewUserCharacterRepository(mysqlClient)
	userQuestRepo := mysql.NewUserQuestRepository(mysqlClient)

	// Sqliteリポジトリ定義
	characterRepo := sqlite.NewCharacterRepository(sqliteClient)
	questRepo := sqlite.NewQuestRepository(sqliteClient)

	// ユースケース定義
	characterUsecase := usecase.NewCharacterUsecase(characterRepo, userCharacterRepo)
	questUsecase := usecase.NewQuestUsecase(questRepo, userQuestRepo)

	// ハンドラ定義
	characterApiHandler := handler.NewCharacterHandler(characterUsecase)
	questApiHandler := handler.NewQuestHandler(questUsecase)

	// コントローラ定義
	characterApiController := guildapi.NewCharacterAPIController(characterApiHandler)
	questApiController := guildapi.NewQuestAPIController(questApiHandler)

	// ルーティング定義
	router := guildapi.NewRouter(characterApiController, questApiController)

	// APIサーバ起動
	http.ListenAndServe(":8080", router)
}
