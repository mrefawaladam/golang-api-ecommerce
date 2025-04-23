package usecase

import (
	"ecommerce-api/internal/domain"
	"ecommerce-api/internal/dto/request"
	"ecommerce-api/internal/repository"
	"ecommerce-api/internal/util"
	"errors"
)

type UserUsecase interface {
	Register(input request.RegisterRequest) error
	Login(input request.LoginRequest) (string, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) Register(input request.RegisterRequest) error {
	hashedPassword, err := util.HashPassword(input.Password)
	if err != nil {
		return err
	}

	user := domain.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
	}

	return u.repo.Create(&user)
}

func (u *userUsecase) Login(input request.LoginRequest) (string, error) {
    // Cari user berdasarkan email
    user, err := u.repo.FindByEmail(input.Email)
    if err != nil {
        return "", errors.New("email atau password salah")
    }

    // Validasi password
    if !util.CheckPasswordHash(input.Password, user.Password) {
        return "", errors.New("email atau password salah")
    }

    // Generate token JWT
    token, err := util.GenerateToken(user.ID)
	if err != nil {
        return "", errors.New("gagal membuat token")
    }

    // Return token jika semua proses berhasil
    return token, nil
}
