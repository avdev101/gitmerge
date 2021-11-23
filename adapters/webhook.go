package adapters

import (
	"eremeev/gitmerge/core"
	"fmt"
)

type ObjectAttributes struct {
	SourceProjectId int    `json:"source_project_id"`
	Action          string `json:"action"`
	IID             int    `json:"iid"`
}

type WebHookCommand struct {
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
}

func (c WebHookCommand) isNewMergeRequest() bool {
	return c.ObjectAttributes.Action == "" || c.ObjectAttributes.Action == "open"
}

type IWebhookCommandHandler interface {
	Handle(c WebHookCommand) error
}

type WebhookCommandHandler struct {
	mergeService core.MergeService
}

func (h WebhookCommandHandler) Handle(c WebHookCommand) error {
	if c.isNewMergeRequest() {
		return h.handleNewMergeRequest(c)
	}

	return nil
}

func (h WebhookCommandHandler) createLinkIssueCommand(c WebHookCommand) core.LinkIssueCommand {
	cmd := core.LinkIssueCommand{
		ProjectID: fmt.Sprint(c.ObjectAttributes.SourceProjectId),
		ID:        c.ObjectAttributes.IID,
	}
	return cmd
}

func (h WebhookCommandHandler) handleNewMergeRequest(c WebHookCommand) error {
	return h.mergeService.LinkIssue(h.createLinkIssueCommand(c))
}
