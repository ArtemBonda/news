package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"

	"github.com/ArtemBonda/news/internal/news"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func SearchMiddleware(newsapi *news.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rawURL, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		params := rawURL.Query()
		searchQuery := params.Get("q")
		page := params.Get("page")
		if page == "" {
			page = "1"
		}

		results, err := newsapi.FetchEverything(searchQuery, page)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		fmt.Printf("%+v", results)

	}
}
