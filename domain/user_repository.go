package domain

type UserRepository interface {
	FindByUsernameAndPassword(username, password string) (*User, error)
}
