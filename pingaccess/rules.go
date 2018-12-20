package pingaccess

import (
	"net/http"
	"strings"
)

type RulesService service

//GetRulesCommand - Get all Rules
//RequestType: GET
//Input: input *GetRulesCommandInput
func (s *RulesService) GetRulesCommand(input *GetRulesCommandInput) (result *RulesView, resp *http.Response, err error) {
	path := "/rules"

	req, err := s.client.newRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

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
func (s *RulesService) AddRuleCommand(input *AddRuleCommandInput) (result *RuleView, resp *http.Response, err error) {
	path := "/rules"

	req, err := s.client.newRequest("POST", path, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type AddRuleCommandInput struct {
	Body RuleView
}

//GetRuleDescriptorsCommand - Get descriptors for all supported Rule types
//RequestType: GET
//Input:
func (s *RulesService) GetRuleDescriptorsCommand() (result *RuleDescriptorsView, resp *http.Response, err error) {
	path := "/rules/descriptors"

	req, err := s.client.newRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetRuleDescriptorCommand - Get descriptor for a Rule type
//RequestType: GET
//Input: input *GetRuleDescriptorCommandInput
func (s *RulesService) GetRuleDescriptorCommand(input *GetRuleDescriptorCommandInput) (result *RuleDescriptor, resp *http.Response, err error) {
	path := "/rules/descriptors/{ruleType}"

	path = strings.Replace(path, "{ruleType}", input.Path.RuleType, -1)

	req, err := s.client.newRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type GetRuleDescriptorCommandInput struct {
	Path struct {
		RuleType string
	}
}

//DeleteRuleCommand - Delete a Rule
//RequestType: DELETE
//Input: input *DeleteRuleCommandInput
func (s *RulesService) DeleteRuleCommand(input *DeleteRuleCommandInput) (resp *http.Response, err error) {
	path := "/rules/{id}"

	path = strings.Replace(path, "{id}", input.Path.Id, -1)

	req, err := s.client.newRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

type DeleteRuleCommandInput struct {
	Path struct {
		Id string
	}
}

//GetRuleCommand - Get a Rule
//RequestType: GET
//Input: input *GetRuleCommandInput
func (s *RulesService) GetRuleCommand(input *GetRuleCommandInput) (result *RuleView, resp *http.Response, err error) {
	path := "/rules/{id}"

	path = strings.Replace(path, "{id}", input.Path.Id, -1)

	req, err := s.client.newRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type GetRuleCommandInput struct {
	Path struct {
		Id string
	}
}

//UpdateRuleCommand - Update a Rule
//RequestType: PUT
//Input: input *UpdateRuleCommandInput
func (s *RulesService) UpdateRuleCommand(input *UpdateRuleCommandInput) (result *RuleView, resp *http.Response, err error) {
	path := "/rules/{id}"

	path = strings.Replace(path, "{id}", input.Path.Id, -1)

	req, err := s.client.newRequest("PUT", path, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type UpdateRuleCommandInput struct {
	Body RuleView
	Path struct {
		Id string
	}
}
