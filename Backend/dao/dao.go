package dao

import (
	"backend/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB es la instancia global de GORM que compartiremos con el resto de la app
var DB *gorm.DB

// Connect inicializa la conexión a MySQL y realiza la migración de las tablas
func Connect() {
	// DSN (Data Source Name): Contiene las credenciales y configuración de la base de datos
	dsn := "root:12345@tcp(127.0.0.1:3306)/conciertos_db?charset=utf8mb4&parseTime=True&loc=Local"
	// Intenta abrir la conexión con el driver de MySQL y la configuración por defecto de GORM
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// Si la conexión falla, detiene la ejecución del programa de inmediato
	if err != nil {
		panic("No se pudo conectar a la base de datos")
	}
	// AutoMigrate crea automáticamente las tablas en la BD basándose en las estructuras de Go
	database.AutoMigrate(
		&domain.Usuario{},
		&domain.Concierto{},
		&domain.Entrada{},
		&domain.ListaEspera{},
	)
	// Asigna la base de datos ya conectada a nuestra variable global
	DB = database
}
