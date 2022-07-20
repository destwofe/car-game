package handlers

import (
	"net/http"
	"strconv"
	"toycar_backend/domains"
	"toycar_backend/models"

	"github.com/gin-gonic/gin"
)

type CarGameHandler struct {
	carGameUsecase domains.CarGameUsecase
}

func NewCarGameHandler(carGameUsecase domains.CarGameUsecase) *CarGameHandler {
	return &CarGameHandler{carGameUsecase}
}

func (g *CarGameHandler) GetAllGames(c *gin.Context) {
	games, err := g.carGameUsecase.GetAllGames()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, games)
}

func (g *CarGameHandler) NewGame(c *gin.Context) {
	var game *models.CarGame
	if err := c.ShouldBind(&game); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := g.carGameUsecase.NewGame(game)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, game)
}

func (g *CarGameHandler) GetGame(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	game, err := g.carGameUsecase.GetGame(uint(id))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, game)
}

func (g *CarGameHandler) Place(c *gin.Context) {
	var req *models.CarGamePlaceRequest
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	game, err := g.carGameUsecase.Place(req.ID, req.CarPositionX, req.CarPositionY, req.CarDirection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, game)
}

func (g *CarGameHandler) Move(c *gin.Context) {
	var req *models.CarGameMoveRequest
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	game, err := g.carGameUsecase.Move(req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, game)
}

func (g *CarGameHandler) Turn(c *gin.Context) {
	var req *models.CarGameTurnRequest
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	game, err := g.carGameUsecase.Turn(req.ID, req.TurnDirection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, game)
}
