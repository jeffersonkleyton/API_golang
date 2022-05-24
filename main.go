package main

import (
	"fmt"
	"os"
	"testando/api"
)

func main() {
	t := os.Getenv("TESTE")
	fmt.Println(t)

	api.Run()
}
