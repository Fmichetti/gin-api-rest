package models

import "gorm.io/gorm"

type Questao struct {
	gorm.Model
	Enunciado    string        `json:"enunciado"`
	MateriaID    uint          `json:"materia_id"`
	Alternativas []Alternativa `gorm:"many2many:questao_alternativas;" json:"alternativas"`
}
