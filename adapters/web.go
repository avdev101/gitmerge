package adapters

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func newWebhookCommand(r *http.Request) (WebHookCommand, error) {
	var cmd WebHookCommand

	err := json.NewDecoder(r.Body).Decode(&cmd)

	return cmd, err
}

func handleError(w http.ResponseWriter, msg string, err error) {
	log.Printf("[error] %v: %v", msg, err)

	http.Error(w, fmt.Sprintf("%v: %v", msg, err), 500)
}

func (s Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func (s Server) handleMergeRequest(w http.ResponseWriter, r *http.Request) {
	command, err := newWebhookCommand(r)

	if err != nil {
		handleError(w, "can't parse payload", err)
		return
	}

	err = s.WebhookCommandHandler.Handle(command)

	if err != nil {
		handleError(w, "can't link issue", err)
		return
	}
}

type Server struct {
	Port                  int
	WebhookCommandHandler IWebhookCommandHandler
}

func (s Server) Start() {
	http.HandleFunc("/status", s.handleStatus)
	http.HandleFunc("/merge", s.handleMergeRequest)

	addr := fmt.Sprintf("0.0.0.0:%v", s.Port)

	log.Fatal(http.ListenAndServe(addr, nil))
}
