package util

import (
	"net/url"
)

func HostnameOfUrl(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return u.Host
}

func RelativePathToAbsolutePath(href string, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return " "
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return " "
	}
	return baseUrl.ResolveReference(uri).String()

}
