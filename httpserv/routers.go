package httpserv

import (
	"github.com/NuttayotSukkum/user-profile/infrastructure"
	"github.com/NuttayotSukkum/user-profile/internal/adaptor/handler"
	"github.com/NuttayotSukkum/user-profile/internal/adaptor/repo"
	"github.com/NuttayotSukkum/user-profile/internal/core/service"
)

func bindHealthRouter(a *infrastructure.App) {

	a.Register(
		infrastructure.AppMethodGet,
		"/health",
		infrastructure.HealthHandler,
	)
}

func bindCustomerProfileRouteCreate(a *infrastructure.App) {
	userRepo := repo.NewUserProfileRepo(infrastructure.DB)
	svc := service.NewUserProfileSvc(userRepo)
	userProfileHandler := handler.NewUserHandler(svc)

	a.Register(
		infrastructure.AppMethodPost,
		"v1/user-profile/create",
		userProfileHandler.CreateUser,
	)
}
