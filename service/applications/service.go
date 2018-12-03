package applications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

type Applications struct {
	*pingaccess.Config
}

//ApplicationView - An application.
type ApplicationView struct {
	AccessValidatorId  int                    `json:"accessValidatorId",omitempty`
	AgentId            int                    `json:"agentId"`
	ApplicationType    string                 `json:"applicationType",omitempty`
	CaseSensitivePath  bool                   `json:"caseSensitivePath",omitempty`
	ContextRoot        string                 `json:"contextRoot"`
	DefaultAuthType    string                 `json:"defaultAuthType"`
	Description        string                 `json:"description",omitempty`
	Destination        string                 `json:"destination",omitempty`
	Enabled            bool                   `json:"enabled",omitempty`
	IdentityMappingIds map[string]int         `json:"identityMappingIds",omitempty`
	Name               string                 `json:"name"`
	Policy             map[string]interface{} `json:"policy",omitempty`
	Realm              string                 `json:"realm",omitempty`
	RequireHTTPS       bool                   `json:"requireHTTPS",omitempty`
	SiteId             int                    `json:"siteId"`
	VirtualHostIds     []*int                 `json:"virtualHostIds"`
	WebSessionId       int                    `json:"webSessionId",omitempty`
}

//ApplicationsView - A collection of applications.
type ApplicationsView struct {
	Items []*ApplicationView `json:"items"`
}

//MethodView - HTTP Method
type MethodView struct {
	Name string `json:"name"`
}

//MethodsView
type MethodsView struct {
	Items []*MethodView `json:"items"`
}

//ResourceView - A resource.
type ResourceView struct {
	Anonymous               bool                   `json:"anonymous",omitempty`
	ApplicationId           int                    `json:"applicationId",omitempty`
	AuditLevel              string                 `json:"auditLevel",omitempty`
	DefaultAuthTypeOverride string                 `json:"defaultAuthTypeOverride"`
	Enabled                 bool                   `json:"enabled",omitempty`
	Methods                 []*string              `json:"methods"`
	Name                    string                 `json:"name"`
	PathPrefixes            []*string              `json:"pathPrefixes"`
	Policy                  map[string]interface{} `json:"policy",omitempty`
	RootResource            bool                   `json:"rootResource",omitempty`
}

//ResourcesView
type ResourcesView struct {
	Items []*ResourceView `json:"items"`
}

// New creates a new instance of the Applications client with a config.
//
// Example:
//     // Create a Applications client from just a config.
//     svc := applications.New(cfg)
//
func New(cfg *pingaccess.Config) *Applications {
	svc := &Applications{
		Config: cfg,
	}
	return svc
}

//GetApplicationsCommand - Get all Applications
//RequestType: GET
//Input: input *GetApplicationsCommandInput
func (svc *Applications) GetApplicationsCommand(input *GetApplicationsCommandInput) (result *ApplicationsView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/applications")

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

type GetApplicationsCommandInput struct {
	Params struct {
		Page          int
		SiteId        int
		NumberPerPage int
		AgentId       int
		VirtualHostId int
		Filter        string
		Name          string
		SortKey       string
		Order         string
	}
}

//AddApplicationCommand - Add an Application
//RequestType: POST
//Input: input *AddApplicationCommandInput
func (svc *Applications) AddApplicationCommand(input *AddApplicationCommandInput) (result *ApplicationView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/applications")

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

type AddApplicationCommandInput struct {
	Body ApplicationView
}

//GetResourcesCommand - Get all Resources
//RequestType: GET
//Input: input *GetResourcesCommandInput
func (svc *Applications) GetResourcesCommand(input *GetResourcesCommandInput) (result *ResourcesView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/applications/resources")

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

type GetResourcesCommandInput struct {
	Params struct {
		Page          int
		NumberPerPage int
	}
}

//GetApplicationsResourcesMethodsCommand - Get Application Resource Methods
//RequestType: GET
//Input:
func (svc *Applications) GetApplicationsResourcesMethodsCommand() (result *MethodsView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/applications/resources/methods")

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

//DeleteApplicationResourceCommand - Delete an Application Resource
//RequestType: DELETE
//Input: input *DeleteApplicationResourceCommandInput
func (svc *Applications) DeleteApplicationResourceCommand(input *DeleteApplicationResourceCommandInput) (err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/applications/{applicationId}/resources/{resourceId}")

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

type DeleteApplicationResourceCommandInput struct {
	Path struct {
		applicationId string
		resourceId    string
	}
}

//GetApplicationResourceCommand - Get an Application Resource
//RequestType: GET
//Input: input *GetApplicationResourceCommandInput
func (svc *Applications) GetApplicationResourceCommand(input *GetApplicationResourceCommandInput) (result *ResourceView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/applications/{applicationId}/resources/{resourceId}")

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

type GetApplicationResourceCommandInput struct {
	Path struct {
		applicationId string
		resourceId    string
	}
}

//UpdateApplicationResourceCommand - Update an Application Resource
//RequestType: PUT
//Input: input *UpdateApplicationResourceCommandInput
func (svc *Applications) UpdateApplicationResourceCommand(input *UpdateApplicationResourceCommandInput) (result *ResourceView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/applications/{applicationId}/resources/{resourceId}")

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

type UpdateApplicationResourceCommandInput struct {
	Body ResourceView
	Path struct {
		applicationId string
		resourceId    string
	}
}

//DeleteApplicationCommand - Delete an Application
//RequestType: DELETE
//Input: input *DeleteApplicationCommandInput
func (svc *Applications) DeleteApplicationCommand(input *DeleteApplicationCommandInput) (result *ApplicationView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/applications/{id}")

	req, err := http.NewRequest("DELETE", url, nil)
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

type DeleteApplicationCommandInput struct {
	Path struct {
		id string
	}
}

//GetApplicationCommand - Get an Application
//RequestType: GET
//Input: input *GetApplicationCommandInput
func (svc *Applications) GetApplicationCommand(input *GetApplicationCommandInput) (result *ApplicationView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/applications/{id}")

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

type GetApplicationCommandInput struct {
	Path struct {
		id string
	}
}

//UpdateApplicationCommand - Update an Application
//RequestType: PUT
//Input: input *UpdateApplicationCommandInput
func (svc *Applications) UpdateApplicationCommand(input *UpdateApplicationCommandInput) (result *ApplicationView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/applications/{id}")

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

type UpdateApplicationCommandInput struct {
	Body ApplicationView
	Path struct {
		id string
	}
}

//GetApplicationResourcesCommand - Get Resources for an Application
//RequestType: GET
//Input: input *GetApplicationResourcesCommandInput
func (svc *Applications) GetApplicationResourcesCommand(input *GetApplicationResourcesCommandInput) (result *ResourceView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/applications/{id}/resources")

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

type GetApplicationResourcesCommandInput struct {
	Params struct {
		Page          int
		NumberPerPage int
		Name          string
		Filter        string
		SortKey       string
		Order         string
	}

	Path struct {
		id string
	}
}

//AddApplicationResourceCommand - Add Resource to an Application
//RequestType: POST
//Input: input *AddApplicationResourceCommandInput
func (svc *Applications) AddApplicationResourceCommand(input *AddApplicationResourceCommandInput) (result *ResourceView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/applications/{id}/resources")

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

type AddApplicationResourceCommandInput struct {
	Body ResourceView
	Path struct {
		id string
	}
}
