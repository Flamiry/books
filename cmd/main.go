package main

import (
	"log"

	"github.com/Flamiry/books.git/internal/config"
	"github.com/Flamiry/books.git/internal/server"
	"github.com/Flamiry/books.git/internal/storage"
)

 func main() {
	cfg := config.Readconfig()
	log.Println(cfg)
	storage := storage.New()
	server := server.New(cfg.Host, storage)

	if err := server.Run(); err != nil {
		panic(err)
	}
 }