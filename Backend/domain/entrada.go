package domain

import "time"

// Entrada representa un ticket o boleto comprado por un usuario para un concierto
type Entrada struct {
	ID          uint `gorm:"primaryKey"`
	UsuarioID   uint
	ConciertoID uint
	Estado      string    `gorm:"not null"`
	FechaCompra time.Time `gorm:"not null"`
	// Relaciones de GORM (Carga los objetos completos en las consultas si se usa Preload)
	Usuario   Usuario
	Concierto Concierto
}
