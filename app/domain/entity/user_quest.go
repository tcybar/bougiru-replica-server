package entity

// UserQuest ユーザーのクエスト進捗状況を表す
type UserQuest struct {
	ID      int  // ユーザーのクエスト進捗状況ID
	UserID  int  // ユーザーID
	QuestID int  // クエストID
	IsClear bool // クエストをクリアしたかどうか
}

// NewUserQuest 新規作成
func NewUserQuest(userID int, questID int, isClear bool) *UserQuest {
	return &UserQuest{
		UserID:  userID,
		QuestID: questID,
		IsClear: isClear,
	}
}
