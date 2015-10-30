package repository

type UserRepository struct {

}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) IsValid(userName, password string) bool {
	return true
}