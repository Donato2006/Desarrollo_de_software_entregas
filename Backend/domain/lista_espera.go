package domain

import "time"

// ListaEspera gestiona la cola de usuarios que quieren ir a un concierto agotado
type ListaEspera struct {
	ID                uint `gorm:"primaryKey"`
	UsuarioID         uint
	ConciertoID       uint
	PosicionCola      int       `gorm:"not null"`
	Estado            string    `gorm:"not null"`
	FechaAlta         time.Time `gorm:"not null"`
	FechaNotificacion *time.Time
	FechaExpiracion   *time.Time
	// Relaciones de GORM
	Usuario   Usuario
	Concierto Concierto
}
