package models

import "gorm.io/gorm"

type Avaliacao struct {
	gorm.Model
	Nome      string    `json:"nome"`
	Descricao string    `json:"descricao"`
	MateriaID uint      `json:"materia_id"`
	TurmaID   uint      `json:"turma_id"`
	Questoes  []Questao `gorm:"many2many:avaliacao_questoes;" json:"questoes"`
}
