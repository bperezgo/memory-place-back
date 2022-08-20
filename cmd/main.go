package main

import (
	"log"

	_ "github.com/99designs/gqlgen"
	"github.com/bperezgo/memory-place/cmd/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
