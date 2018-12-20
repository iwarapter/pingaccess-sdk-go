package pingaccess

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	Username   string
	Password   string
	BaseURL    *url.URL
	httpClient *http.Client

	Applications       *ApplicationsService
	IdentityMappings   *IdentityMappingsService
	Rules              *RulesService
	SiteAuthenticators *SiteAuthenticatorsService
	Sites              *SitesService
	Virtualhosts       *VirtualhostsService
}

type service struct {
	client *Client
}

func NewClient(username string, password string, baseUrl *url.URL, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{httpClient: httpClient}
	c.Username = username
	c.Password = password
	c.BaseURL = baseUrl

	c.Applications = &ApplicationsService{client: c}
	c.IdentityMappings = &IdentityMappingsService{client: c}
	c.Rules = &RulesService{client: c}
	c.SiteAuthenticators = &SiteAuthenticatorsService{client: c}
	c.Sites = &SitesService{client: c}
	c.Virtualhosts = &VirtualhostsService{client: c}
	return c
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: fmt.Sprintf("pa-admin-api/v3%s", path)}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Add("X-Xsrf-Header", "PingAccess")
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	log.Println("[CLIENT] executing request")
	log.Printf("[CLIENT] METHOD: %s", req.Method)
	log.Printf("[CLIENT] URL: %s", req.URL.String())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}
	return resp, err
}
