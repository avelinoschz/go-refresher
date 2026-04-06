package main

import "testing"

func TestInMemoryAppStoreSaveAndGetByName(t *testing.T) {
	t.Parallel()

	store := NewInMemoryAppStore()
	app := App{
		Name:      "catalog",
		OwnerTeam: "platform",
	}

	if err := store.Save(app); err != nil {
		t.Fatalf("save app: %v", err)
	}

	got, found := store.GetByName("catalog")
	if !found {
		t.Fatal("expected app to be found")
	}

	if got != app {
		t.Fatalf("unexpected app: got %+v want %+v", got, app)
	}
}

func TestAppServiceRegisterValidatesRequiredFields(t *testing.T) {
	t.Parallel()

	service := NewAppService(NewInMemoryAppStore())

	testCases := []App{
		{Name: "", OwnerTeam: "platform"},
		{Name: "catalog", OwnerTeam: ""},
	}

	for _, app := range testCases {
		app := app
		t.Run(app.Name+"-"+app.OwnerTeam, func(t *testing.T) {
			t.Parallel()

			if err := service.Register(app); err == nil {
				t.Fatal("expected validation error")
			}
		})
	}
}

func TestAppServiceRegisterRejectsDuplicateNames(t *testing.T) {
	t.Parallel()

	store := NewInMemoryAppStore()
	service := NewAppService(store)

	app := App{
		Name:      "catalog",
		OwnerTeam: "platform",
	}

	if err := service.Register(app); err != nil {
		t.Fatalf("register initial app: %v", err)
	}

	if err := service.Register(app); err == nil {
		t.Fatal("expected duplicate-name error")
	}
}

func TestAppServiceRegisterDelegatesPersistenceToStore(t *testing.T) {
	t.Parallel()

	store := NewInMemoryAppStore()
	service := NewAppService(store)

	app := App{
		Name:      "catalog",
		OwnerTeam: "platform",
	}

	if err := service.Register(app); err != nil {
		t.Fatalf("register app: %v", err)
	}

	got, found := store.GetByName("catalog")
	if !found {
		t.Fatal("expected app to be saved in store")
	}

	if got != app {
		t.Fatalf("unexpected saved app: got %+v want %+v", got, app)
	}
}
