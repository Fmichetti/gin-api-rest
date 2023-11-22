package models

// AlunoResponse representa a estrutura desejada para a resposta JSON
type AlunoResponse struct {
	Matricula string `json:"matricula" gorm:"unique"`
	Nome      string `json:"nome"`
	Turma     string `json:"turma"`
}
