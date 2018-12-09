package sites

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

type Sites struct {
	*pingaccess.Config
}

//SiteView - A site.
type SiteView struct {
	AvailabilityProfileId     int       `json:"availabilityProfileId",omitempty`
	ExpectedHostname          string    `json:"expectedHostname",omitempty`
	Id                        int       `json:"id",omitempty`
	KeepAliveTimeout          int       `json:"keepAliveTimeout",omitempty`
	LoadBalancingStrategyId   int       `json:"loadBalancingStrategyId",omitempty`
	MaxConnections            int       `json:"maxConnections",omitempty`
	MaxWebSocketConnections   int       `json:"maxWebSocketConnections",omitempty`
	Name                      string    `json:"name"`
	Secure                    bool      `json:"secure",omitempty`
	SendPaCookie              bool      `json:"sendPaCookie",omitempty`
	SiteAuthenticatorIds      []*int    `json:"siteAuthenticatorIds",omitempty`
	SkipHostnameVerification  bool      `json:"skipHostnameVerification",omitempty`
	Targets                   []*string `json:"targets"`
	TrustedCertificateGroupId int       `json:"trustedCertificateGroupId",omitempty`
	UseProxy                  bool      `json:"useProxy",omitempty`
	UseTargetHostHeader       bool      `json:"useTargetHostHeader",omitempty`
}

//SitesView - A collection of sites items.
type SitesView struct {
	Items []*SiteView `json:"items"`
}

// New creates a new instance of the Sites client with a config.
//
// Example:
//     // Create a Sites client from just a config.
//     svc := sites.New(cfg)
//
func New(cfg *pingaccess.Config) *Sites {
	svc := &Sites{
		Config: cfg,
	}
	return svc
}

//GetSitesCommand - Get all Sites
//RequestType: GET
//Input: input *GetSitesCommandInput
func (svc *Sites) GetSitesCommand(input *GetSitesCommandInput) (result *SitesView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/sites")

	log.Printf("[CLIENT] URL: %s", url)

	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(svc.Config.Username, svc.Config.Password)
	req.Header.Add("X-Xsrf-Header", "PingAccess")

	if err != nil {
		log.Println(err.Error())
		return result, err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println(err.Error())
		return result, err
	} else {
		defer resp.Body.Close()
		respBody, _ := ioutil.ReadAll(resp.Body)
		log.Printf("[CLIENT] ResponseBody: %v", string(respBody))
		json.Unmarshal(respBody, &result)
	}
	return result, nil
}

type GetSitesCommandInput struct {
	Params struct {
		Page          int
		NumberPerPage int
		Filter        string
		Name          string
		SortKey       string
		Order         string
	}
}

//AddSiteCommand - Create a Site
//RequestType: POST
//Input: input *AddSiteCommandInput
func (svc *Sites) AddSiteCommand(input *AddSiteCommandInput) (result *SiteView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/sites")

	log.Printf("[CLIENT] URL: %s", url)

	b, err := json.Marshal(input.Body)
	log.Printf("[CLIENT] RequestBody: %s", b)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.SetBasicAuth(svc.Config.Username, svc.Config.Password)
	req.Header.Add("X-Xsrf-Header", "PingAccess")

	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		log.Println(err.Error())
		return result, err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println(err.Error())
		return result, err
	} else {
		defer resp.Body.Close()
		respBody, _ := ioutil.ReadAll(resp.Body)
		log.Printf("[CLIENT] ResponseBody: %v", string(respBody))
		json.Unmarshal(respBody, &result)
	}
	return result, nil
}

type AddSiteCommandInput struct {
	Body SiteView
}

//DeleteSiteCommand - Delete a Site
//RequestType: DELETE
//Input: input *DeleteSiteCommandInput
func (svc *Sites) DeleteSiteCommand(input *DeleteSiteCommandInput) (err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/sites/{id}")

	url = strings.Replace(url, "{id}", input.Path.Id, -1)

	log.Printf("[CLIENT] URL: %s", url)

	req, err := http.NewRequest("DELETE", url, nil)
	req.SetBasicAuth(svc.Config.Username, svc.Config.Password)
	req.Header.Add("X-Xsrf-Header", "PingAccess")

	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = http.DefaultClient.Do(req)

	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

type DeleteSiteCommandInput struct {
	Path struct {
		Id string
	}
}

//GetSiteCommand - Get a Site
//RequestType: GET
//Input: input *GetSiteCommandInput
func (svc *Sites) GetSiteCommand(input *GetSiteCommandInput) (result *SiteView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/sites/{id}")

	url = strings.Replace(url, "{id}", input.Path.Id, -1)

	log.Printf("[CLIENT] URL: %s", url)

	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(svc.Config.Username, svc.Config.Password)
	req.Header.Add("X-Xsrf-Header", "PingAccess")

	if err != nil {
		log.Println(err.Error())
		return result, err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println(err.Error())
		return result, err
	} else {
		defer resp.Body.Close()
		respBody, _ := ioutil.ReadAll(resp.Body)
		log.Printf("[CLIENT] ResponseBody: %v", string(respBody))
		json.Unmarshal(respBody, &result)
	}
	return result, nil
}

type GetSiteCommandInput struct {
	Path struct {
		Id string
	}
}

//UpdateSiteCommand - Update a Site
//RequestType: PUT
//Input: input *UpdateSiteCommandInput
func (svc *Sites) UpdateSiteCommand(input *UpdateSiteCommandInput) (result *SiteView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/sites/{id}")

	url = strings.Replace(url, "{id}", input.Path.Id, -1)

	log.Printf("[CLIENT] URL: %s", url)

	b, err := json.Marshal(input.Body)
	log.Printf("[CLIENT] RequestBody: %s", b)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(b))
	req.SetBasicAuth(svc.Config.Username, svc.Config.Password)
	req.Header.Add("X-Xsrf-Header", "PingAccess")

	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		log.Println(err.Error())
		return result, err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println(err.Error())
		return result, err
	} else {
		defer resp.Body.Close()
		respBody, _ := ioutil.ReadAll(resp.Body)
		log.Printf("[CLIENT] ResponseBody: %v", string(respBody))
		json.Unmarshal(respBody, &result)
	}
	return result, nil
}

type UpdateSiteCommandInput struct {
	Body SiteView
	Path struct {
		Id string
	}
}
