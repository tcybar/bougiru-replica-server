.PHONY: setup
## setup: 開発環境のセットアップ
setup:
	@bash ./scripts/setup.sh

.PHONY: mysql-cli
## mysql-cli: mysqlクライアントを起動
mysql-cli:
	docker compose run mysql-cli

.PHONY: sqlite-cli
## sqlite-cli: sqliteクライアントを起動
sqlite-cli:
	sqlite3 ./db/master.db

.PHONY: server
## server: サーバーを起動
server:
	@go run ./cmd/server/main.go

.PHONY: create-masterdb
## create-masterdb: マスタデータベースを作成
create-masterdb:
	@rm -f ./db/master.db
	sqlite3 ./db/master.db < ./schema/master/schema.sql
	sqlite3 ./db/master.db < ./schema/master/seed.sql

.PHONY: create-mysqldb
## create-mysqldb: mysqlデータベースを作成
create-mysqldb:
	@mysqldef -uuser -ppassword db < ./schema/schema.sql

.PHONY: sqlboiler-userdb
## sqlboiler-userdb: sqlboilerを実行
sqlboiler-userdb:
	@sqlboiler mysql -o app/infrastracture/mysql/model -p model --wipe --no-tests

.PHONY: sqlboiler-masterdb
## sqlboiler-masterdb: sqlboilerを実行
sqlboiler-masterdb:
	@sqlboiler sqlite3 -o app/infrastracture/sqlite/model -p model --wipe --no-tests

.PHONY: openapi
## openapi: apiのリクエストレスポンスコードを自動生成
openapi:
  ## guild_battle_records_api
	docker run --rm -v ${PWD}:/local -u $(shell id -u):$(shell id -g) openapitools/openapi-generator-cli:v7.1.0 generate \
		-g go-server \
		-i /local/api/guild_battle_records_api.yaml \
		-o /local/app/gen \
    --additional-properties=onlyInterfaces=true,outputAsLibrary=true,packageName=guildapi,router=chi,sourceFolder=guildapi
  ## 不要なimport文を削除しないとコンパイルエラーになるため、goimportsでimport文を整形する
	goimports -w ./app/gen
