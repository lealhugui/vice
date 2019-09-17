package main

import (
	"time"

	"github.com/lealhugui/vice/data"

	"github.com/lealhugui/vice/cmd/master/routes"
)

func main() {

	data.GlobalConn.DB.AutoMigrate(&data.Task{})
	data.GlobalConn.DB.AutoMigrate(&data.Worker{})

	routes.StartServer()
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
