package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	DIAGNOSIS   string    `json:"diagnosis"`
	ACTION      string    `json:"action"`
	DESCRIPTION string    `json:"description"`
	PATIENTID   uuid.UUID `gorm:"type:uuid;" json:"id"`
}
