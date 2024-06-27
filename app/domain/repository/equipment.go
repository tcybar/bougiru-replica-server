package repository

import (
	"context"

	"github.com/tcybar/guild-battle-records-server/app/domain/entity"
)

type IEquipmentRepository interface {
	ListAll(ctx context.Context) ([]*entity.Equipment, error)
}
