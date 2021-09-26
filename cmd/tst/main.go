package main

import (
	"eremeev/gitmerge/adapters"
	"fmt"
	"log"
)

var BASE_PATH = "https://gitlab.com/api/v4"
var TOKEN = "_NA2XQAwwJDksNKxYLzE"

func main() {
	store := adapters.GitlabStore{
		Token:    TOKEN,
		BasePath: BASE_PATH,
	}

	merge, err := store.GetMerge("15905077", 3)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(merge.Description)
}
