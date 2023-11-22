package models

import (
	"time"

	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome           string    `json:"nome"`
	Matricula      string    `json:"matricula"`
	DataNascimento time.Time `json:"data_nascimento"`
	UserID         uint      `json:"user_id" gorm:"unique"`
	TurmaID        uint      `json:"turma_id"`
	Turma          Turma     `gorm:"foreignKey:TurmaID"`
}
