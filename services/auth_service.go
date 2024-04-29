package services

import (
	"github.com/sebsvt/ATNL001/logs"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
}

func NewAuthService() AuthService {
	return authService{}
}

func (srv authService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	return string(hashedPassword), nil
}
func (srv authService) VerifyPassword(password string, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		logs.Error(err)
		return false
	}
	return true
}
