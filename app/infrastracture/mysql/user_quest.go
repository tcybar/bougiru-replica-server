package mysql

import (
	"database/sql"

	"github.com/tcybar/guild-battle-records-server/app/domain/entity"
)

type UserQuestRepository struct {
	db *sql.DB
}

func NewUserQuestRepository(db *sql.DB) *UserQuestRepository {
	return &UserQuestRepository{db: db}
}

func (r *UserQuestRepository) GetListByUserID(userID int) ([]*entity.UserQuest, error) {
	return []*entity.UserQuest{
		{
			ID: 1,
		},
	}, nil
}
