package mysql

import (
	"context"
	"database/sql"

	"github.com/tcybar/guild-battle-records-server/app/domain/entity"
	"github.com/tcybar/guild-battle-records-server/app/infrastracture/mysql/model"
)

type UserCharacterRepository struct {
	db *sql.DB
}

func NewUserCharacterRepository(db *sql.DB) *UserCharacterRepository {
	return &UserCharacterRepository{db: db}
}

func (r *UserCharacterRepository) ListByUserID(ctx context.Context, userID int) ([]*entity.UserCharacter, error) {
	userCharacters, err := model.UserCharacters(model.UserCharacterWhere.UserID.EQ(uint64(userID))).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return toUserCharacterEntities(userCharacters), nil
}

func (r *UserCharacterRepository) Create(userCharacter *entity.UserCharacter) error {
	_, err := r.db.Exec(
		"INSERT INTO user_characters (user_id, character_id, level, exp) VALUES (?, ?, ?, ?)",
		userCharacter.UserID, userCharacter.CharacterID, userCharacter.Level, userCharacter.Exp,
	)
	return err
}

// func toUserCharacterModel(e *entity.UserCharacter) *model.UserCharacter {
// 	return &model.UserCharacter{
// 		UserID:      uint64(e.UserID),
// 		CharacterID: uint64(e.CharacterID),
// 		Level:       e.Level,
// 		Exp:         e.Exp,
// 	}
// }

func toUserCharacterEntity(m *model.UserCharacter) *entity.UserCharacter {
	return &entity.UserCharacter{
		UserID:      int(m.UserID),
		CharacterID: int32(m.CharacterID),
		Level:       m.Level,
		Exp:         m.Exp,
	}
}

func toUserCharacterEntities(ms []*model.UserCharacter) []*entity.UserCharacter {
	userCharacters := make([]*entity.UserCharacter, 0, len(ms))
	for _, m := range ms {
		userCharacters = append(userCharacters, toUserCharacterEntity(m))
	}
	return userCharacters
}
