package core

import (
	"fmt"
	"testing"
)

func TestGetIssueId(t *testing.T) {
	merge := MergeRequest{
		ProjectID:    1,
		ID:           2,
		Description:  "description",
		SourceBranch: "feature/issue/123",
	}

	issueId := merge.getIssueId()

	if issueId != "123" {
		t.Errorf("expect %v get %v", 123, issueId)
	}
}

func TestLinkIssue(t *testing.T) {
	merge := MergeRequest{
		ProjectID:    1,
		ID:           2,
		Description:  "description",
		SourceBranch: "feature/issue/123",
	}

	expected := fmt.Sprintf("#%v\n\n%v", 123, merge.Description)

	merge.LinkIssue()

	if merge.Description != expected {
		t.Errorf("desc: expected: %s, get %s", expected, merge.Description)
	}
}

func TestNoIssueIdMatch(t *testing.T) {
	merge := MergeRequest{
		ProjectID:    1,
		ID:           2,
		Description:  "description",
		SourceBranch: "feature/issue/test123",
	}

	issueId := merge.getIssueId()

	if issueId != "" {
		t.Errorf("expect %v get %v", 123, issueId)
	}
}
