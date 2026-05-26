package domain

import "time"

// Concierto representa la entidad de un evento musical en el sistema
type Concierto struct {
	ID               uint      `gorm:"primaryKey"`
	Nombre           string    `gorm:"not null"`
	Fecha            time.Time `gorm:"not null"`
	Lugar            string    `gorm:"not null"`
	CupoTotal        int       `gorm:"not null"`
	CuposDisponibles int       `gorm:"not null"`
}
