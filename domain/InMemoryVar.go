package domain

var InMemoryDatabase map[string]string

type InMemoryVar struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
