package controller

import (
	"go-hexagonal-v2/src/maestro/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetMaestroController struct {
	service *application.MaestroService
}

func NewGetMaestroController(service *application.MaestroService) *GetMaestroController {
    return &GetMaestroController{service: service}
}

func (ctrl *GetMaestroController) Run(c *gin.Context) {
	id := c.Param("id") 

	maestro, err := ctrl.service.GetMaestro(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if maestro == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Maestro not found"})
		return
	}

	c.JSON(http.StatusOK, maestro)
}
