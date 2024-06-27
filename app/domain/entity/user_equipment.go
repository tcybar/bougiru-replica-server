package entity

// UserEquipment ユーザーの装備
type UserEquipment struct {
	ID          int // ユーザーの装備ID
	UserID      int // ユーザーID
	EquipmentID int // 装備ID
}

// NewUserEquipment 新規作成
func NewUserEquipment(userID int, equipmentID int) *UserEquipment {
	return &UserEquipment{
		UserID:      userID,
		EquipmentID: equipmentID,
	}
}
