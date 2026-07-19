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

	prefix := os.Getenv("BANG_PREFIX")
	if prefix == "" {
		prefix = "!"
	}

	for _, part := range parts {
		if strings.HasPrefix(part, prefix) && len(part) > len(prefix) && trigger == "" {
			trigger = strings.ToLower(part[len(prefix):])
			continue
		}

		searchParts = append(searchParts, part)
	}

	bang, ok := Bangs[trigger]
	if !ok {
		defaultSearchPrefix := os.Getenv("DEFAULT_SEARCH_PREFIX")
		defaultSearchSuffix := os.Getenv("DEFAULT_SEARCH_SUFFIX")

		bang = Bang{
			Prefix: defaultSearchPrefix,
			Suffix: defaultSearchSuffix,
		}

		if defaultSearchPrefix == "" {
			bang = Bang{
				Prefix: "https://www.google.com/search?q=",
			}
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
