package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tcybar/guild-battle-records-server/app/gen/guildapi"
	"github.com/tcybar/guild-battle-records-server/app/usecase"
)

// EquipmentHandler 装備関連のAPIハンドラ
type EquipmentHandler struct {
	iEquipmentUsecase usecase.IEquipmentUsecase
}

// NewEquipmentHandler 装備関連のAPIハンドラを作成
func NewEquipmentHandler(iEquipmentUsecase usecase.IEquipmentUsecase) *EquipmentHandler {
	return &EquipmentHandler{iEquipmentUsecase: iEquipmentUsecase}
}

// V1EquipmentGet 装備情報一覧取得
func (h *EquipmentHandler) V1EquipmentGet(ctx context.Context) (guildapi.ImplResponse, error) {
	ms, err := h.iEquipmentUsecase.GetUserEquipmentList(ctx, 1)
	if err != nil {
		return guildapi.ImplResponse{}, err
	}
	fmt.Println(ms)

	return guildapi.ImplResponse{Code: http.StatusOK, Body: nil}, nil
}
