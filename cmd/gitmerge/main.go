package main

import (
	"eremeev/gitmerge/adapters"
	"eremeev/gitmerge/core"
	"fmt"
)

func main() {
	fmt.Println("start...")

	store := adapters.GitlabStore{}

	mergeService := core.MergeService{
		MergeStore: store,
	}

	server := adapters.Server{
		MergeService: mergeService,
		Port:         9191,
	}

	server.Start()

}
