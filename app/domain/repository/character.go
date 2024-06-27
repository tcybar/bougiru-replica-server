package repository

import (
	"context"

	"github.com/tcybar/guild-battle-records-server/app/domain/entity"
)

type ICharacterRepository interface {
	Find(ctx context.Context, id int32) (*entity.Character, error)
	ListAll(ctx context.Context) ([]*entity.Character, error)
	ListByIDs(ctx context.Context, ids []int32) ([]*entity.Character, error)
}
