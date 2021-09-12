package urls_data

import (
	"io/ioutil"

	urls_domain "github.com/arnaugomez/url_shortener_yaml/urls/domain"
	"gopkg.in/yaml.v2"
)

func saveYamlUrl(u urls_domain.Url) error {
	yu := urlToYamlUrl(u)
	urls := append(getYamlUrls(), yu)

	ydata, err := yaml.Marshal(&urls)
	if (err != nil) {
		return err
	}
	err = ioutil.WriteFile("./assets/urls.yaml", ydata, 0)
	return err
}
