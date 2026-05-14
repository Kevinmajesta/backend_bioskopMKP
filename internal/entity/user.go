package entity

import (
	"github.com/google/uuid"
)

type User struct {
	Id_user  uuid.UUID `json:"id_user" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"-"`
	Auditable
}
