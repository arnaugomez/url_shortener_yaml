package urls_routes

import (
	"encoding/json"
	"net/http"

	urls_data "github.com/arnaugomez/url_shortener_yaml/urls/data"
)

func PostUrl(w http.ResponseWriter, r *http.Request) {
	u, err := urlFromJson(r)
	enc := json.NewEncoder(w)
	if err != nil {
		enc.Encode(err.Error())
		return
	}

	err = urls_data.SaveUrl(&u)
	if err != nil {
		enc.Encode(err.Error())
		return
	}
	enc.Encode(u)
}
