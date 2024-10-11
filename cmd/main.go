package main

import (
	"fmt"

	"github.com/silazemli/lab1-template/internal/server"
	"github.com/silazemli/lab1-template/internal/storage"
)

func main() {
	db, err := storage.NewDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	srv := server.NewServer(db)
	err = srv.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
}
