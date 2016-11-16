package plex

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

const plexURL = "http://plex.tv"

type (
	Plex struct {
		HTTPClient *http.Client
		URL        string
		Token      string
	}
)

func New(token string) *Plex {
	return &Plex{
		Token: token,
		HTTPClient: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
}
func (p *Plex) SetURL(url string) {
	p.URL = url
}
func (p *Plex) SetToken(token string) {
	p.Token = token
}
func (p *Plex) GetServers() ([]PMSServer, error) {
	url := fmt.Sprintf("%s/pms/servers.xml?X-Plex-Token=%s", plexURL, p.Token)
	resp, err := p.get(url)
	if err != nil {
		return []PMSServer{}, err
	}
	defer resp.Body.Close()
	response := &ServersResponse{}
	if err := xml.NewDecoder(resp.Body).Decode(response); err != nil {
		return []PMSServer{}, err
	}
	return response.Servers, nil
}

func (p *Plex) get(url string) (*http.Response, error) {
	client := p.HTTPClient

	req, err := http.NewRequest("GET", query, nil)

	if reqErr != nil {
		return &http.Response{}, reqErr
	}

	resp, err := client.Do(req)

	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}
