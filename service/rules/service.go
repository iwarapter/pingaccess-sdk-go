package rules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

type Rules struct {
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

//RuleDescriptor
type RuleDescriptor struct {
	AgentCachingDisabled bool                  `json:"agentCachingDisabled"`
	Category             string                `json:"category"`
	ClassName            string                `json:"className"`
	ConfigurationFields  []*ConfigurationField `json:"configurationFields"`
	Label                string                `json:"label"`
	Modes                string                `json:"modes"`
	Type                 string                `json:"type"`
}

//RuleDescriptorsView
type RuleDescriptorsView struct {
	Items []*RuleDescriptor `json:"items"`
}

//RuleView
type RuleView struct {
	ClassName             string                 `json:"className"`
	Configuration         map[string]interface{} `json:"configuration"`
	Name                  string                 `json:"name"`
	SupportedDestinations []*string              `json:"supportedDestinations",omitempty`
}

//RulesView
type RulesView struct {
	Items []*RuleView `json:"items"`
}

// New creates a new instance of the Rules client with a config.
//
// Example:
//     // Create a Rules client from just a config.
//     svc := rules.New(cfg)
//
func New(cfg *pingaccess.Config) *Rules {
	svc := &Rules{
		Config: cfg,
	}
	return svc
}

//GetRulesCommand - Get all Rules
//RequestType: GET
//Input: input *GetRulesCommandInput
func (svc *Rules) GetRulesCommand(input *GetRulesCommandInput) (result *RulesView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/rules")

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

type GetRulesCommandInput struct {
	Params struct {
		Page          int
		NumberPerPage int
		Filter        string
		Name          string
		SortKey       string
		Order         string
	}
}

//AddRuleCommand - Add a Rule
//RequestType: POST
//Input: input *AddRuleCommandInput
func (svc *Rules) AddRuleCommand(input *AddRuleCommandInput) (result *RuleView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/rules")

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

type AddRuleCommandInput struct {
	Body RuleView
}

//GetRuleDescriptorsCommand - Get descriptors for all supported Rule types
//RequestType: GET
//Input:
func (svc *Rules) GetRuleDescriptorsCommand() (result *RuleDescriptorsView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/rules/descriptors")

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

//GetRuleDescriptorCommand - Get descriptor for a Rule type
//RequestType: GET
//Input: input *GetRuleDescriptorCommandInput
func (svc *Rules) GetRuleDescriptorCommand(input *GetRuleDescriptorCommandInput) (result *RuleDescriptor, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/rules/descriptors/{ruleType}")

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

type GetRuleDescriptorCommandInput struct {
	Path struct {
		ruleType string
	}
}

//DeleteRuleCommand - Delete a Rule
//RequestType: DELETE
//Input: input *DeleteRuleCommandInput
func (svc *Rules) DeleteRuleCommand(input *DeleteRuleCommandInput) (err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/rules/{id}")

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

type DeleteRuleCommandInput struct {
	Path struct {
		id string
	}
}

//GetRuleCommand - Get a Rule
//RequestType: GET
//Input: input *GetRuleCommandInput
func (svc *Rules) GetRuleCommand(input *GetRuleCommandInput) (result *RuleView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/rules/{id}")

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

type GetRuleCommandInput struct {
	Path struct {
		id string
	}
}

//UpdateRuleCommand - Update a Rule
//RequestType: PUT
//Input: input *UpdateRuleCommandInput
func (svc *Rules) UpdateRuleCommand(input *UpdateRuleCommandInput) (result *RuleView, err error) {
	url := fmt.Sprintf("%s%s", svc.Config.BaseURL, "/rules/{id}")

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

type UpdateRuleCommandInput struct {
	Body RuleView
	Path struct {
		id string
	}
}
