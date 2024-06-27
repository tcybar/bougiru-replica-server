package usecase

import (
	"context"
	"fmt"

	"github.com/tcybar/guild-battle-records-server/app/domain/entity"
	"github.com/tcybar/guild-battle-records-server/app/domain/repository"
	"github.com/tcybar/guild-battle-records-server/app/pkg/sliceutil"
)

// equipmentUsecase 装備に関するユースケースの構造体
type equipmentUsecase struct {
	equipmentRepo     repository.IEquipmentRepository
	userEquipmentRepo repository.IUserEquipmentRepository
}

// IEquipmentUsecase 装備に関するユースケースのインターフェース
type IEquipmentUsecase interface {
	GetUserEquipmentList(ctx context.Context, userID int) ([]*entity.UserEquipment, error)
}

// NewEquipmentUsecase 装備に関するユースケースを作成する
func NewEquipmentUsecase(equipmentRepo repository.IEquipmentRepository, userEquipmentRepo repository.IUserEquipmentRepository) *equipmentUsecase {
	return &equipmentUsecase{equipmentRepo: equipmentRepo, userEquipmentRepo: userEquipmentRepo}
}

// GetUserEquipmentList ユーザーの装備一覧を取得する
func (u *equipmentUsecase) GetUserEquipmentList(ctx context.Context, userID int) ([]*entity.UserEquipment, error) {
	// 装備マスタ取得
	equipments, err := u.equipmentRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	userEquipmentMapByID := sliceutil.ToMapByKey(equipments, func(e *entity.Equipment) int { return e.ID })

	// ユーザーの装備取得
	userEquipment, err := u.userEquipmentRepo.ListByUserID(userID)
	if err != nil {
		return nil, err
	}
	fmt.Println(userEquipment)

	for _, equipment := range equipments {
		if userEquipment, ok := userEquipmentMapByID[equipment.ID]; ok {
			fmt.Println(userEquipment)
		}
	}

	return nil, nil
}
