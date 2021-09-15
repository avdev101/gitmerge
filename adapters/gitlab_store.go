package adapters

import (
	"eremeev/gitmerge/core"
	"fmt"
)

type GitlabStore struct {
}

func (s GitlabStore) GetMerge(projectId string, id int) (core.MergeRequest, error) {
	fmt.Printf("get merge for project: %v and id: %v\n", projectId, id)
	return core.MergeRequest{}, nil
}

func (s GitlabStore) SaveDescription(merge core.MergeRequest) error {
	fmt.Println("save description")
	return nil
}
