package urls_data

import (
	"errors"
	"fmt"

	urls_domain "github.com/arnaugomez/url_shortener_yaml/urls/domain"
)

func pathExists(u *urls_domain.Url, urls *[]urls_domain.Url) bool {
	for _, url := range *urls {
		if u.Path == url.Path {
			return true
		}
	}
	return false
}

func SaveUrl(u *urls_domain.Url) error {
	urls := getUrls()
	if u.Path != "" {
		if pathExists(u, &urls) {
			return errors.New(fmt.Sprint("Path", u.Path, "already exists"))
		}
		saveYamlUrl(*u)
		return nil
	}

	i := 1
	for {
		u.Path = fmt.Sprint(i)
		if !pathExists(u, &urls) {
			break
		}
		i = i + 1
	}

	saveYamlUrl(*u)
	return nil
}
