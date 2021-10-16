package handler

import (
	"YS_TestCase/common/dto"
	"YS_TestCase/domain"
	"YS_TestCase/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type GameHandler struct {
	gameService service.GameService
}

func (handler *GameHandler) GetAllGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response dto.RestResponse

	id := r.URL.Query().Get("id")
	if id != "" {
		if id, err := strconv.Atoi(id); err != nil {
			badRequest := dto.NewBadRequest("ID should be numeric!")
			w.WriteHeader(badRequest.Code)
			json.NewEncoder(w).Encode(badRequest)
		} else {
			response = handler.gameService.FindById(id)
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
		}
	} else {
		response = handler.gameService.FindAllGames()
		w.WriteHeader(response.Code)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	}
}

func (handler *GameHandler) FlushGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := handler.gameService.Flush()

	w.WriteHeader(response.Code)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func (handler *GameHandler) AddGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var game domain.Game
	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		panic(err)
	} else {
		response := handler.gameService.AddGame(game)
		json.NewEncoder(w).Encode(response)
	}
	return
}

func (handler *GameHandler) DeleteGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqId := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(reqId, 10, 32)

	if  err != nil {
		badRequest := dto.NewBadRequest("ID should be numeric!")
		w.WriteHeader(badRequest.Code)
		json.NewEncoder(w).Encode(badRequest)
	} else {
		response := handler.gameService.DeleteGame(int(id))
		json.NewEncoder(w).Encode(response)
	}
}
