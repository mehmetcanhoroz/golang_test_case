package repository

import (
	"YS_TestCase/domain"
)

type GameRepository struct {
}

func (repo *GameRepository) FindAll() []domain.Game {
	return domain.GameDatabase
}
func (repo *GameRepository) FindById(id int) *domain.Game {
	for _, game := range domain.GameDatabase {
		if game.Id == id {
			return &game
		}
	}
	return nil
}
func (repo *GameRepository) Flush() []domain.Game {
	domain.GameDatabase = []domain.Game{}
	return domain.GameDatabase
}
func (repo *GameRepository) Insert(game domain.Game) domain.Game {
	domain.GameDatabase = append(domain.GameDatabase, game)
	return game
}
func (repo *GameRepository) Delete(id int) bool {
	for i, game := range domain.GameDatabase {
		if game.Id == id {
			arr1 := domain.GameDatabase[:i]
			arr2 := domain.GameDatabase[i+1:]

			domain.GameDatabase = append(arr1, arr2...)
			return true
		}
	}
	return false
}
