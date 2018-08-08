package pubtools

import (
	"net/http"
	"strconv"
)

// HTTPRequestLib struct
type HTTPRequestLib struct {
	r *http.Request
}

// HTTPRequest returns HTTPRequestLib
func HTTPRequest(r *http.Request) *HTTPRequestLib {
	return &HTTPRequestLib{
		r: r,
	}
}

// GetQuery returns guery value
func (h *HTTPRequestLib) GetQuery(key string) (string, bool) {
	vals := h.r.URL.Query()
	val, ok := vals[key]
	if !ok {
		return "", false
	}

	return val[0], true
}

// GetIntQuery returns guery int value
func (h *HTTPRequestLib) GetIntQuery(key string) (int, bool) {
	valToRe, ex := h.GetQuery(key)

	if !ex {
		return 0, false
	}
	vtr, _ := strconv.Atoi(valToRe)
	return vtr, true
}
