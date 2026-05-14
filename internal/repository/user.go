package repository

import (
	"Kevinmajesta/backend_bioskopMKP/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByID(id uuid.UUID) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	CreateUser(user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindUserByID(id uuid.UUID) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.Where("id_user = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindUserByEmail(email string) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user *entity.User) error {
	return r.db.Create(user).Error
}
