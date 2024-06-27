package sqlite

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/tcybar/guild-battle-records-server/app/domain/entity"
	"github.com/tcybar/guild-battle-records-server/app/infrastracture/sqlite/model"
)

type QuestRepository struct {
	db *sqlx.DB
}

func NewQuestRepository(db *sqlx.DB) *QuestRepository {
	return &QuestRepository{db: db}
}

func (r *QuestRepository) ListAll(ctx context.Context) ([]*entity.Quest, error) {
	quests, err := model.Quests().All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	fmt.Println(quests)

	return nil, nil
}
