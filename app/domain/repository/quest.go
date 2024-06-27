package repository

import (
	"context"

	"github.com/tcybar/guild-battle-records-server/app/domain/entity"
)

type IQuestRepository interface {
	ListAll(ctx context.Context) ([]*entity.Quest, error)
}
