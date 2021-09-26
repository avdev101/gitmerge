package adapters

import (
	"bytes"
	"encoding/json"
	"eremeev/gitmerge/core"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GitlabStore struct {
	BasePath string
	Token    string
}

func (s GitlabStore) GetMerge(projectId string, id int) (core.MergeRequest, error) {
	result := core.MergeRequest{}

	fmt.Printf("get merge for project: %v and id: %v\n", projectId, id)

	url := fmt.Sprintf("%v/projects/%v/merge_requests/%v", s.BasePath, projectId, id)
	fmt.Println(url)

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

	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

type UpdatePayload struct {
	Description string `json:"description"`
}

func (s GitlabStore) SaveDescription(merge core.MergeRequest) error {

	fmt.Println("==== save ====")

	url := fmt.Sprintf("%v/projects/%v/merge_requests/%v", s.BasePath, merge.ProjectID, merge.ID)
	fmt.Println(url)

	payload := UpdatePayload{
		Description: merge.Description,
	}

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url, payloadBuf)

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", s.Token))
	req.Header.Add("content-type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return nil
}
