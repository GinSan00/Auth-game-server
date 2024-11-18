package main

import (
	"fmt"
	"main/internal/config"
)

func main() {
	fmt.Println("Serv")

	cfg := config.MustLoad()
}
