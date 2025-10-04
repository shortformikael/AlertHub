package main

import (
	"fmt"

	"github.com/shortformikael/AlertHub/src/config"
)

func main() {
	fmt.Println("Starting AlertHub...")
	var c = config.NewConfig()
	fmt.Println(c)
	fmt.Println("AlertHub configuration loaded successfully!")
}
