package services

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/sebsvt/ATNL001/logs"
	"github.com/sebsvt/ATNL001/repositories"
)

var (
	ErrUserEmailAlreadyInUse = errors.New("email already in use")
	ErrUserDoesNotExists     = errors.New("this user account does not exists")
)

type userService struct {
	userRepo repositories.UserRepository
	authSrv  AuthService
}

func NewUserService(userRepo repositories.UserRepository, authSrv AuthService) UserService {
	return userService{userRepo: userRepo, authSrv: authSrv}
}

func (srv userService) CreateUserAccount(newUser CreateNewUserRequest) (*UserResposne, error) {

	// checking email is already in used
	_, err := srv.userRepo.FromEmail(newUser.Email)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return nil, ErrUserDoesNotExists
		}
		return nil, err
	}

	// hashing password
	hashed, err := srv.authSrv.HashPassword(newUser.Password)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	newUser.Password = hashed

	// create new user
	user, err := srv.userRepo.CreateNewUser(repositories.User{
		Email:          newUser.Email,
		FirstName:      newUser.FirstName,
		LastName:       newUser.LastName,
		HashedPassword: newUser.Password,
		CreatedAt:      time.Now().Format("2006-1-2 15:04:05"),
	})

	if err != nil {
		logs.Error(err)
		return nil, err
	}

	// returing data
	return &UserResposne{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}
func (srv userService) GetUser(id int) (*UserResposne, error) {
	user, err := srv.userRepo.FromID(id)
	if err != nil {
		logs.Error(err)
		if err == sql.ErrNoRows {
			return nil, ErrUserDoesNotExists
		}
		return nil, err
	}

	return &UserResposne{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}
