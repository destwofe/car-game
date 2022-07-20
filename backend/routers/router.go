package routers

import (
	"toycar_backend/database"
	"toycar_backend/handlers"
	"toycar_backend/repositories"
	"toycar_backend/usecases"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	carGameRepository := repositories.NewCarGameRepository(database.DB)
	carGameUsecase := usecases.NewCarGameUsecase(carGameRepository)
	carGameHandler := handlers.NewCarGameHandler(carGameUsecase)

	r := gin.Default()

	// cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	v1 := r.Group("/v1")

	game := v1.Group("/game")
	game.GET("/", carGameHandler.GetAllGames)
	game.POST("/", carGameHandler.NewGame)
	game.GET("/:id", carGameHandler.GetGame)
	game.POST("/place", carGameHandler.Place)
	game.POST("/move", carGameHandler.Move)
	game.POST("/turn", carGameHandler.Turn)

	return r
}
