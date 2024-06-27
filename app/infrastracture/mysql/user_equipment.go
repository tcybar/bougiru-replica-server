package mysql

import (
	"database/sql"

	"github.com/tcybar/guild-battle-records-server/app/domain/entity"
)

type UserEquipmentRepository struct {
	db *sql.DB
}

func NewUserEquipmentRepository(db *sql.DB) *UserEquipmentRepository {
	return &UserEquipmentRepository{db: db}
}

func (r *UserEquipmentRepository) ListByUserID(userID int) ([]*entity.UserEquipment, error) {
	return []*entity.UserEquipment{
		{
			ID: 1,
		},
	}, nil
}
