package adapters

import (
	"eremeev/gitmerge/core"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Port         int
	MergeService core.IMergeService
}

func createLinkCommand(r *http.Request) (core.LinkIssueCommand, error) {
	cmd := core.LinkIssueCommand{
		ProjectID: "1",
		ID:        1,
	}

	return cmd, nil

}

func (s Server) handleMergeRequest(w http.ResponseWriter, r *http.Request) {
	cmd, err := createLinkCommand(r)

	if err != nil {
		log.Printf("[error] can't create command: %v", err)

		http.Error(w, fmt.Sprintf("can't create command: %v", err), 500)

		return
	}

	err = s.MergeService.LinkIssue(cmd)

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
