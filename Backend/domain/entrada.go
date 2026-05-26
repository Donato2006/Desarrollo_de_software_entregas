package domain

import "time"

// Entrada representa un ticket o boleto comprado por un usuario para un concierto
type Entrada struct {
	ID uint `gorm:"primaryKey"`

	UsuarioID uint `gorm:"not null"`
	Usuario   Usuario

	ConciertoID uint `gorm:"not null"`
	Concierto   Concierto

	Estado string `gorm:"default:'activa'"`

	FechaCompra time.Time `gorm:"autoCreateTime"`
}
