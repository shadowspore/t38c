package t38c

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

// Tile38Client ...
type Tile38Client struct {
	addr       string
	debug      bool
	httpClient *http.Client
}

// ClientOption ...
type ClientOption func(*Tile38Client)

// WithHTTPClient ...
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Tile38Client) {
		c.httpClient = httpClient
	}
}

// Debug ...
func Debug() ClientOption {
	return func(c *Tile38Client) {
		c.debug = true
	}
}

// New ...
func New(address string, opts ...ClientOption) (*Tile38Client, error) {
	client := &Tile38Client{
		addr: address,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(client)
	}

	if err := client.Ping(); err != nil {
		return nil, err
	}

	return client, nil
}

// Execute command
func (client *Tile38Client) execute(command string, response interface{}) error {
	resp, err := client.httpClient.Post(client.addr, "", strings.NewReader(command))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if client.debug {
		log.Printf("[%s]: %s", command, b)
	}

	if !gjson.GetBytes(b, "ok").Bool() {
		return fmt.Errorf("command [%s]: %s", command, gjson.GetBytes(b, "err").String())
	}

	if response != nil {
		if err := json.Unmarshal(b, response); err != nil {
			return err
		}
	}

	return nil
}
