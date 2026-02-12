package port

import "github.com/NuttayotSukkum/user-profile/internal/core/domain"

type UserProfileRepo interface {
	Save(user domain.UserProfile) error
	FindUserById(id string) (*domain.UserProfile, error)
	// findUserAll() error
	// UpdateUserProfile(UserProfile domain.UserProfile) error
}
