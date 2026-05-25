package services

import (
	"backend/dao"
	"backend/domain"
)

// ObtenerTodosLosConciertos busca todos los registros de la tabla conciertos
func ObtenerTodosLosConciertos() ([]domain.Concierto, error) {

	var conciertos []domain.Concierto
	// Ejecuta un "SELECT * FROM conciertos" y llena el slice con los resultados
	resultado := dao.DB.Find(&conciertos)
	// Devuelve el array obtenido y el error si es que hubo alguno
	return conciertos, resultado.Error
}

// CrearConcierto inserta un nuevo registro de concierto en la BD
func CrearConcierto(concierto domain.Concierto) (domain.Concierto, error) {
	// Ejecuta un "INSERT INTO" usando los datos de la estructura
	resultado := dao.DB.Create(&concierto)
	// Retorna el concierto (ahora con su ID autoincremental asignado) y el error
	return concierto, resultado.Error
}

// ObtenerConciertoPorID busca un concierto que coincida exactamente con la clave primaria dada
func ObtenerConciertoPorID(id string) (domain.Concierto, error) {

	var concierto domain.Concierto
	// Busca el primer registro que coincida con el ID enviado. Lanza error si no existe.
	resultado := dao.DB.First(&concierto, id)

	return concierto, resultado.Error
}

// ActualizarConcierto modifica un registro existente combinando búsqueda y actualización
func ActualizarConcierto(id string, datos domain.Concierto) (domain.Concierto, error) {

	var concierto domain.Concierto
	// 1. Verifica primero si el concierto existe en la BD
	resultado := dao.DB.First(&concierto, id)

	if resultado.Error != nil {
		return concierto, resultado.Error // Retorna temprano si no se encuentra
	}
	// 2. Si existe, actualiza solo los campos que cambiaron usando el método Updates de GORM
	dao.DB.Model(&concierto).Updates(datos)

	return concierto, nil
}

// EliminarConcierto borra un registro de concierto según su ID
func EliminarConcierto(id string) error {

	var concierto domain.Concierto
	// 1. Verifica si el concierto existe antes de intentar borrarlo
	resultado := dao.DB.First(&concierto, id)

	if resultado.Error != nil {
		return resultado.Error // Si no existe, devuelve el error (como ErrRecordNotFound)
	}
	// 2. Si existe, lo remueve de la base de datos
	dao.DB.Delete(&concierto)

	return nil
}
