package service

import (
	"YS_TestCase/common/dto"
	"YS_TestCase/repository"
	"net/http"
)

type InMemoryService struct {
	repo repository.InMemoryRepository
}

func (gs *InMemoryService) FindAllValues() dto.RestResponse {

	records := gs.repo.FindAll()

	return dto.RestResponse{
		Data: records,
		Code: http.StatusOK,
	}
}

func (gs *InMemoryService) FindByKey(key string) dto.RestResponse {

	record := gs.repo.FindByKey(key)

	if record.Value == "" {
		return dto.NewNotFoundError("No value was found!")
	} else {
		return dto.RestResponse{
			Data: record,
			Code: http.StatusOK,
		}
	}
}

func (gs *InMemoryService) Flush() dto.RestResponse {

	records := gs.repo.Flush()

	return dto.RestResponse{
		Data: records,
		Code: http.StatusOK,
	}

}

func (gs *InMemoryService) AddValue(key string, value string) dto.RestResponse {

	if key == "" {
		return dto.NewBadRequest("Key cannot be empty!")
	} else if value == "" {
		return dto.NewBadRequest("Value cannot be empty!")
	}

	record := gs.repo.Insert(key, value)

	return dto.RestResponse{
		Data: record,
		Code: http.StatusCreated,
	}

}

func (gs *InMemoryService) DeleteValue(key string) dto.RestResponse {

	gs.repo.Delete(key)

	return dto.RestResponse{
		Message: "Value deleted.",
		Code:    http.StatusAccepted,
	}
}
