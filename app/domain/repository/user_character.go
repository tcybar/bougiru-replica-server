package repository

import (
	"context"

	"github.com/tcybar/guild-battle-records-server/app/domain/entity"
)

type IUserCharacterRepository interface {
	ListByUserID(ctx context.Context, userID int) ([]*entity.UserCharacter, error)
	Create(userCharacter *entity.UserCharacter) error
}
