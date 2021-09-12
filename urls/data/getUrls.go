package urls_data

import urls_domain "github.com/arnaugomez/url_shortener_yaml/urls/domain"

func getUrls() []urls_domain.Url {
	yu := getYamlUrls()
	u := make([]urls_domain.Url, len(yu))
	for i, yu := range yu {
		u[i] = yu.toUrl()
	}
	return u
}