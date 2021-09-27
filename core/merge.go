package core

import "fmt"

type MergeRequest struct {
	ProjectID    int `json:"project_id"`
	ID           int `json:"iid"`
	Description  string
	SourceBranch string
}

func (m MergeRequest) getIssueId() int {
	return 2
}

func (m *MergeRequest) LinkIssue() {
	issueId := m.getIssueId()
	m.Description = fmt.Sprintf("#%v\n\n%v", issueId, m.Description)
}
