package main

import (
	"fmt"

	"github.com/joelgarciajr84/go-project-structure-template/internal/engine"
)

func main() {

	fmt.Println("Hello, starting engines")
	totalEngines := 40
	engine.StartEngines(totalEngines)

}
