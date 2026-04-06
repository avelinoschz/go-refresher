package main

import (
	"errors"
	"fmt"
)

type App struct {
	Name      string
	OwnerTeam string
}

type AppStore interface {
	Save(app App) error
	GetByName(name string) (App, bool)
}

type InMemoryAppStore struct {
	apps map[string]App
}

func NewInMemoryAppStore() *InMemoryAppStore {
	return &InMemoryAppStore{
		apps: make(map[string]App),
	}
}

func (s *InMemoryAppStore) Save(app App) error {
	// TODO: implement
	return errors.New("not implemented")
}

func (s *InMemoryAppStore) GetByName(name string) (App, bool) {
	// TODO: implement
	return App{}, false
}

type AppService struct {
	store AppStore
}

func NewAppService(store AppStore) AppService {
	return AppService{store: store}
}

func (s AppService) Register(app App) error {
	// TODO: implement
	return errors.New("not implemented")
}

func main() {
	store := NewInMemoryAppStore()
	service := NewAppService(store)

	err := service.Register(App{
		Name:      "catalog",
		OwnerTeam: "platform",
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("app registered")
}
