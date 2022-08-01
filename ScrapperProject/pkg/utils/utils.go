package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(req *http.Request, in interface{}) {
	if body, err := ioutil.ReadAll(req.Body); err == nil {
		if err := json.Unmarshal([]byte(body), in); err != nil {
			return
		}
	}
}
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}
