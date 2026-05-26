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
