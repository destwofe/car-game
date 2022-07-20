package domains

import "toycar_backend/models"

type CarGameUsecase interface {
	GetAllGames() (gs []models.CarGame, err error)
	NewGame(g *models.CarGame) (err error)
	GetGame(id uint) (g *models.CarGame, err error)
	Place(id uint, x uint, y uint, direction string) (g *models.CarGame, err error)
	Move(id uint) (g *models.CarGame, err error)
	Turn(id uint, direction string) (g *models.CarGame, err error)
}

type CarGameRepository interface {
	GetAllGames() (gs []models.CarGame, err error)
	NewGame(g *models.CarGame) (err error)
	GetGame(id uint) (g *models.CarGame, err error)
	UpdateGame(g *models.CarGame) (err error)
}
