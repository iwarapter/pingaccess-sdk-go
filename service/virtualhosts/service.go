package virtualhosts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

type Virtualhosts struct {
	*pingaccess.Config
}

//VirtualHost
type VirtualHost struct {
	AgentResourceCacheTTL     int    `json:"agentResourceCacheTTL"`
	Host                      string `json:"host"`
	Id                        int    `json:"id"`
	KeyPairId                 int    `json:"keyPairId"`
	Port                      int    `json:"port"`
	TrustedCertificateGroupId int    `json:"trustedCertificateGroupId"`
}

//VirtualHostView
type VirtualHostView struct {
	AgentResourceCacheTTL     int    `json:"agentResourceCacheTTL",omitempty`
	Host                      string `json:"host"`
	KeyPairId                 int    `json:"keyPairId",omitempty`
	Port                      int    `json:"port"`
	TrustedCertificateGroupId int    `json:"trustedCertificateGroupId",omitempty`
}

//VirtualHostsView
type VirtualHostsView struct {
	Items []*VirtualHostView `json:"items"`
}

// New creates a new instance of the Virtualhosts client with a config.
//
// Example:
//     // Create a Virtualhosts client from just a config.
//     svc := virtualhosts.New(cfg)
//
func New(cfg *pingaccess.Config) *Virtualhosts {
	svc := &Virtualhosts{
		Config: cfg,
	}
	return svc
}

//GetVirtualHostsCommand - Get all Virtual Hosts
//RequestType: GET
//Input: input *GetVirtualHostsCommandInput
func (svc *Virtualhosts) GetVirtualHostsCommand(input *GetVirtualHostsCommandInput) (result *VirtualHostsView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/virtualhosts")

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
		json.Unmarshal(respBody, &result)
	}
	return result, nil
}

type GetVirtualHostsCommandInput struct {
	Params struct {
		Page          int
		NumberPerPage int
		Filter        string
		VirtualHost   string
		SortKey       string
		Order         string
	}
}

//AddVirtualHostCommand - Create a Virtual Host
//RequestType: POST
//Input: input *AddVirtualHostCommandInput
func (svc *Virtualhosts) AddVirtualHostCommand(input *AddVirtualHostCommandInput) (result *VirtualHostView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/virtualhosts")

	b, err := json.Marshal(input)
	log.Printf("[CLIENT] %s", b)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
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
		json.Unmarshal(respBody, &result)
	}
	return result, nil
}

type AddVirtualHostCommandInput struct {
	Body VirtualHostView
}

//DeleteVirtualHostCommand - Delete a Virtual Host
//RequestType: DELETE
//Input: input *DeleteVirtualHostCommandInput
func (svc *Virtualhosts) DeleteVirtualHostCommand(input *DeleteVirtualHostCommandInput) (err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/virtualhosts/{id}")

	req, err := http.NewRequest("DELETE", url, nil)
	req.SetBasicAuth(svc.Config.Username, svc.Config.Password)
	req.Header.Add("X-Xsrf-Header", "PingAccess")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

type DeleteVirtualHostCommandInput struct {
	Path struct {
		id string
	}
}

//GetVirtualHostCommand - Get a Virtual Host
//RequestType: GET
//Input: input *GetVirtualHostCommandInput
func (svc *Virtualhosts) GetVirtualHostCommand(input *GetVirtualHostCommandInput) (result *VirtualHost, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/virtualhosts/{id}")

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
		json.Unmarshal(respBody, &result)
	}
	return result, nil
}

type GetVirtualHostCommandInput struct {
	Path struct {
		id string
	}
}

//UpdateVirtualHostCommand - Update a Virtual Host
//RequestType: PUT
//Input: input *UpdateVirtualHostCommandInput
func (svc *Virtualhosts) UpdateVirtualHostCommand(input *UpdateVirtualHostCommandInput) (result *VirtualHostView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/virtualhosts/{id}")

	b, err := json.Marshal(input)
	log.Printf("[CLIENT] %s", b)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(b))
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
		json.Unmarshal(respBody, &result)
	}
	return result, nil
}

type UpdateVirtualHostCommandInput struct {
	Body VirtualHostView
	Path struct {
		id string
	}
}
