package usecases

import (
	"errors"
	"toycar_backend/domains"
	"toycar_backend/models"
	"toycar_backend/utils"
)

type carGameUsecase struct {
	carGameRepo domains.CarGameRepository
}

func NewCarGameUsecase(carGameRepo domains.CarGameRepository) domains.CarGameUsecase {
	return &carGameUsecase{carGameRepo}
}

func (c *carGameUsecase) GetAllGames() (gs []models.CarGame, err error) {
	return c.carGameRepo.GetAllGames()
}

func (c *carGameUsecase) NewGame(g *models.CarGame) (err error) {
	return c.carGameRepo.NewGame(g)
}

func (c *carGameUsecase) GetGame(id uint) (g *models.CarGame, err error) {
	return c.carGameRepo.GetGame(id)
}

func (c *carGameUsecase) Place(id uint, x uint, y uint, direction string) (g *models.CarGame, err error) {
	game, err := c.carGameRepo.GetGame(id)
	if err != nil {
		return nil, err
	}

	if game == nil {
		return nil, errors.New("game not found")
	}
	if game.CarPositionX != nil || game.CarPositionY != nil || game.CarDirection != nil {
		return nil, errors.New("already place")
	}

	game.CarPositionX = &x
	game.CarPositionY = &y
	game.CarDirection = &direction

	err = c.carGameRepo.UpdateGame(game)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (c *carGameUsecase) Move(id uint) (g *models.CarGame, err error) {
	game, err := c.carGameRepo.GetGame(id)
	if err != nil {
		return nil, err
	}

	if game == nil {
		return nil, errors.New("game not found")
	}
	if game.CarPositionX == nil || game.CarPositionY == nil || game.CarDirection == nil {
		return nil, errors.New("not place car yet")
	}

	if *game.CarDirection == "NORTH" {
		newY := *game.CarPositionY + 1
		if newY > game.DimensionY {
			return nil, errors.New("cannot move to east")
		}
		game.CarPositionY = &(newY)
	} else if *game.CarDirection == "SOUTH" {
		newY := *game.CarPositionY - 1
		if newY == 0 {
			return nil, errors.New("cannot move to south")
		}
		game.CarPositionY = &(newY)
	} else if *game.CarDirection == "EAST" {
		newX := *game.CarPositionX + 1
		if newX > game.DimensionX {
			return nil, errors.New("cannot move to east")
		}
		game.CarPositionX = &(newX)
	} else if *game.CarDirection == "WEST" {
		newX := *game.CarPositionX - 1
		if newX == 0 {
			return nil, errors.New("cannot move to west")
		}
		game.CarPositionX = &(newX)
	}

	err = c.carGameRepo.UpdateGame(game)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (c *carGameUsecase) Turn(id uint, direction string) (g *models.CarGame, err error) {
	game, err := c.carGameRepo.GetGame(id)
	if err != nil {
		return nil, err
	}

	if game == nil {
		return nil, errors.New("game not found")
	}
	if game.CarPositionX == nil || game.CarPositionY == nil || game.CarDirection == nil {
		return nil, errors.New("not place car yet")
	}

	directions := []string{"NORTH", "EAST", "SOUTH", "WEST"}

	direction_index := utils.IndexOf(*game.CarDirection, directions)

	if direction == "LEFT" {
		direction_index = utils.Mod(direction_index-1, 4)
	} else if direction == "RIGHT" {
		direction_index = utils.Mod(direction_index+1, 4)
	}

	game.CarDirection = &directions[direction_index]

	err = c.carGameRepo.UpdateGame(game)
	if err != nil {
		return nil, err
	}

	return game, nil
}
