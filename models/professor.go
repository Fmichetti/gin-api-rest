package models

import "gorm.io/gorm"

type Professor struct {
	gorm.Model
	Nome           string `json:"nome"`
	Especializacao string `json:"especializacao"`
	Contato        string `json:"contato"`
	UserID         uint   `json:"user_id"`
}
