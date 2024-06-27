package repository

import "github.com/tcybar/guild-battle-records-server/app/domain/entity"

type IUserQuestRepository interface {
	GetListByUserID(userID int) ([]*entity.UserQuest, error)
}
