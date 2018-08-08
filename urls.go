package pubtools

import (
	"net/url"
)

// URLLib struct
type URLLib struct{}

// URL returns URLLib
func URL() *URLLib {
	return &URLLib{}
}

// ResolveReference resolves new url by base url base and target
func (u *URLLib) ResolveReference(b, t string) (string, error) {
	base, err := url.Parse(b)
	if err != nil {
		return "", err
	}

	target, err := url.Parse(t)
	if err != nil {
		return "", err
	}

	ur := base.ResolveReference(target)
	return ur.String(), nil
}
