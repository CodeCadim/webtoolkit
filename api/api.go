package api

import (
	"encoding/json"
	"net/http"
)

type API struct {
	biz business
}

func (api API) Init(b business) {
	api.biz = b
	api.InitRoutes()
}

// respondJSON retourne une reponse json avec le statut et le contenu passés en paramètre
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}