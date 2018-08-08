package pubtools

import (
	"net/url"
)

// URLResolveReference resolves new url by base url base and target
func URLResolveReference(b, t string) (string, error) {
	base, err := url.Parse(b)
	if err != nil {
		return "", err
	}

	target, err := url.Parse(t)
	if err != nil {
		return "", err
	}

	u := base.ResolveReference(target)
	return u.String(), nil
}
