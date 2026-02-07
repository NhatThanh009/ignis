package api

import (
	"fmt"
	"ignisgo/templates"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	content, err := templates.FS.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Could not read template", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write(content)
}

func IgnisHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello IgnisGo!")
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
