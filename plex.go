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

func (p *Plex) GetPlaylistInfo(playlistID string) ([]PMSPlaylistInfo, error) {
	url := fmt.Sprintf("%s/playlists/%s/items/?X-Plex-Token=%s", p.URL, playlistID, p.Token)
	resp, err := p.get(url)
	if err != nil {
		return PMSPlaylistInfo{}, err
	}
	response := &PlaylistInfo{}
	if err := xml.NewDecoder(resp.Body).Decode(respone); err != nil {
		return PMSPlaylistInfo{}, err
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

	return resp, nil
}
