package services

import (
	"backend/dao"
	"backend/domain"
)

func ObtenerTodosLosConciertos() ([]domain.Concierto, error) {

	var conciertos []domain.Concierto

	resultado := dao.DB.Find(&conciertos)

	return conciertos, resultado.Error
}
func CrearConcierto(concierto domain.Concierto) (domain.Concierto, error) {

	resultado := dao.DB.Create(&concierto)

	return concierto, resultado.Error
}
func ObtenerConciertoPorID(id string) (domain.Concierto, error) {

	var concierto domain.Concierto

	resultado := dao.DB.First(&concierto, id)

	return concierto, resultado.Error
}
func ActualizarConcierto(id string, datos domain.Concierto) (domain.Concierto, error) {

	var concierto domain.Concierto

	resultado := dao.DB.First(&concierto, id)

	if resultado.Error != nil {
		return concierto, resultado.Error
	}

	dao.DB.Model(&concierto).Updates(datos)

	return concierto, nil
}
func EliminarConcierto(id string) error {

	var concierto domain.Concierto

	resultado := dao.DB.First(&concierto, id)

	if resultado.Error != nil {
		return resultado.Error
	}

	dao.DB.Delete(&concierto)

	return nil
}
