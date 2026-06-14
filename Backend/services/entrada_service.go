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

	var siguiente domain.ListaEspera

	resultadoLista := dao.DB.
		Where(
			"concierto_id = ? AND estado = ?",
			concierto.ID,
			"esperando",
		).
		Order("posicion_cola asc").
		First(&siguiente)

	if resultadoLista.Error == nil {

		nuevaEntrada := domain.Entrada{
			UsuarioID:   siguiente.UsuarioID,
			ConciertoID: concierto.ID,
			Estado:      "activa",
			FechaCompra: time.Now(),
		}

		dao.DB.Create(&nuevaEntrada)

		ahora := time.Now()

		siguiente.Estado = "asignado"
		siguiente.FechaNotificacion = &ahora

		dao.DB.Save(&siguiente)

		concierto.CuposDisponibles--

		dao.DB.Save(&concierto)

		var listasPosteriores []domain.ListaEspera

		dao.DB.
			Where(
				"concierto_id = ? AND posicion_cola > ? AND estado = ?",
				concierto.ID,
				siguiente.PosicionCola,
				"esperando",
			).
			Find(&listasPosteriores)

		for _, item := range listasPosteriores {

			item.PosicionCola--

			dao.DB.Save(&item)

		}
	}

	return nil
}

func TransferirEntrada(usuarioID uint, entradaID string, correoDestino string) error {
	var entrada domain.Entrada

	resultado := dao.DB.First(&entrada, entradaID)

	if resultado.Error != nil {
		return errors.New("entrada no encontrada")
	}

	if entrada.UsuarioID != usuarioID {
		return errors.New("no podés transferir una entrada que no es tuya")
	}

	if entrada.Estado != "activa" {
		return errors.New("solo se pueden transferir entradas activas")
	}

	var usuarioDestino domain.Usuario

	resultado = dao.DB.Where("correo = ?", correoDestino).First(&usuarioDestino)

	if resultado.Error != nil {
		return errors.New("usuario destino no encontrado")
	}

	if usuarioDestino.ID == usuarioID {
		return errors.New("no podés transferirte la entrada a vos mismo")
	}

	entrada.UsuarioID = usuarioDestino.ID

	dao.DB.Save(&entrada)

	return nil
}
