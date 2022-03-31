package engine

import "fmt"

func StartEngines(engines int) bool {
	for i := 0; i < engines; i++ {
		fmt.Println("Engine Nr ", i, "is running")
	}

	return true

}
