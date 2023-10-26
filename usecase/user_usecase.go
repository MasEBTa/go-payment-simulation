package usecase

import (
	"errors"
	"go-payment-simulation/model"
	"go-payment-simulation/repository"
	"go-payment-simulation/utils/common"
	"go-payment-simulation/utils/security"
)

type UserUsecase interface {
	FindById(id string) (model.UserCredential, error)
	Login(payload model.UserCredential) (model.UserCredential, error)
	FindByRole(role string) ([]model.UserCredential, error)
	NewUser(payload model.UserCredential) (model.UserCredential, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

// FindByRole implements UserUsecase.
func (u *userUsecase) FindByRole(role string) ([]model.UserCredential, error) {
	return u.repo.FindByRole(role)
}

// Login implements UserUsecase.
func (u *userUsecase) Login(payload model.UserCredential) (model.UserCredential, error) {
	var user model.UserCredential
	var err error
	if payload.Email != "" {
		user, err = u.repo.FindByEmail(payload.Email)
		if err != nil {
			return model.UserCredential{}, err
		}
	}
	if payload.Username != "" {
		user, err = u.repo.FIndByUsername(payload.Username)
		if err != nil {
			return model.UserCredential{}, err
		}
	}
	// Validasi Password
	err = security.VerifyPassword(user.Password, payload.Password)
	if err != nil {
		return model.UserCredential{}, err
	}
	// Generate Token
	token, err := security.GenerateJwtToken(user)
	if err != nil {
		return model.UserCredential{}, err
	}
	return model.UserCredential{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
		UserRole: user.UserRole,
		Password: "*********",
	}, nil
}

// NewUser implements UserUsecase.
func (u *userUsecase) NewUser(payload model.UserCredential) (model.UserCredential, error) {
	// cek apakah email sudah digunakan
	user, _ := u.repo.FindByEmail(payload.Email)
	if user.Email != "" {
		return model.UserCredential{}, errors.New("email is used")
	}
	// cek apakah username sudah digunakan
	user, _ = u.repo.FIndByUsername(payload.Username)
	if user.Username != "" {
		return model.UserCredential{}, errors.New("username is used")
	}
	// hashing password
	hashPassword, err := security.HashPassword(payload.Password)
	payload.Password = hashPassword
	if err != nil {
		return model.UserCredential{}, err
	}
	// buat id
	payload.Id = common.GenerateID()

	return u.repo.NewUser(payload)
}

// FindById implements UserUsecase.
func (u *userUsecase) FindById(id string) (model.UserCredential, error) {
	return u.repo.FindById(id)
}

func NewUserUseCase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}
