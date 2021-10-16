package repository

import (
	"YS_TestCase/domain"
)

type InMemoryRepository struct {
}

func (repo *InMemoryRepository) FindAll() map[string]string {
	return domain.InMemoryDatabase
}
func (repo *InMemoryRepository) FindByKey(key string) domain.InMemoryVar {
	return domain.InMemoryVar{
		Key: key,
		Value: domain.InMemoryDatabase["key"],
	}
}
func (repo *InMemoryRepository) Flush() map[string]string {
	domain.InMemoryDatabase = make(map[string]string)
	return domain.InMemoryDatabase
}
func (repo *InMemoryRepository) Insert(key string, value string) domain.InMemoryVar {
	domain.InMemoryDatabase[key] = value
	return domain.InMemoryVar{Key: key, Value: value}
}
func (repo *InMemoryRepository) Delete(key string) {
	delete(domain.InMemoryDatabase, key)
}
