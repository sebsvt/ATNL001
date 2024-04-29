package services

type AuthService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password string, hashedPassword string) bool
}
