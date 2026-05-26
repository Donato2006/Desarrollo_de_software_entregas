package services

import (
	"backend/dao"
	"backend/domain"
	"errors"
	"time"
)

func ComprarEntrada(usuarioID uint, conciertoID uint) (domain.Entrada, error) {

	var concierto domain.Concierto

	resultado := dao.DB.First(&concierto, conciertoID)

	if resultado.Error != nil {
		return domain.Entrada{}, errors.New("concierto no encontrado")
	}

	if concierto.CuposDisponibles <= 0 {
		return domain.Entrada{}, errors.New("no hay cupos disponibles")
	}

	entrada := domain.Entrada{
		UsuarioID:   usuarioID,
		ConciertoID: conciertoID,
		Estado:      "activa",
		FechaCompra: time.Now(),
	}

	resultado = dao.DB.Create(&entrada)

	if resultado.Error != nil {
		return domain.Entrada{}, resultado.Error
	}

	concierto.CuposDisponibles--

	dao.DB.Save(&concierto)

	return entrada, nil
}

func ObtenerEntradasUsuario(usuarioID uint) ([]domain.Entrada, error) {

	var entradas []domain.Entrada

	resultado := dao.DB.
		Preload("Concierto").
		Where("usuario_id = ?", usuarioID).
		Find(&entradas)

	return entradas, resultado.Error
}

func CancelarEntrada(usuarioID uint, entradaID string) error {
	var entrada domain.Entrada

	resultado := dao.DB.First(&entrada, entradaID)

	if resultado.Error != nil {
		return errors.New("entrada no encontrada")
	}

	if entrada.UsuarioID != usuarioID {
		return errors.New("no podés cancelar una entrada que no es tuya")
	}

	if entrada.Estado == "cancelada" {
		return errors.New("la entrada ya está cancelada")
	}

	entrada.Estado = "cancelada"
	dao.DB.Save(&entrada)

	var concierto domain.Concierto
	dao.DB.First(&concierto, entrada.ConciertoID)

	concierto.CuposDisponibles++
	dao.DB.Save(&concierto)

	return nil
}
