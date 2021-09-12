package urls_routes

import (
	"encoding/json"
	"errors"
	"net/http"

	urls_domain "github.com/arnaugomez/url_shortener_yaml/urls/domain"
)

type jsonUrl struct {
	Url  string `json:"url"`
	Path string `json:"path"`
}

func (u jsonUrl) toUrl() urls_domain.Url {
	return urls_domain.Url{Path: u.Path, Url: u.Url}
}

func urlFromJson(r *http.Request) (url urls_domain.Url, err error) {
	var ju jsonUrl
	json.NewDecoder(r.Body).Decode(&ju)

	if ju.Url == "" {
		err = errors.New("param `url` is empty or nonexistent")
	}
	url = ju.toUrl()
	return
}
