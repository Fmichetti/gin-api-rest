package models

import "gorm.io/gorm"

type Turma struct {
	gorm.Model
	Nome string `json:"nome"`
}
