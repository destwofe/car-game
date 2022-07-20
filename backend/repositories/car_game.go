package repositories

import (
	"toycar_backend/domains"
	"toycar_backend/models"

	"gorm.io/gorm"
)

type carGameRepository struct {
	conn *gorm.DB
}

func NewCarGameRepository(conn *gorm.DB) domains.CarGameRepository {
	return &carGameRepository{conn}
}

func (c *carGameRepository) GetAllGames() (gs []models.CarGame, err error) {
	err = c.conn.Find(&gs).Error
	return gs, err
}

func (c *carGameRepository) NewGame(g *models.CarGame) (err error) {
	err = c.conn.Create(g).Error
	return err
}

func (c *carGameRepository) GetGame(id uint) (g *models.CarGame, err error) {
	err = c.conn.Where("id = ?", id).Find(&g).Error
	return g, err
}

func (c *carGameRepository) UpdateGame(g *models.CarGame) (err error) {
	err = c.conn.Save(&g).Error
	return err
}
