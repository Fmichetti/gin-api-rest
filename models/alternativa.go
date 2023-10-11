package models

import "gorm.io/gorm"

type Alternativa struct {
	gorm.Model
	Texto    string `json:"texto"`
	Resposta bool   `json:"resposta"`
}
