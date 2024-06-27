package repository

import "github.com/tcybar/guild-battle-records-server/app/domain/entity"

type IUserEquipmentRepository interface {
	ListByUserID(userID int) ([]*entity.UserEquipment, error)
}
