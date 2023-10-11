package models

import "gorm.io/gorm"

type Materia struct {
	gorm.Model
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
}
