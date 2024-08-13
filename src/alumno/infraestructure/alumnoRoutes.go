package infraestructure

import (
    "github.com/gin-gonic/gin"
    "go-hexagonal-v2/src/alumno/application"
    "go-hexagonal-v2/src/alumno/infraestructure/controller"
	"go-hexagonal-v2/src/shared/middleware"
)

func RegisterAlumnoRoutes(r *gin.Engine, service *application.AlumnoService) {
    createCtrl := controller.NewCreateAlumnoController(service)
    getCtrl := controller.NewGetAlumnoController(service)
    getAllCtrl := controller.NewGetAllAlumnosController(service)
    deleteCtrl := controller.NewDeleteAlumnoController(service)

    r.POST("/alumnos", middleware.JWTAuthMiddleware(), createCtrl.Run)
    r.GET("/alumnos/:id", middleware.JWTAuthMiddleware(), getCtrl.Run)
    r.GET("/alumnos",middleware.JWTAuthMiddleware(), getAllCtrl.Run) 
    r.DELETE("/alumnos/:id", middleware.JWTAuthMiddleware(), deleteCtrl.Run)
}
