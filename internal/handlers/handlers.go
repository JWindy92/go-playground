package handlers

import "fmt"

type ModelsInterface interface {
	GetSomethingFromDb()
}

type HandlerImplementation struct {
	models ModelsInterface
}

func NewHandlerImplementation(models ModelsInterface) *HandlerImplementation {
	return &HandlerImplementation{
		models: models,
	}
}

func (h *HandlerImplementation) Get(req string) {
	fmt.Println("Get request recievied, I'll handle it...")
	h.askModelsForSomething(req)
}

func (h *HandlerImplementation) askModelsForSomething(target string) {
	fmt.Printf("Models knows a guy in the database department that can get %s for us. I'll ask them...\n", target)
	h.models.GetSomethingFromDb()
}
