package domain

import "time"

type Concierto struct {
	ID               uint      `gorm:"primaryKey"`
	Nombre           string    `gorm:"not null"`
	Fecha            time.Time `gorm:"not null"`
	Lugar            string    `gorm:"not null"`
	CupoTotal        int       `gorm:"not null"`
	CuposDisponibles int       `gorm:"not null"`
}
