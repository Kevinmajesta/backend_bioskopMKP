package service

import (
	"errors"

	"Kevinmajesta/backend_bioskopMKP/internal/entity"
	"Kevinmajesta/backend_bioskopMKP/internal/repository"
	"Kevinmajesta/backend_bioskopMKP/pkg/token"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(name, email, password string) (*entity.User, error)
	Login(email, password string) (string, error)
}

type userService struct {
	userRepository repository.UserRepository
	tokenUseCase   token.TokenUseCase
}

func NewUserService(userRepository repository.UserRepository, tokenUseCase token.TokenUseCase) UserService {
	return &userService{
		userRepository: userRepository,
		tokenUseCase:   tokenUseCase,
	}
}

func (s *userService) Register(name, email, password string) (*entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := s.userRepository.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.userRepository.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("email/password yang anda masukkan salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("email/password yang anda masukkan salah")
	}

	claims := token.JwtCustomClaims{
		ID:    user.Id_user.String(),
		Email: user.Email,
	}

	accessToken, err := s.tokenUseCase.GenerateAccessToken(claims)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
