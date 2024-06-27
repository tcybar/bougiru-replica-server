package repository

import "github.com/tcybar/guild-battle-records-server/app/infrastracture/mysql/model"

type UserRepository interface {
	GetAll() ([]*model.User, error)
}
