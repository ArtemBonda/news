package data

import "time"

type Results struct {
	Status       string   `json:"status"`
	TotalResults int      `json:"totalResults"`
	Articles     Articles `json:"articles"`
}

type Articles []struct {
	Source      Source    `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

type Source struct {
	ID   interface{} `json:"id"`
	Name string      `json:"name"`
}
