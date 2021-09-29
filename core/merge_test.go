package core

import "testing"

func TestGetIssueId(t *testing.T) {
	merge := MergeRequest{
		ProjectID:    1,
		ID:           2,
		Description:  "description",
		SourceBranch: "feature/issue/123",
	}

	issueId := merge.getIssueId()

	if issueId != 123 {
		t.Errorf("expect %v get %v", 123, issueId)
	}
}
