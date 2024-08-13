package controller

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "go-hexagonal-v2/src/alumno/application"
)

type GetAllAlumnosController struct {
    service *application.AlumnoService
}

func NewGetAllAlumnosController(service *application.AlumnoService) *GetAllAlumnosController {
    return &GetAllAlumnosController{service: service}
}

func (ctrl *GetAllAlumnosController) Run(c *gin.Context) {
    alumnos, err := ctrl.service.GetAllAlumnos()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, alumnos)
}
