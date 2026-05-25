package database

import (
	"backend/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:12345@tcp(127.0.0.1:3306)/conciertos_db?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("No se pudo conectar a la base de datos")
	}

	database.AutoMigrate(
		&domain.Usuario{},
		&domain.Concierto{},
		&domain.Entrada{},
		&domain.ListaEspera{},
	)
	DB = database
}
