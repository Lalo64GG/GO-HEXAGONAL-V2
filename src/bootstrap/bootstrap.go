package bootstrap

import (
    "go-hexagonal-v2/src/alumno/application"
    "go-hexagonal-v2/src/alumno/infraestructure"

    maestroApplication "go-hexagonal-v2/src/maestro/application"
    maestroInfraestructure "go-hexagonal-v2/src/maestro/infraestructure"
    "go-hexagonal-v2/src/shared/database"
    "github.com/gin-gonic/gin"
)

func Initialize() *gin.Engine {

    database.InitDB()

    alumnoRepo := infraestructure.NewMySQLAlumnoRepository(database.DB)
    alumnoService := application.NewAlumnoService(alumnoRepo)

    maestroRepo := maestroInfraestructure.NewMySQLMaestroRepository(database.DB)
    maestroService := maestroApplication.NewMaestroService(maestroRepo)



    r := gin.Default()

    infraestructure.RegisterAlumnoRoutes(r, alumnoService)
    maestroInfraestructure.RegisterMaestroRoutes(r, maestroService)

    return r
}
