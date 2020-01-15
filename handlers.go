package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/harrydrippin/go-quick-link/link"
)

// HandlerGo handles the `/go` route.
func HandlerGo(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[4:] // `/go/a/b/c/` => `a/b/c`

	splitted := strings.Split(path, "/")
	if splitted[len(splitted)-1] == "" {
		splitted = splitted[0 : len(splitted)-1]
	}

	command := splitted[0]
	params := splitted[1:]

	log.Printf("Command: %s", command)
	log.Printf("Params: %v", params)

	query := link.Query{
		Command:    command,
		ParamCount: len(params),
	}

	quickLink, err := link.GetDatabase().QueryQuickLink(query)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf("Command %s with %d parameters is not found", command, len(params))))
		return
	}

	link := quickLink.Path

	for i := 0; i < quickLink.ParamCount; i++ {
		link = strings.Replace(link, fmt.Sprintf("{%d}", i), params[i], 1)
	}

	http.Redirect(w, r, link, http.StatusPermanentRedirect)
}

// HandlerGetQuickLinks handles the GET `/manage` route
func HandlerGetQuickLinks(w http.ResponseWriter, r *http.Request) {
	result := link.GetDatabase().GetQuickLinks()
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Println(err)
		return
	}
}

// HandlerSetQuickLinks handles the POST `/manage` route
func HandlerSetQuickLinks(w http.ResponseWriter, r *http.Request) {
	var quickLink link.QuickLink

	if err := json.NewDecoder(r.Body).Decode(&quickLink); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := link.GetDatabase().AddQuickLink(quickLink); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// HandlerDeleteQuickLinks handles the DELETE `/manage` route
func HandlerDeleteQuickLinks(w http.ResponseWriter, r *http.Request) {
	var query link.Query

	if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := link.GetDatabase().RemoveQuickLink(query); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

// HandlerHealthCheck handles the `/health` route.
func HandlerHealthCheck(w http.ResponseWriter, r *http.Request) {
	type HealthCheckReturn struct {
		Result    bool  `json:"result"`
		Timestamp int64 `json:"timestamp"`
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	result := &HealthCheckReturn{
		Result:    true,
		Timestamp: time.Now().Unix() * 1000,
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Println(err)
		return
	}
}
