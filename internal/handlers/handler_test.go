package handlers

import (
	"fmt"
	"testing"

	"github.com/JWindy92/go-playground/internal/models"
	"github.com/stretchr/testify/mock"
)

type mock_models struct{}
type mock_database struct{}

func (m *mock_models) GetSomethingFromDb() {
	fmt.Println("Mock Models speaking. This is just a test for handlers, I want to make sure models.GetSomethingFromDb is called")
	fmt.Println("But since we are just testing, there's no need to actually talk to a database")
}

func (m *mock_database) DoDbStuff() {
	fmt.Println("Hello, I'm just a mock database. I just answer when models asks me to, but I'm not going to do anything.")
}

func TestGet(t *testing.T) {
	fmt.Println("I'm testing my handler Get() logic, but I don't actually need models to do anything")
	handler_imp := NewHandlerImplementation(&mock_models{})
	handler_imp.Get("test request")
}

func TestAnother(t *testing.T) {
	fmt.Println("I'm testing my handler Get() again, but this time I'd like the models to function normally.")
	fmt.Println("However, I don't need the database to do anything.")
	handler_imp := NewHandlerImplementation(models.NewModelsImplementation(&mock_database{}))
	handler_imp.Get("test request")
}

type MockedModels struct {
	mock.Mock
}

// func (m *MockedModels) GetSomethingFromDb() {
// 	fmt.Println("Called Mocked GetSomethingFromDb")
// 	m.Called()
// }

// func TestingWithMockingLibrary(t *testing.T) {
// 	mockedModels := new(MockedModels)
// 	handler_imp := NewHandlerImplementation(&mockedModels)
// }
