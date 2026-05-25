package domain

import "time"

type Entrada struct {
	ID          uint `gorm:"primaryKey"`
	UsuarioID   uint
	ConciertoID uint
	Estado      string    `gorm:"not null"`
	FechaCompra time.Time `gorm:"not null"`

	Usuario   Usuario
	Concierto Concierto
}
