# ==============================
# ベースイメージ
# ==============================
ARG GO_VERSION=1.22.4
FROM golang:${GO_VERSION}-bookworm

# ==============================
# 必要なパッケージのインストール
# direnv            環境変数設定のため
# bash-completion   makeコマンド実行時のTab補完を有効にするため
# sqlite3           マスタデータをsqliteで管理するため
# ==============================
RUN apt-get update && apt-get install -y \
    direnv \
    bash-completion \
    sqlite3

# ==============================
# vscodeからログインするためのユーザーを設定
# ==============================
ARG USERNAME=tcybar
ARG GROUPNAME=tcybar
ARG UID=1000
ARG GID=1000
RUN groupadd -g ${GID} ${GROUPNAME} && \
    useradd -m -s /bin/bash -u ${UID} -g ${GID} $USERNAME

# ==============================
# 作業ディレクトリの設定
# ==============================
WORKDIR /home/${USERNAME}/workspace

# ==============================
# direnvの設定
# ==============================
RUN echo eval "$(direnv hook bash)" >> /home/${USERNAME}/.bashrc && \
    . /home/${USERNAME}/.bashrc

# ==============================
# エントリーポイント
# ==============================
ENTRYPOINT ["/bin/bash"]
