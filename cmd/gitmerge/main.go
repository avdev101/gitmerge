package main

import (
	"eremeev/gitmerge/adapters"
	"eremeev/gitmerge/core"
	"flag"
	"fmt"
)

var basePath = flag.String("base_path", "https://gitlab.com/api/v4", "api base path")
var token = flag.String("token", "_NA2XQAwwJDksNKxYLzE", "api token")
var port = flag.Int("port", 9191, "listen port")

func main() {
	fmt.Println("start...")

	flag.Parse()

	store := adapters.GitlabStore{
		Token:    *token,
		BasePath: *basePath,
	}

	mergeService := core.MergeService{
		MergeStore: store,
	}

	webhookCommandHandler := adapters.NewWebhookCommandHandler(mergeService)

	server := adapters.Server{
		WebhookCommandHandler: webhookCommandHandler,
		Port:                  *port,
	}

	server.Start()

}
