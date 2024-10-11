package main

import (
	"fmt"

	"github.com/JWindy92/go-playground/internal/database"
	"github.com/JWindy92/go-playground/internal/handlers"
	"github.com/JWindy92/go-playground/internal/models"
)

func main() {
	fmt.Println("Running interfaces-example main.go")
	fmt.Println("This is how the code will be run in the REAL world")
	handler_imp := handlers.NewHandlerImplementation(
		models.NewModelsImplementation(
			database.NewDB("some dummy database"),
		),
	)
	handler_imp.Get("something")
}
