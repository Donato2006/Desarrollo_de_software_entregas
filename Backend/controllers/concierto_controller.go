package controllers

import (
	"backend/domain"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ObtenerConciertos maneja la petición GET para listar todos los conciertos.
func ObtenerConciertos(c *gin.Context) {
	// Llama al servicio para obtener el listado de la base de datos
	conciertos, err := services.ObtenerTodosLosConciertos()
	// Si ocurre un error en la base de datos, responde con código 500
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener conciertos",
		})
		return
	}
	// Si todo sale bien, retorna la lista de conciertos con un código 200 OK
	c.JSON(http.StatusOK, conciertos)
}

// CrearConcierto maneja la petición POST para registrar un nuevo concierto.
func CrearConcierto(c *gin.Context) {
	var concierto domain.Concierto
	// Intenta mapear el cuerpo JSON de la petición a la estructura 'Concierto'
	if err := c.ShouldBindJSON(&concierto); err != nil {
		// Si el JSON está mal formado o faltan datos requeridos, devuelve 400 Bad Request
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos inválidos",
		})
		return
	}
	// Llama al servicio para insertar el nuevo concierto en la base de datos
	conciertoCreado, err := services.CrearConcierto(concierto)
	// Si falla la inserción en la base de datos, responde con código 500
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al crear concierto",
		})
		return
	}
	// Retorna el objeto concierto recién creado con el código 201 Created
	c.JSON(http.StatusCreated, conciertoCreado)
}

// ObtenerConciertoPorID maneja la petición GET para buscar un concierto específico usando su ID.
func ObtenerConciertoPorID(c *gin.Context) {
	// Extrae el parámetro 'id' de la URL (ej: /conciertos/5)
	id := c.Param("id")
	// Llama al servicio para buscar el registro por su ID
	concierto, err := services.ObtenerConciertoPorID(id)
	// Si no se encuentra el concierto o hay un error, responde con 404 Not Found
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Concierto no encontrado",
		})
		return
	}
	// Si se encuentra, lo devuelve con un código 200 OK
	c.JSON(http.StatusOK, concierto)
}

// ActualizarConcierto maneja la petición PUT para modificar los datos de un concierto existente.
func ActualizarConcierto(c *gin.Context) {
	// Obtiene el ID del concierto desde los parámetros de la URL
	id := c.Param("id")

	var datosActualizados domain.Concierto
	// Valida que el JSON recibido en el cuerpo coincida con la estructura esperada
	if err := c.ShouldBindJSON(&datosActualizados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos inválidos",
		})
		return
	}
	// Llama al servicio para buscar y actualizar el concierto con los nuevos datos
	conciertoActualizado, err := services.ActualizarConcierto(id, datosActualizados)
	// Si el concierto no existe en la BD, devuelve un error 404
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Concierto no encontrado",
		})
		return
	}
	// Retorna el concierto ya modificado con un código 200 OK
	c.JSON(http.StatusOK, conciertoActualizado)
}

// EliminarConcierto maneja la petición DELETE para remover un concierto del sistema.
func EliminarConcierto(c *gin.Context) {
	// Obtiene el ID del concierto a eliminar desde la URL
	id := c.Param("id")
	// Llama al servicio encargado de la eliminación
	err := services.EliminarConcierto(id)
	// Si el registro no existe o no se puede borrar, devuelve 404
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Concierto no encontrado",
		})
		return
	}
	// Confirma el éxito de la operación con un mensaje JSON y código 200 OK
	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Concierto eliminado correctamente",
	})
}
