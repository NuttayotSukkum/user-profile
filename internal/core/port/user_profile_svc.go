package port

type UserProFileRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"  validate:"required"`
	Email     string `json:"email"      validate:"required,email"`
	Password  string `json:"password"   validate:"required,min=8"`
}

type UserProfileSvc interface {
	CreateUserProfile(req UserProFileRequest) error
}
