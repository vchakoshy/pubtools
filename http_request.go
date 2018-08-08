package pubtools

import (
	"net/http"
	"strconv"
)

func GetRequestQuery(key string, r *http.Request) (string, bool) {
	vals := r.URL.Query()
	val, ok := vals[key]
	if !ok {
		return "", false
	}

	return val[0], true
}

func GetRequestIntQuery(key string, r *http.Request) (int, bool) {
	valToRe, ex := GetRequestQuery(key, r)

	if !ex {
		return 0, false
	}
	vtr, _ := strconv.Atoi(valToRe)
	return vtr, true
}
