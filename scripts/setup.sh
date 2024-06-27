#!/bin/sh
set -eu

# Install golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $GOBIN v1.55.2

# Install goimports for openapi-generatorで生成したコードの不要なimportを削除するため
# golangci-lintではtypecheckエラーになって、importが削除されない
go install golang.org/x/tools/cmd/goimports@latest

# Install gopls for VSCode
# VSCodeでセーブ時に自動フォーマットするため
go install golang.org/x/tools/gopls@latest

# Install staticcheck for VSCode
# セーブ時にコンパイルエラーに気付くため
go install honnef.co/go/tools/cmd/staticcheck@latest

# Install delve for VSCode
# VSCodeでデバッグするため
go install github.com/go-delve/delve/cmd/dlv@latest

# sqldef
curl -OL https://github.com/k0kubun/sqldef/releases/download/v0.17.11/mysqldef_linux_amd64.tar.gz
tar xf mysqldef_linux_amd64.tar.gz -C ${GOBIN}
rm mysqldef_linux_amd64.tar.gz

# sqlBoiler
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@latest
