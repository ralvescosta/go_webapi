package app

type IUserService interface {
	Register()
}

type userService struct{}

func (s *userService) Register() {}

func NewUserService() IUserService {
	return &userService{}
}
