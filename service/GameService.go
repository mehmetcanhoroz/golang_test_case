package service

import (
	"YS_TestCase/common/dto"
	"YS_TestCase/domain"
	"YS_TestCase/repository"
	"math/rand"
	"net/http"
)

type GameService struct {
	repo repository.GameRepository
}

func (gs *GameService) FindAllGames() dto.RestResponse {

	records := gs.repo.FindAll()

	return dto.RestResponse{
		Data: records,
		Code: http.StatusOK,
	}
}

func (gs *GameService) FindById(id int) dto.RestResponse {

	record := gs.repo.FindById(id)

	if record == nil {
		return dto.NewNotFoundError("No game was found!")
	} else {
		return dto.RestResponse{
			Data: record,
			Code: http.StatusOK,
		}
	}
}

func (gs *GameService) Flush() dto.RestResponse {

	records := gs.repo.Flush()

	return dto.RestResponse{
		Data: records,
		Code: http.StatusOK,
	}

}

func (gs *GameService) AddGame(game domain.Game) dto.RestResponse {

	if game.Name == "" {
		return dto.NewBadRequest("Name cannot be empty!")
	} else if game.Price == 0 {
		return dto.NewBadRequest("Price cannot be empty!")
	}

	game.Id = rand.Int()
	record := gs.repo.Insert(game)

	return dto.RestResponse{
		Data: record,
		Code: http.StatusCreated,
	}

}


func (gs *GameService) DeleteGame(id int) dto.RestResponse {

	record := gs.FindById(id)

	if record.Code == http.StatusNotFound {
		return dto.NewNotFoundError("No game was found!")
	} else {
		deleted := gs.repo.Delete(id)

		if deleted {
			return dto.RestResponse{
				Message: "Game deleted.",
				Code: http.StatusAccepted,
			}
		}

		return dto.NewNotFoundError("Game wasn't found!")
	}
}