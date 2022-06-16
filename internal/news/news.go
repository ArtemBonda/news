package news

import "net/http"

type Client struct {
	client *http.Client
	key    string
	page   int
}

func NewClient(client *http.Client, key string, page int) *Client {
	if page > 100 {
		page = 100
	}
	return &Client{client: client, key: key, page: page}
}
