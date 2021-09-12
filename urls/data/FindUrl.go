package urls_data

import (
	"errors"
	"fmt"

	urls_domain "github.com/arnaugomez/url_shortener_yaml/urls/domain"
)

func FindUrl(p string) (url urls_domain.Url, err error) {
	u := getUrls()
	for _, u := range u {
		if u.Path == p {
			return u, nil
		}
	}
	return urls_domain.Url{}, errors.New(fmt.Sprint("Website with path ", p, " not found"))
}
