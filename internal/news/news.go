package news

import (
	"encoding/json"
	"fmt"
	"github.com/ArtemBonda/news/internal/data"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	client   *http.Client
	key      string
	pageSize int
}

func NewClient(client *http.Client, key string, pageSize int) *Client {
	if pageSize > 100 {
		pageSize = 100
	}
	return &Client{client: client, key: key, pageSize: pageSize}
}

func (c *Client) FetchEverything(query, page string) (*data.Results, error) {
	request := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&pageSize=%d&page=%s&sortBy=publishedAt&apiKey=%s&language=ru",
		url.QueryEscape(query), c.pageSize, page, c.key)
	resp, err := c.client.Get(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	result := &data.Results{}
	err = json.Unmarshal(body, result)
	return result, err
}
