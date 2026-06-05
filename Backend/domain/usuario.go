package domain

// Usuario representa a las personas registradas en la plataforma (clientes o administradores)
type Usuario struct {
	ID       uint   `gorm:"primaryKey"`
	Correo   string `gorm:"unique;not null" binding:"required,email"`
	Rol      string `gorm:"not null"`
	Password string `gorm:"not null" binding:"required,min=6"`
	Nombre   string `gorm:"not null" binding:"required"`
}
