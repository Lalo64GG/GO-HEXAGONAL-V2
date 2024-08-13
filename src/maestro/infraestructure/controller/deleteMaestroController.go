package controller 

import (
	"net/http"

    "github.com/gin-gonic/gin"
    "go-hexagonal-v2/src/maestro/application"
)

type DeleteMaestroController struct {
	service *application.MaestroService
}

func NewDeleteMaestroController(service *application.MaestroService) *DeleteMaestroController {
    return &DeleteMaestroController{service: service}
}

func (ctrl *DeleteMaestroController) Run(c *gin.Context) {
	id := c.Param("id")

	maestro, err := ctrl.service.GetMaestro(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if maestro == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Maestro not found"})
		return
	}

	if err := ctrl.service.DeleteMaestro(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message": "Maestro deleted"})
}
