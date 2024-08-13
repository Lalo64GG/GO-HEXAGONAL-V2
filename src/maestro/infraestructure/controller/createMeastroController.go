package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go-hexagonal-v2/src/maestro/application"
    "go-hexagonal-v2/src/maestro/domain"
    "net/http"
)

var validate *validator.Validate

type CreateMaestroController struct {
    service *application.MaestroService
}

func NewCreateMaestroContoller(service *application.MaestroService) *CreateMaestroController {
    validate = validator.New()
    return &CreateMaestroController{service: service}
}

func (ctrl *CreateMaestroController) Run(c *gin.Context) {
    var maestro domain.Maestro

    if err := c.ShouldBindJSON(&maestro); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := validate.Struct(maestro); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
        return
    }

    if err := ctrl.service.CreateMaestro(maestro); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "id":    maestro.ID, 
        "name":  maestro.Name,
        "email": maestro.Email,
    })
}
