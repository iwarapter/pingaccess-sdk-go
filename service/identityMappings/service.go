package identityMappings

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

type IdentityMappings struct {
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

//IdentityMappingDescriptor
type IdentityMappingDescriptor struct {
	ClassName           string                `json:"className"`
	ConfigurationFields []*ConfigurationField `json:"configurationFields"`
	Label               string                `json:"label"`
	Type                string                `json:"type"`
}

//IdentityMappingDescriptorsView
type IdentityMappingDescriptorsView struct {
	Items []*IdentityMappingDescriptor `json:"items"`
}

//IdentityMappingView
type IdentityMappingView struct {
	ClassName     string                 `json:"className"`
	Configuration map[string]interface{} `json:"configuration"`
	Name          string                 `json:"name"`
}

//IdentityMappingsView
type IdentityMappingsView struct {
	Items []*IdentityMappingView `json:"items"`
}

// New creates a new instance of the IdentityMappings client with a config.
//
// Example:
//     // Create a IdentityMappings client from just a config.
//     svc := identityMappings.New(cfg)
//
func New(cfg *pingaccess.Config) *IdentityMappings {
	svc := &IdentityMappings{
		Config: cfg,
	}
	return svc
}

//GetIdentityMappingsCommand - Get all Identity Mappings
//RequestType: GET
//Input: input *GetIdentityMappingsCommandInput
func (svc *IdentityMappings) GetIdentityMappingsCommand(input *GetIdentityMappingsCommandInput) (result *IdentityMappingsView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/identityMappings")

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

type GetIdentityMappingsCommandInput struct {
	Params struct {
		Page          int
		NumberPerPage int
		Filter        string
		Name          string
		SortKey       string
		Order         string
	}
}

//AddIdentityMappingCommand - Create an Identity Mapping
//RequestType: POST
//Input: input *AddIdentityMappingCommandInput
func (svc *IdentityMappings) AddIdentityMappingCommand(input *AddIdentityMappingCommandInput) (result *IdentityMappingView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/identityMappings")

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

type AddIdentityMappingCommandInput struct {
	Body IdentityMappingView
}

//GetIdentityMappingDescriptorsCommand - Get descriptors for all supported Identity Mapping types
//RequestType: GET
//Input:
func (svc *IdentityMappings) GetIdentityMappingDescriptorsCommand() (result *IdentityMappingDescriptorsView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/identityMappings/descriptors")

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

//GetIdentityMappingDescriptorCommand - Get descriptor for an Identity Mapping type
//RequestType: GET
//Input: input *GetIdentityMappingDescriptorCommandInput
func (svc *IdentityMappings) GetIdentityMappingDescriptorCommand(input *GetIdentityMappingDescriptorCommandInput) (result *IdentityMappingDescriptor, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/identityMappings/descriptors/{identityMappingType}")

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

type GetIdentityMappingDescriptorCommandInput struct {
	Path struct {
		identityMappingType string
	}
}

//DeleteIdentityMappingCommand - Delete an Identity Mapping
//RequestType: DELETE
//Input: input *DeleteIdentityMappingCommandInput
func (svc *IdentityMappings) DeleteIdentityMappingCommand(input *DeleteIdentityMappingCommandInput) (err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/identityMappings/{id}")

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

type DeleteIdentityMappingCommandInput struct {
	Path struct {
		id string
	}
}

//GetIdentityMappingCommand - Get an Identity Mapping
//RequestType: GET
//Input: input *GetIdentityMappingCommandInput
func (svc *IdentityMappings) GetIdentityMappingCommand(input *GetIdentityMappingCommandInput) (result *IdentityMappingView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/identityMappings/{id}")

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

type GetIdentityMappingCommandInput struct {
	Path struct {
		id string
	}
}

//UpdateIdentityMappingCommand - Update an Identity Mapping
//RequestType: PUT
//Input: input *UpdateIdentityMappingCommandInput
func (svc *IdentityMappings) UpdateIdentityMappingCommand(input *UpdateIdentityMappingCommandInput) (result *IdentityMappingView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/identityMappings/{id}")

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

type UpdateIdentityMappingCommandInput struct {
	Body IdentityMappingView
	Path struct {
		id string
	}
}
