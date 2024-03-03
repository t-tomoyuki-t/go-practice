package repository

type IAuthRepository interface {
	GetSession() error
	SetSession() error
	DeleteSession() error
}
