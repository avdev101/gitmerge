package main

import (
	"eremeev/gitmerge/adapters"
	"eremeev/gitmerge/core"
	"fmt"
)

var BASE_PATH = "https://gitlab.com/api/v4"
var TOKEN = "_NA2XQAwwJDksNKxYLzE"

func main() {
	fmt.Println("start...")

	store := adapters.GitlabStore{
		Token:    TOKEN,
		BasePath: BASE_PATH,
	}

	mergeService := core.MergeService{
		MergeStore: store,
	}

	server := adapters.Server{
		MergeService: mergeService,
		Port:         9191,
	}

	server.Start()

}
