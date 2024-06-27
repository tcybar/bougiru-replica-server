package handler

import (
	"context"
	"net/http"

	"github.com/tcybar/guild-battle-records-server/app/gen/guildapi"
	"github.com/tcybar/guild-battle-records-server/app/usecase"
)

// CharacterHandler キャラクター関連のAPIハンドラ
type CharacterHandler struct {
	iCharacterUsecase usecase.ICharacterUsecase
}

// NewCharacterHandler キャラクター関連のAPIハンドラを作成
func NewCharacterHandler(iCharacterUsecase usecase.ICharacterUsecase) *CharacterHandler {
	return &CharacterHandler{iCharacterUsecase: iCharacterUsecase}
}

// GetCharacters キャラクター情報一覧取得
func (h *CharacterHandler) GetCharacters(ctx context.Context) (guildapi.ImplResponse, error) {
	getUserCharacterListOutput, err := h.iCharacterUsecase.GetUserCharacterList(ctx, 1)
	if err != nil {
		return guildapi.ImplResponse{}, err
	}

	characterList := make([]int32, 0, len(getUserCharacterListOutput.List))
	for _, v := range getUserCharacterListOutput.List {
		characterList = append(characterList, v.ID)
	}

	res := guildapi.GetCharactersResponse{
		CharacterList: characterList,
	}

	return guildapi.ImplResponse{Code: http.StatusOK, Body: res}, nil
}

// PostCharacters キャラクター登録
func (h *CharacterHandler) PostCharacters(ctx context.Context, req guildapi.PostCharactersRequest) (guildapi.ImplResponse, error) {
	userCharacter, err := h.iCharacterUsecase.CreateUserCharacter(ctx, 1, req.CharacterId)
	if err != nil {
		return guildapi.ImplResponse{}, err
	}

	return guildapi.ImplResponse{Code: http.StatusOK, Body: userCharacter}, nil
}
