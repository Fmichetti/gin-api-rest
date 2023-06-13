package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	NAME        string    `json:"name"`
	CPF         string    `json:"cpf" gorm:"unique"`
	DESCRIPTION string    `json:"description"`
	CITY        string    `json:"city"`
}
