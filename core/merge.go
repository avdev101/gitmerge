package core

import "fmt"

type MergeRequest struct {
	ProjectID    string
	ID           int
	Description  string
	SourceBranch string
}

func (m MergeRequest) getIssueId() int {
	return 123
}

func (m MergeRequest) LinkIssue() {
	issueId := m.getIssueId()
	m.Description = fmt.Sprintf("#%v\n\n%v", issueId, m.Description)
}
