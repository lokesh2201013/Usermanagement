package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Age       int     `json:"age" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	Createdby uuid.UUID `json:"created_by"`
    UpdatedAt time.Time`json:"updated_at"`
	UpdatedBy  uuid.UUID `json:"updated_by"`
}

type Admin struct{
    ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"` 
    Name string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`	
}