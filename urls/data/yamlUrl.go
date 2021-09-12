package urls_data

import urls_domain "github.com/arnaugomez/url_shortener_yaml/urls/domain"

type yamlUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func (u yamlUrl) toUrl() urls_domain.Url {
	return urls_domain.Url{Path: u.Path, Url: u.Url}
}

func urlToYamlUrl(u urls_domain.Url) yamlUrl {
	return yamlUrl{Path: u.Path, Url: u.Url}
}
