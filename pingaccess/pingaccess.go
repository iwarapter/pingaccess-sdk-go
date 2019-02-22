package pingaccess

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	Username   string
	Password   string
	BaseURL    *url.URL
	Context    string
	httpClient *http.Client

	Applications       *ApplicationsService
	IdentityMappings   *IdentityMappingsService
	Rules              *RulesService
	Rulesets           *RulesetsService
	SiteAuthenticators *SiteAuthenticatorsService
	Sites              *SitesService
	Virtualhosts       *VirtualhostsService
	WebSessions        *WebSessionsService
}

type service struct {
	client *Client
}

func NewClient(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{httpClient: httpClient}
	c.Username = username
	c.Password = password
	c.BaseURL = baseUrl
	c.Context = context

	c.Applications = &ApplicationsService{client: c}
	c.IdentityMappings = &IdentityMappingsService{client: c}
	c.Rules = &RulesService{client: c}
	c.Rulesets = &RulesetsService{client: c}
	c.SiteAuthenticators = &SiteAuthenticatorsService{client: c}
	c.Sites = &SitesService{client: c}
	c.Virtualhosts = &VirtualhostsService{client: c}
	c.WebSessions = &WebSessionsService{client: c}
	return c
}

func (c *Client) newRequest(method string, path *url.URL, body interface{}) (*http.Request, error) {
	u := c.BaseURL.ResolveReference(path)
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
	// log.Println("[CLIENT] executing request")
	// log.Printf("[CLIENT] METHOD: %s", req.Method)
	// log.Printf("[CLIENT] URL: %s", req.URL.String())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()
	// if resp.StatusCode < 200 && resp.StatusCode > 299 {
	// 	if w, ok := v.(io.Writer); ok {
	// 		io.Copy(w, resp.Body)
	// 	} else {
	// 		err = json.NewDecoder(resp.Body).Decode(failure)
	// 	}
	// 	return resp, failure, err
	// }
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

type FailureResponse struct {
	Form  map[string][]string `json:"form"`
	Flash []string            `json:"flash"`
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range.
// API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other
// response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := FailureResponse{}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, &errorResponse)
	}

	return &PingAccessError{
		Response: errorResponse,
	}
}

// PingAccessError occurs when PingAccess returns a non 2XX response
type PingAccessError struct {
	Response FailureResponse
}

func (r *PingAccessError) Error() (message string) {
	if r.Response.Flash != nil {
		message = strings.Join(r.Response.Flash, "\n")
	}
	for _, v := range r.Response.Form {
		for _, s := range v {
			message = fmt.Sprintf(message + s + "\n")
		}
	}
	return
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
