package sqlite

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/tcybar/guild-battle-records-server/app/domain/entity"
	"github.com/tcybar/guild-battle-records-server/app/infrastracture/sqlite/model"
	"github.com/tcybar/guild-battle-records-server/app/pkg/sliceutil"
)

// CharacterRepository キャラクターマスタのリポジトリ
type CharacterRepository struct {
	db *sqlx.DB
}

// NewCharacterRepository キャラクターマスタのリポジトリを作成する
func NewCharacterRepository(db *sqlx.DB) *CharacterRepository {
	return &CharacterRepository{db: db}
}

// Find 主キーを指定してキャラクターマスタを取得する
func (r *CharacterRepository) Find(ctx context.Context, id int32) (*entity.Character, error) {
	character, err := model.Characters(model.CharacterWhere.ID.EQ(int64(id))).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return toCharacterEntity(character), nil
}

func (r *CharacterRepository) ListAll(ctx context.Context) ([]*entity.Character, error) {
	characters, err := model.Characters().All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	fmt.Println(characters)

	return nil, nil
}

// ListByIDs IDsからキャラクターマスタを取得する
func (r *CharacterRepository) ListByIDs(ctx context.Context, ids []int32) ([]*entity.Character, error) {
	ids64 := sliceutil.ConvertNumberSlice[int32, int64](ids)
	characters, err := model.Characters(model.CharacterWhere.ID.IN(ids64)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return toCharacterEntities(characters), nil
}

// toCharacterEntity モデルをエンティティに変換する
func toCharacterEntity(c *model.Character) *entity.Character {
	return &entity.Character{
		ID:   int32(c.ID),
		Name: c.Name,
	}
}

// toCharacterEntities モデルスライスをエンティティスライスに変換する
func toCharacterEntities(ms []*model.Character) []*entity.Character {
	characters := make([]*entity.Character, 0, len(ms))
	for _, m := range ms {
		characters = append(characters, toCharacterEntity(m))
	}
	return characters
}
