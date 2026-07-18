package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Bang struct {
	Prefix string
	Suffix string
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	parts := strings.Fields(query)

	var trigger string
	var searchParts []string

	for _, part := range parts {
		if strings.HasPrefix(part, "!") && len(part) > 1 && trigger == "" {
			trigger = strings.ToLower(part[1:])
			continue
		}

		searchParts = append(searchParts, part)
	}

	bang, ok := Bangs[trigger]
	if !ok {
		bang = Bang{
			Prefix: "https://www.google.com/search?q=",
		}
	}

	search := strings.Join(searchParts, " ")
	target := bang.Prefix + url.QueryEscape(search) + bang.Suffix

	// w.Header().Set("Location", target)
	// w.WriteHeader(http.StatusFound)
	http.Redirect(w, r, target, http.StatusFound)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", redirectHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8479"
	}

	log.Printf("Listening on :%s", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
