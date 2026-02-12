package service

import (
	"strings"
	"time"

	"github.com/NuttayotSukkum/user-profile/internal/core/domain"
	"github.com/NuttayotSukkum/user-profile/internal/core/port"
	"github.com/NuttayotSukkum/user-profile/internal/utils"
	"github.com/google/uuid"
)

type UserProfileSvc struct {
	userProfileRepo port.UserProfileRepo
}

func NewUserProfileSvc(userProfileRepo port.UserProfileRepo) *UserProfileSvc {
	return &UserProfileSvc{
		userProfileRepo: userProfileRepo,
	}
}

func (repo *UserProfileSvc) CreateUserProfile(req port.UserProFileRequest) error {
	hashPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := domain.UserProfile{
		ID:        uuid.NewString(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashPass,
		Role:      strings.ToUpper("user"),
		Status:    domain.StatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := repo.userProfileRepo.Save(user); err != nil {
		return err
	}

	return nil
}
