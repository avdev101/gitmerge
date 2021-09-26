package adapters

import (
	"encoding/json"
	"eremeev/gitmerge/core"
	"fmt"
	"net/http"
)

type GitlabStore struct {
	BasePath string
	Token    string
}

func (s GitlabStore) GetMerge(projectId string, id int) (core.MergeRequest, error) {
	result := core.MergeRequest{}

	fmt.Printf("get merge for project: %v and id: %v\n", projectId, id)

	url := fmt.Sprintf("%v/projects/:%v/merge_requests/%v", s.BasePath, projectId, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", s.Token))

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return result, err
	}

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (s GitlabStore) SaveDescription(merge core.MergeRequest) error {
	fmt.Println("save description")
	return nil
}
