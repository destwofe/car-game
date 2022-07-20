package usecases_test

import (
	"testing"
	"toycar_backend/models"
	"toycar_backend/repositories/mocks"
	"toycar_backend/usecases"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mockRequest := &models.CarGame{
		DimensionX: 5,
		DimensionY: 5,
	}

	t.Run("success", func(t *testing.T) {
		mockCarGameRepo := new(mocks.CarGameRepository)

		mockCarGameRepo.On("NewGame", mockRequest).Return(nil).Once()

		u := usecases.NewCarGameUsecase(mockCarGameRepo)

		err := u.NewGame(mockRequest)

		assert.NoError(t, err)

		mockCarGameRepo.AssertExpectations(t)
	})
}
