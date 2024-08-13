package controller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
    "go-hexagonal-v2/src/maestro/application"
)

type GetAllMaestrosController struct {
	service *application.MaestroService
}

func NewGetAllMaestroController(service *application.MaestroService) *GetAllMaestrosController {
	return &GetAllMaestrosController{service: service}
}

func (ctrl GetAllMaestrosController) Run(c *gin.Context){
	maestros, err := ctrl.service.GetAllMaestros()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, maestros)
}