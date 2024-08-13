package controller

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "go-hexagonal-v2/src/maestro/application"
    "go-hexagonal-v2/src/shared/auth"
)

type LoginMaestroController struct {
    service *application.MaestroService
}

func NewLoginMaestroController(service *application.MaestroService) *LoginMaestroController {
    return &LoginMaestroController{service: service}
}

func (ctrl *LoginMaestroController) Login(c *gin.Context) {
    var request struct {
        Email    string `json:"email" binding:"required"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    maestro, err := ctrl.service.Authenticate(request.Email, request.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    // Genera el token JWT
    token, err := auth.GenerateJWT(maestro.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
