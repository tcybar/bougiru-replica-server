package entity

// UserCharacter ユーザーのキャラクターを表す
type UserCharacter struct {
	ID          int   // ユーザーキャラクターID
	UserID      int   // ユーザーID
	CharacterID int32 // キャラクターID
	Level       int   // レベル
	Exp         int   // 経験値
}

const (
	userCharacterInitialLevel = 1 // 初期レベル
)

// NewUserCharacter 新規作成
func NewUserCharacter(userID int, characterID int32) *UserCharacter {
	return &UserCharacter{
		ID:          0, // NOTE: IDはDBで自動採番されるため、0で初期化
		UserID:      userID,
		CharacterID: characterID,
		Level:       userCharacterInitialLevel,
		Exp:         0,
	}
}
