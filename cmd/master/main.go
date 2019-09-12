package main

import (
	"time"

	"github.com/lealhugui/vice/routes"
)

func main() {
	routes.StartServer()
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
