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
	s.apps[app.Name] = app
	return nil
}

func (s *InMemoryAppStore) GetByName(name string) (App, bool) {
	app, found := s.apps[name]
	return app, found
}

type AppService struct {
	store AppStore
}

func NewAppService(store AppStore) AppService {
	return AppService{store: store}
}

func (s AppService) Register(app App) error {
	if app.Name == "" {
		return errors.New("name is required")
	}

	if app.OwnerTeam == "" {
		return errors.New("owner team is required")
	}

	if _, found := s.store.GetByName(app.Name); found {
		return fmt.Errorf("app %q already exists", app.Name)
	}

	return s.store.Save(app)
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
