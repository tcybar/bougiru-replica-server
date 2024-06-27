package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tcybar/guild-battle-records-server/app/gen/guildapi"
	"github.com/tcybar/guild-battle-records-server/app/usecase"
)

// QuestHandler クエスト関連のAPIハンドラ
type QuestHandler struct {
	iQuestUsecase usecase.IQuestUsecase
}

// NewQuestHandler クエスト関連のAPIハンドラを作成
func NewQuestHandler(iQuestUsecase usecase.IQuestUsecase) *QuestHandler {
	return &QuestHandler{iQuestUsecase: iQuestUsecase}
}

// V1QuestsGet クエスト情報一覧取得
func (h *QuestHandler) V1QuestsGet(ctx context.Context) (guildapi.ImplResponse, error) {
	ms, err := h.iQuestUsecase.GetUserQuestList(ctx, 1)
	if err != nil {
		return guildapi.ImplResponse{}, err
	}
	fmt.Println(ms)

	return guildapi.ImplResponse{Code: http.StatusOK, Body: nil}, nil
}
