package rest

import (
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

func (h *Handler) parseBody(r *http.Request, req interface{}) bool {
	if r.Body == nil {
		return false
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.Warn(err)
		return false
	}
	err = jsoniter.Unmarshal(data, req)
	if err != nil {
		h.logger.Warn(err)
		return false
	}
	return true
}
