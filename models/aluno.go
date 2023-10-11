package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome           string `json:"nome"`
	DataNascimento string `json:"data_nascimento"`
	UserID         uint   `json:"user_id"`
	TurmaID        uint   `json:"turma_id"`
}
