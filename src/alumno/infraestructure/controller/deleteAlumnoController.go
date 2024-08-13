package controller

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "go-hexagonal-v2/src/alumno/application"
)

type DeleteAlumnoController struct {
    service *application.AlumnoService
}

func NewDeleteAlumnoController(service *application.AlumnoService) *DeleteAlumnoController {
    return &DeleteAlumnoController{service: service}
}

func (ctrl *DeleteAlumnoController) Run(c *gin.Context) {
    id := c.Param("id")

    alumno, err := ctrl.service.GetAlumno(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding Alumno"})
        return
    }

    if alumno == nil { 
        c.JSON(http.StatusBadRequest, gin.H{"error": "Alumno not found"})
        return
    }

    if err := ctrl.service.DeleteAlumno(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Alumno deleted"})
}
