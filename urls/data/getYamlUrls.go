package urls_data

import (
	"os"

	"gopkg.in/yaml.v2"
)

func getYamlUrls() []yamlUrl {
	text, err := os.ReadFile("./assets/urls.yaml")
	if err != nil {
		panic(err)
	}
	var urls []yamlUrl
	err2 := yaml.Unmarshal(text, &urls)
	if err2 != nil {
		panic(err)
	}
	return urls
}