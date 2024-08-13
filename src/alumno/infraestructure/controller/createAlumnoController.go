package controller

import (
    "net/http"
    "github.com/go-playground/validator/v10"
    "github.com/gin-gonic/gin"
    "go-hexagonal-v2/src/alumno/application"
    "go-hexagonal-v2/src/alumno/domain"
)

var validate *validator.Validate

type CreateAlumnoController struct {
    service *application.AlumnoService
}

func NewCreateAlumnoController(service *application.AlumnoService) *CreateAlumnoController {
    validate = validator.New()
    return &CreateAlumnoController{service: service}
}

func (ctrl *CreateAlumnoController) Run(c *gin.Context) {
    var alumno domain.Alumno
    if err := c.ShouldBindJSON(&alumno); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := validate.Struct(alumno); err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
        return
    }

    if err := ctrl.service.CreateAlumno(alumno); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, alumno)
}
