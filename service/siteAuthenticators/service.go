package siteAuthenticators

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

type SiteAuthenticators struct {
	*pingaccess.Config
}

//ConfigurationDependentFieldOption - Configuration of the dependent field option.
type ConfigurationDependentFieldOption struct {
	Options []*ConfigurationOption `json:"options"`
	Value   string                 `json:"value"`
}

//ConfigurationField - Details for configuration in the administrator console.
type ConfigurationField struct {
	Advanced     bool                     `json:"advanced"`
	ButtonGroup  string                   `json:"buttonGroup"`
	Default      string                   `json:"default"`
	Deselectable bool                     `json:"deselectable"`
	Fields       []*ConfigurationField    `json:"fields"`
	Help         Help                     `json:"help"`
	Label        string                   `json:"label"`
	Name         string                   `json:"name"`
	Options      []*ConfigurationOption   `json:"options"`
	ParentField  ConfigurationParentField `json:"parentField"`
	Required     bool                     `json:"required"`
	Type         string                   `json:"type"`
}

//ConfigurationOption - The configuration option attributes.
type ConfigurationOption struct {
	Category string `json:"category"`
	Label    string `json:"label"`
	Value    string `json:"value"`
}

//ConfigurationParentField - Configuration of the parent field.
type ConfigurationParentField struct {
	DependentFieldOptions []*ConfigurationDependentFieldOption `json:"dependentFieldOptions"`
	FieldName             string                               `json:"fieldName"`
}

//Help - Configuration of the help context of a configuration field.
type Help struct {
	Content string `json:"content"`
	Title   string `json:"title"`
	Url     string `json:"url"`
}

//SiteAuthenticatorDescriptor
type SiteAuthenticatorDescriptor struct {
	ClassName           string                `json:"className"`
	ConfigurationFields []*ConfigurationField `json:"configurationFields"`
	Label               string                `json:"label"`
	Type                string                `json:"type"`
}

//SiteAuthenticatorDescriptorsView
type SiteAuthenticatorDescriptorsView struct {
	Items []*SiteAuthenticatorDescriptor `json:"items"`
}

//SiteAuthenticatorView
type SiteAuthenticatorView struct {
	ClassName     string                 `json:"className"`
	Configuration map[string]interface{} `json:"configuration"`
	Id            int                    `json:"id",omitempty`
	Name          string                 `json:"name"`
}

//SiteAuthenticatorsView
type SiteAuthenticatorsView struct {
	Items []*SiteAuthenticatorView `json:"items"`
}

// New creates a new instance of the SiteAuthenticators client with a config.
//
// Example:
//     // Create a SiteAuthenticators client from just a config.
//     svc := siteAuthenticators.New(cfg)
//
func New(cfg *pingaccess.Config) *SiteAuthenticators {
	svc := &SiteAuthenticators{
		Config: cfg,
	}
	return svc
}

//GetSiteAuthenticatorsCommand - Get all Site Authenticators
//RequestType: GET
//Input: input *GetSiteAuthenticatorsCommandInput
func (svc *SiteAuthenticators) GetSiteAuthenticatorsCommand(input *GetSiteAuthenticatorsCommandInput) (result *SiteAuthenticatorsView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/siteAuthenticators")

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

type GetSiteAuthenticatorsCommandInput struct {
	Params struct {
		Page          int
		NumberPerPage int
		Filter        string
		Name          string
		SortKey       string
		Order         string
	}
}

//AddSiteAuthenticatorCommand - Create a Site Authenticator
//RequestType: POST
//Input: input *AddSiteAuthenticatorCommandInput
func (svc *SiteAuthenticators) AddSiteAuthenticatorCommand(input *AddSiteAuthenticatorCommandInput) (result *SiteAuthenticatorView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/siteAuthenticators")

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

type AddSiteAuthenticatorCommandInput struct {
	Body SiteAuthenticatorView
}

//GetSiteAuthenticatorDescriptorsCommand - Get descriptors for all supported Site Authenticator types
//RequestType: GET
//Input:
func (svc *SiteAuthenticators) GetSiteAuthenticatorDescriptorsCommand() (result *SiteAuthenticatorDescriptorsView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/siteAuthenticators/descriptors")

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

//GetSiteAuthenticatorDescriptorCommand - Get descriptor for a Site Authenticator type
//RequestType: GET
//Input: input *GetSiteAuthenticatorDescriptorCommandInput
func (svc *SiteAuthenticators) GetSiteAuthenticatorDescriptorCommand(input *GetSiteAuthenticatorDescriptorCommandInput) (result *SiteAuthenticatorDescriptor, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/siteAuthenticators/descriptors/{siteAuthenticatorType}")

	url = strings.Replace(url, "{siteAuthenticatorType}", input.Path.SiteAuthenticatorType, -1)

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

type GetSiteAuthenticatorDescriptorCommandInput struct {
	Path struct {
		SiteAuthenticatorType string
	}
}

//DeleteSiteAuthenticatorCommand - Delete a Site Authenticator
//RequestType: DELETE
//Input: input *DeleteSiteAuthenticatorCommandInput
func (svc *SiteAuthenticators) DeleteSiteAuthenticatorCommand(input *DeleteSiteAuthenticatorCommandInput) (err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/siteAuthenticators/{id}")

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

type DeleteSiteAuthenticatorCommandInput struct {
	Path struct {
		Id string
	}
}

//GetSiteAuthenticatorCommand - Get a Site Authenticator
//RequestType: GET
//Input: input *GetSiteAuthenticatorCommandInput
func (svc *SiteAuthenticators) GetSiteAuthenticatorCommand(input *GetSiteAuthenticatorCommandInput) (result *SiteAuthenticatorView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/siteAuthenticators/{id}")

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

type GetSiteAuthenticatorCommandInput struct {
	Path struct {
		Id string
	}
}

//UpdateSiteAuthenticatorCommand - Update a Site Authenticator
//RequestType: PUT
//Input: input *UpdateSiteAuthenticatorCommandInput
func (svc *SiteAuthenticators) UpdateSiteAuthenticatorCommand(input *UpdateSiteAuthenticatorCommandInput) (result *SiteAuthenticatorView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/siteAuthenticators/{id}")

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

type UpdateSiteAuthenticatorCommandInput struct {
	Body SiteAuthenticatorView
	Path struct {
		Id string
	}
}
