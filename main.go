package main

import (
	"fmt"
	"transactionsApi/routes"
)

func main() {
	fmt.Println("[main] Ligando sistema ...")

	server := routes.Routes()

	server.Run()
}
