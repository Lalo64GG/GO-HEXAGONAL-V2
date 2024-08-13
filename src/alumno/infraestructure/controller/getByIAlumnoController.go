package controller

import (
    "net/http"


    "github.com/gin-gonic/gin"
    "go-hexagonal-v2/src/alumno/application"
)

type GetAlumnoController struct {
    service *application.AlumnoService
}

func NewGetAlumnoController(service *application.AlumnoService) *GetAlumnoController {
    return &GetAlumnoController{service: service}
}

func (ctrl *GetAlumnoController) Run(c *gin.Context) {
    id := c.Param("id")
 

    alumno, err := ctrl.service.GetAlumno(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    if alumno == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Alumno not found"})
        return
    }

    c.JSON(http.StatusOK, alumno)
}
