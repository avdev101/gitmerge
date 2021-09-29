package core

import (
	"fmt"
	"regexp"
)

type MergeRequest struct {
	ProjectID    int `json:"project_id"`
	ID           int `json:"iid"`
	Description  string
	SourceBranch string
}

func (m MergeRequest) getIssueId() string {
	re := regexp.MustCompile("issue/([0-9]+)")
	match := re.FindStringSubmatch(m.SourceBranch)

	if len(match) == 2 {
		return match[1]
	}

	return ""
}

func (m *MergeRequest) LinkIssue() {
	issueId := m.getIssueId()
	m.Description = fmt.Sprintf("#%v\n\n%v", issueId, m.Description)
}
