package adapters

import (
	"encoding/json"
	"eremeev/gitmerge/core"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Port         int
	MergeService core.IMergeService
}

type ObjectAttributes struct {
	SourceProjectId string `json:"source_project_id"`
	Action          string `json:"action"`
	IID             int    `json:"iid"`
}

type MergeRequestPayload struct {
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
}

func NewMergeRequestPayload(r *http.Request) (MergeRequestPayload, error) {
	var payload MergeRequestPayload

	err := json.NewDecoder(r.Body).Decode(&payload)

	return payload, err
}

func (p MergeRequestPayload) createLinkIssueCommand() core.LinkIssueCommand {
	cmd := core.LinkIssueCommand{
		ProjectID: p.ObjectAttributes.SourceProjectId,
		ID:        p.ObjectAttributes.IID,
	}
	return cmd
}

func (s Server) handleMergeRequest(w http.ResponseWriter, r *http.Request) {
	payload, err := NewMergeRequestPayload(r)

	if err != nil {
		log.Printf("[error] can't parse payload: %v", err)

		http.Error(w, fmt.Sprintf("can't parse paylod: %v", err), 500)

		return
	}

	if payload.ObjectAttributes.Action != "open" {
		return
	}

	err = s.MergeService.LinkIssue(payload.createLinkIssueCommand())

	if err != nil {
		log.Printf("[error] can't link issue: %v", err)

		http.Error(w, fmt.Sprintf("can't link issue: %v", err), 500)

		return

	}
}

func (s Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func (s Server) Start() {
	http.HandleFunc("/status", s.handleStatus)
	http.HandleFunc("/merge", s.handleMergeRequest)

	addr := fmt.Sprintf("0.0.0.0:%v", s.Port)

	log.Fatal(http.ListenAndServe(addr, nil))
}
