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

func (s Server) HandleMergeRequest(w http.ResponseWriter, r *http.Request) {
	cmd := core.LinkIssueCommand{
		ProjectID: "1",
		ID:        1,
	}
	s.MergeService.LinkIssue(cmd)
}

func (s Server) Start() {
	http.HandleFunc("/merge", s.HandleMergeRequest)

	addr := fmt.Sprintf("0.0.0.0:%v", s.Port)

	log.Fatal(http.ListenAndServe(addr, nil))
}
