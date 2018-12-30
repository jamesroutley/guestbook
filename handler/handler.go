package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jamesroutley/guestbook/datastore"
	"github.com/jamesroutley/guestbook/domain"
	"github.com/jamesroutley/guestbook/dto"
)

var ds datastore.Datastore

func Serve(port int, store datastore.Datastore) {
	ds = store
	http.HandleFunc("/log", logHandler)
	log.Printf("Listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	// Validation
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("Error reading body: %v", err)
		log.Print(msg)
		http.Error(
			w, msg,
			http.StatusInternalServerError,
		)
		return
	}
	req := &dto.LogRequest{}
	if err := json.Unmarshal(b, req); err != nil {
		msg := fmt.Sprintf("Error unmarshaling body: %v", err)
		log.Print(msg)
		http.Error(
			w, msg,
			http.StatusInternalServerError,
		)
		return
	}

	rsp, err := handleLog(r.Context(), r.RemoteAddr, req)
	if err != nil {
		msg := fmt.Sprintf("Error logging: %v", err)
		log.Print(msg)
		http.Error(
			w, msg,
			http.StatusInternalServerError,
		)
		return
	}

	rspBytes, err := json.Marshal(rsp)
	if err != nil {
		msg := fmt.Sprintf("Error marshaling response: %v", err)
		log.Print(msg)
		http.Error(
			w, msg,
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(rspBytes)
	return
}

func handleLog(ctx context.Context, remoteAddr string, req *dto.LogRequest) (*dto.LogResponse, error) {
	switch {
	case req.URL == "":
		return nil, fmt.Errorf("param 'url' missing")
	}

	created := time.Now().UTC()

	if err := ds.Store(domain.Visit{
		URL:      req.URL,
		Referrer: req.Referrer,
		IP:       remoteAddr,
		Created:  created,
	}); err != nil {
		log.Print(err)
		return nil, err
	}

	fmt.Println(created.Format(time.RFC3339), req.URL, req.Referrer, remoteAddr)

	return &dto.LogResponse{}, nil
}
