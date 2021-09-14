package adapters

import "eremeev/gitmerge/core"

type GitlabStore struct {
}

func (s GitlabStore) GetMerge(projectId string, id int) (core.MergeRequest, error) {
	return core.MergeRequest{}, nil
}

func (s GitlabStore) SaveDescription(merge core.MergeRequest) error {
	return nil
}
