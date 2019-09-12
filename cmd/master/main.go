package main

import (
	"time"

	"github.com/lealhugui/vice/data"
	"github.com/lealhugui/vice/routes"
)

func main() {
	data.GetData()
	routes.StartServer()
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
