package main

import (
	"log"

	"github.com/cweiers1/plant-sensor-array/server"
)

func main() {
	pool, err := server.DbStart()
	if err != nil {
		log.Fatal("unable to start db %w", err)
	}
	defer pool.Close()

	server.Start(pool)
}
