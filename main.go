package main

import (
	"log"
	"os"
)

func main() {
	r := newRouter()

	port := os.Getenv("PORT")
	log.Fatal(r.Run(":" + port))
}
