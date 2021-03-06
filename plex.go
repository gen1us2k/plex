package plex

import (
	"crypto/tls"
	"encoding/xml"
	"errors"
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
			Timeout: 15 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}
func (p *Plex) SetURL(url string) {
	p.URL = url
}
func (p *Plex) SetToken(token string) {
	p.Token = token
}

// GetServers returns a list of your plex servers
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

func (p *Plex) GetPlaylists() ([]PMSPlaylist, error) {
	url := fmt.Sprintf("%s/playlists/all?X-Plex-Token=%s", p.URL, p.Token)
	resp, err := p.get(url)
	if err != nil {
		return []PMSPlaylist{}, err
	}
	defer resp.Body.Close()

	response := &PlaylistResponse{}
	if err := xml.NewDecoder(resp.Body).Decode(response); err != nil {
		return []PMSPlaylist{}, err
	}
	return response.Playlists, nil
}


func (p *Plex) GetPlaylistInfo(playlistID string) (*PMSPlaylistInfo, error) {
	url := fmt.Sprintf("%s/playlists/%s/items/?X-Plex-Token=%s", p.URL, playlistID, p.Token)
	resp, err := p.get(url)
	if err != nil {
		return &PMSPlaylistInfo{}, err
	}
	response := &PMSPlaylistInfo{}
	if err := xml.NewDecoder(resp.Body).Decode(response); err != nil {
		return &PMSPlaylistInfo{}, err
	}
	return response, err
}

func (p *Plex) get(url string) (*http.Response, error) {
	client := p.HTTPClient

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return &http.Response{}, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return &http.Response{}, err
	}
	if resp.StatusCode == 401 {
		return &http.Response{}, errors.New("Unauthorized")
	}

	return resp, nil
}

func defaultHeaders() headers {
	return headers{
		Platform:        "Linux",
		PlatformVersion: "0.0.0",
		Product:         "Go Plex Client",
		Version:         "0.0.1",
		Device:          "Ololo ",
		ContainerSize:   "Plex-Container-Size=50",
		ContainerStart:  "X-Plex-Container-Start=0",
		Accept:          "application/json",
		ContentType:     "application/json",
	}
}
