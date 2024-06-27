package usecase

import (
	"context"
	"fmt"

	"github.com/tcybar/guild-battle-records-server/app/domain/entity"
	"github.com/tcybar/guild-battle-records-server/app/domain/repository"
	"github.com/tcybar/guild-battle-records-server/app/pkg/sliceutil"
)

// questUsecase クエストに関するユースケースの構造体
type questUsecase struct {
	questRepo     repository.IQuestRepository
	userQuestRepo repository.IUserQuestRepository
}

// IQuestUsecase クエストに関するユースケースのインターフェース
type IQuestUsecase interface {
	GetUserQuestList(ctx context.Context, userID int) ([]*entity.UserQuest, error)
}

// NewQuestUsecase クエストに関するユースケースを作成する
func NewQuestUsecase(questRepo repository.IQuestRepository, userQuestRepo repository.IUserQuestRepository) *questUsecase {
	return &questUsecase{questRepo: questRepo, userQuestRepo: userQuestRepo}
}

// GetUserQuestList ユーザーのクエスト一覧を取得する
func (u *questUsecase) GetUserQuestList(ctx context.Context, userID int) ([]*entity.UserQuest, error) {
	// クエストマスタ取得
	quests, err := u.questRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	userQuestMapByID := sliceutil.ToMapByKey(quests, func(e *entity.Quest) int { return e.ID })

	// ユーザーのクエスト進捗状況取得
	userQuests, err := u.userQuestRepo.GetListByUserID(userID)
	if err != nil {
		return nil, err
	}
	fmt.Println(userQuests)

	for _, quest := range quests {
		if userQuest, ok := userQuestMapByID[quest.ID]; ok {
			fmt.Println(userQuest)
		}
	}

	return nil, nil
}
