package urls_routes

import (
	"encoding/json"
	"net/http"

	urls_data "github.com/arnaugomez/url_shortener_yaml/urls/data"
	"github.com/gorilla/mux"
)

func GetUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p := vars["path"]
	u, err := urls_data.FindUrl(p)

	enc := json.NewEncoder(w)
	if err != nil {
		enc.Encode(err.Error())
		return
	}
	http.Redirect(w, r, u.Url, http.StatusPermanentRedirect)
}
