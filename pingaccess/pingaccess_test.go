package pingaccess

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

func config() *Client {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	url, _ := url.Parse("https://localhost:9000/")
	return NewClient("Administrator", "2Access2", url, nil)
}
