package usecase

import (
	"context"
	"errors"

	"github.com/tcybar/guild-battle-records-server/app/domain/entity"
	"github.com/tcybar/guild-battle-records-server/app/domain/repository"
	"github.com/tcybar/guild-battle-records-server/app/pkg/sliceutil"
)

// characterUsecase キャラクターに関するユースケースの構造体
type characterUsecase struct {
	characterRepo     repository.ICharacterRepository
	userCharacterRepo repository.IUserCharacterRepository
}

// ICharacterUsecase キャラクターに関するユースケースのインターフェース
type ICharacterUsecase interface {
	GetUserCharacterList(ctx context.Context, userID int) (*GetUserCharacterListOutput, error)
	CreateUserCharacter(ctx context.Context, userID int, characterID int32) (*entity.UserCharacter, error)
}

// GetUserCharacterList ユーザーのキャラクター一覧を取得する
type GetUserCharacterListOutput struct {
	List []*UserCharacterOutput
}
type UserCharacterOutput struct {
	ID   int32
	Name string
}

// NewCharacterUsecase キャラクターに関するユースケースを作成する
func NewCharacterUsecase(characterRepo repository.ICharacterRepository, userCharacterRepo repository.IUserCharacterRepository) *characterUsecase {
	return &characterUsecase{characterRepo: characterRepo, userCharacterRepo: userCharacterRepo}
}

func (u *characterUsecase) GetUserCharacterList(ctx context.Context, userID int) (*GetUserCharacterListOutput, error) {
	// ユーザーのキャラクター取得
	userCharacters, err := u.userCharacterRepo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// キャラクターマスタ取得
	characterIDs := sliceutil.ConvertSlice(userCharacters, func(e *entity.UserCharacter) int32 { return e.CharacterID })
	uniqueCharacterIDs := sliceutil.ConvertUniqueSlice(characterIDs)
	characters, err := u.characterRepo.ListByIDs(ctx, uniqueCharacterIDs)
	if err != nil {
		return nil, err
	}
	characterMapByID := sliceutil.ToMapByKey(characters, func(e *entity.Character) int32 { return e.ID })

	userCharacterOutputList := make([]*UserCharacterOutput, 0, len(userCharacters))
	for _, userCharacter := range userCharacters {
		if character, ok := characterMapByID[userCharacter.CharacterID]; ok {
			userCharacterOutputList = append(userCharacterOutputList, &UserCharacterOutput{
				ID:   character.ID,
				Name: character.Name,
			})
		}
	}

	// 出力
	output := &GetUserCharacterListOutput{
		List: userCharacterOutputList,
	}

	return output, nil
}

// CreateUserCharacter ユーザーのキャラクターを登録する
func (u *characterUsecase) CreateUserCharacter(ctx context.Context, userID int, characterID int32) (*entity.UserCharacter, error) {
	// キャラクターマスタ取得
	character, err := u.characterRepo.Find(ctx, characterID)
	if err != nil {
		return nil, err
	}

	// キャラクターマスタが存在しない場合はエラー
	if character == nil {
		return nil, errors.New("character not found")
	}

	// ユーザーのキャラクターを新規作成
	userCharacter := entity.NewUserCharacter(userID, character.ID)
	if err = u.userCharacterRepo.Create(userCharacter); err != nil {
		return nil, err
	}

	return userCharacter, nil
}
