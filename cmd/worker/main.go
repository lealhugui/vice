package main

import (
	"fmt"

	"github.com/lealhugui/vice/cmd/worker/config"
)

func main() {
	fmt.Print("worker")
	err := config.LoadCfg()
	if err != nil {
		panic(err)
	}

	RegisterOnInit()

}
