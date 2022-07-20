package models

type CarGame struct {
	ID uint `json:"id"`

	DimensionX uint `json:"dimensionX"`
	DimensionY uint `json:"dimensionY"`

	CarPositionX *uint   `json:"carPositionX"`
	CarPositionY *uint   `json:"carPositionY"`
	CarDirection *string `json:"carDirection"`
}

type CarGamePlaceRequest struct {
	ID           uint   `json:"id"`
	CarPositionX uint   `json:"carPositionX"`
	CarPositionY uint   `json:"carPositionY"`
	CarDirection string `json:"carDirection"`
}

type CarGameMoveRequest struct {
	ID uint `json:"id"`
}

type CarGameTurnRequest struct {
	ID            uint   `json:"id"`
	TurnDirection string `json:"turnDirection"`
}
