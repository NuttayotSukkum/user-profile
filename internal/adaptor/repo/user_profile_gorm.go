package repo

import (
	"errors"

	"github.com/NuttayotSukkum/user-profile/internal/core/domain"
	"gorm.io/gorm"
)

type UserProfileRepo struct {
	db *gorm.DB
}

func NewUserProfileRepo(db *gorm.DB) *UserProfileRepo {
	return &UserProfileRepo{db: db}
}

func (repo *UserProfileRepo) Save(user domain.UserProfile) error {
	if err := repo.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserProfileRepo) FindUserById(id string) (*domain.UserProfile, error) {
	var user domain.UserProfile
	tx := repo.db.
		Where("id = ? AND status <> ?", id, domain.StatusInactive).
		First(&user)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}

	return &user, nil
}
