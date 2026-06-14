package services

import (
	"backend/dao"
	"backend/domain"
)

type ReporteOcupacion struct {
	ID               uint
	Nombre           string
	CupoTotal        int
	CuposDisponibles int
	EntradasVendidas int
	Porcentaje       float64
}

func ObtenerReporteOcupacion() ([]ReporteOcupacion, error) {

	var conciertos []domain.Concierto

	resultado := dao.DB.Find(&conciertos)

	if resultado.Error != nil {
		return nil, resultado.Error
	}

	var reporte []ReporteOcupacion

	for _, concierto := range conciertos {

		vendidas := concierto.CupoTotal - concierto.CuposDisponibles

		var porcentaje float64

		if concierto.CupoTotal > 0 {
			porcentaje = (float64(vendidas) / float64(concierto.CupoTotal)) * 100
		}

		reporte = append(reporte, ReporteOcupacion{
			ID:               concierto.ID,
			Nombre:           concierto.Nombre,
			CupoTotal:        concierto.CupoTotal,
			CuposDisponibles: concierto.CuposDisponibles,
			EntradasVendidas: vendidas,
			Porcentaje:       porcentaje,
		})
	}

	return reporte, nil
}
