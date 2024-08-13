package infraestructure

import (
	"github.com/gin-gonic/gin"
	"go-hexagonal-v2/src/maestro/application"
    "go-hexagonal-v2/src/maestro/infraestructure/controller"
)

func RegisterMaestroRoutes(r *gin.Engine, service *application.MaestroService){
	createCtrl := controller.NewCreateMaestroContoller(service)
	getCtrl := controller.NewGetMaestroController(service)
	getAll := controller.NewGetAllMaestroController(service)
	delete := controller.NewDeleteMaestroController(service)
	login := controller.NewLoginMaestroController(service)

	r.POST("/maestro", createCtrl.Run)
	r.POST("/maestro/login", login.Login)
	r.GET("/maestro/:id", getCtrl.Run)
	r.GET("/maestros", getAll.Run)
	r.DELETE("/maestro/:id", delete.Run)


}