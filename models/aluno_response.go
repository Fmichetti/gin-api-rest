package models

// AlunoResponse representa a estrutura desejada para a resposta JSON
type AlunoResponse struct {
	ID    uint   `json:"id"`
	Nome  string `json:"nome"`
	Turma string `json:"turma"`
}
