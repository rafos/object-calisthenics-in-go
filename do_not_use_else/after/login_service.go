package after

import (
	"github.com/rafos/object-calisthenics-in-go/do_not_use_else/repository"
)

type loginService struct {
	userRepository *repository.UserRepository
}

func NewLoginService() *loginService {
	return &loginService{
		userRepository: repository.NewUserRepository(),
	}
}

func (l *loginService) Login(userName, password string) {
	if l.userRepository.IsValid(userName, password) {
		redirect("homepage")
		return
	}

	addFlash("error", "Bad credentials")

	redirect("login")
}

func redirect(page string) {
	// redirect to given page
}

func addFlash(msgType, msg string) {
	// create flash message
}
