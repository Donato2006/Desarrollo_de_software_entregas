package services

import (
	"backend/dao"
	"backend/domain"
	"errors"
	"time"
)

func AnotarseListaEspera(usuarioID uint, conciertoID uint) (domain.ListaEspera, error) {
	var concierto domain.Concierto

	resultado := dao.DB.First(&concierto, conciertoID)

	if resultado.Error != nil {
		return domain.ListaEspera{}, errors.New("concierto no encontrado")
	}

	if concierto.CuposDisponibles > 0 {
		return domain.ListaEspera{}, errors.New("todavía hay cupos disponibles, no hace falta lista de espera")
	}

	var existente domain.ListaEspera

	resultado = dao.DB.
		Where("usuario_id = ? AND concierto_id = ? AND estado = ?", usuarioID, conciertoID, "esperando").
		First(&existente)

	if resultado.Error == nil {
		return domain.ListaEspera{}, errors.New("ya estás anotado en la lista de espera")
	}

	var cantidad int64

	dao.DB.
		Model(&domain.ListaEspera{}).
		Where("concierto_id = ? AND estado = ?", conciertoID, "esperando").
		Count(&cantidad)

	lista := domain.ListaEspera{
		UsuarioID:    usuarioID,
		ConciertoID:  conciertoID,
		PosicionCola: int(cantidad) + 1,
		Estado:       "esperando",
		FechaAlta:    time.Now(),
	}

	resultado = dao.DB.Create(&lista)

	return lista, resultado.Error
}

func ObtenerListasEsperaUsuario(usuarioID uint) ([]domain.ListaEspera, error) {

	var listas []domain.ListaEspera

	resultado := dao.DB.
		Preload("Concierto").
		Where("usuario_id = ?", usuarioID).
		Find(&listas)

	return listas, resultado.Error
}

func SalirListaEspera(usuarioID uint, listaID string) error {

	var lista domain.ListaEspera

	resultado := dao.DB.First(&lista, listaID)

	if resultado.Error != nil {
		return errors.New("registro de lista de espera no encontrado")
	}

	if lista.UsuarioID != usuarioID {
		return errors.New("no podés eliminar una lista de espera que no es tuya")
	}

	conciertoID := lista.ConciertoID
	posicionEliminada := lista.PosicionCola

	dao.DB.Delete(&lista)

	var listasPosteriores []domain.ListaEspera

	dao.DB.
		Where("concierto_id = ? AND posicion_cola > ? AND estado = ?", conciertoID, posicionEliminada, "esperando").
		Find(&listasPosteriores)

	for _, item := range listasPosteriores {
		item.PosicionCola--
		dao.DB.Save(&item)
	}

	return nil
}

func ObtenerListaEsperaConcierto(conciertoID string) ([]domain.ListaEspera, error) {
	var listas []domain.ListaEspera

	resultado := dao.DB.
		Preload("Usuario").
		Where("concierto_id = ? AND estado = ?", conciertoID, "esperando").
		Order("posicion_cola asc").
		Find(&listas)

	return listas, resultado.Error
}
