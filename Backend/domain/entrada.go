package domain

import "time"

type Entrada struct {
	ID uint `gorm:"primaryKey"`

	UsuarioID uint `gorm:"not null"`
	Usuario   Usuario

	ConciertoID uint `gorm:"not null"`
	Concierto   Concierto

	Estado string `gorm:"default:'activa'"`

	FechaCompra time.Time `gorm:"autoCreateTime"`
}
