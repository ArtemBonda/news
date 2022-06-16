package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
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

	fmt.Println("Search Query is: ", searchQuery)
	fmt.Println("Page is: ", page)
}
