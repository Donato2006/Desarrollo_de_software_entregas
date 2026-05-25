package domain

import "time"

type ListaEspera struct {
	ID                uint `gorm:"primaryKey"`
	UsuarioID         uint
	ConciertoID       uint
	PosicionCola      int       `gorm:"not null"`
	Estado            string    `gorm:"not null"`
	FechaAlta         time.Time `gorm:"not null"`
	FechaNotificacion *time.Time
	FechaExpiracion   *time.Time

	Usuario   Usuario
	Concierto Concierto
}
