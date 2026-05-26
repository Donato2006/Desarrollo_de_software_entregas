package domain

// Usuario representa a las personas registradas en la plataforma (clientes o administradores)
type Usuario struct {
	ID       uint   `gorm:"primaryKey"`
	Correo   string `gorm:"unique;not null"`
	Rol      string `gorm:"not null"`
	Password string `gorm:"not null"`
	Nombre   string `gorm:"not null"`
}
