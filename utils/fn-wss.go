package utils

import (
	"encoding/json"
	"net/http"
)

func FetchPayloadData(r *http.Request) (payload map[string]interface{}, err error) {
	payload = make(map[string]interface{})
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&payload); err != nil {
		return payload, err
	}
	return payload, nil
}
func HttpResponseError(err error) string {
	var maps = map[string]interface{}{
		"error": err.Error(),
	}
	return MapToString(maps)
}
func MapToString(input map[string]interface{}) string {
	b, _ := json.Marshal(input)
	return string(b)
}
