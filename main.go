package main

import (
	"tugas13/database"
	"tugas13/routers"

	_ "github.com/lib/pq"
)

func main() {
	database.Initiator()
	routers.StartServer()
}
