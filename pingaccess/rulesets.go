package pingaccess

import (
	"net/http"
	"strings"
)

type RulesetsService service

//GetRuleSetsCommand - Get all Rule Sets
//RequestType: GET
//Input: input *GetRuleSetsCommandInput
func (s *RulesetsService) GetRuleSetsCommand(input *GetRuleSetsCommandInput) (result *RuleSetsView, resp *http.Response, err error) {
	path := "/rulesets"

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

type GetRuleSetsCommandInput struct {
	Params struct {
		Page          int
		NumberPerPage int
		Filter        string
		Name          string
		SortKey       string
		Order         string
	}
}

//AddRuleSetCommand - Add a Rule Set
//RequestType: POST
//Input: input *AddRuleSetCommandInput
func (s *RulesetsService) AddRuleSetCommand(input *AddRuleSetCommandInput) (result *RuleSetView, resp *http.Response, err error) {
	path := "/rulesets"

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

type AddRuleSetCommandInput struct {
	Body RuleSetView
}

//GetRuleSetElementTypesCommand - Get all Rule Set Element Types
//RequestType: GET
//Input:
func (s *RulesetsService) GetRuleSetElementTypesCommand() (result *RuleSetElementTypesView, resp *http.Response, err error) {
	path := "/rulesets/elementTypes"

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

//GetRuleSetSuccessCriteriaCommand - Get all Success Criteria
//RequestType: GET
//Input:
func (s *RulesetsService) GetRuleSetSuccessCriteriaCommand() (result *RuleSetSuccessCriteriaView, resp *http.Response, err error) {
	path := "/rulesets/successCriteria"

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

//DeleteRuleSetCommand - Delete a Rule Set
//RequestType: DELETE
//Input: input *DeleteRuleSetCommandInput
func (s *RulesetsService) DeleteRuleSetCommand(input *DeleteRuleSetCommandInput) (resp *http.Response, err error) {
	path := "/rulesets/{id}"

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

type DeleteRuleSetCommandInput struct {
	Path struct {
		Id string
	}
}

//GetRuleSetCommand - Get a Rule Set
//RequestType: GET
//Input: input *GetRuleSetCommandInput
func (s *RulesetsService) GetRuleSetCommand(input *GetRuleSetCommandInput) (result *RuleSetView, resp *http.Response, err error) {
	path := "/rulesets/{id}"

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

type GetRuleSetCommandInput struct {
	Path struct {
		Id string
	}
}

//UpdateRuleSetCommand - Update a Rule Set
//RequestType: PUT
//Input: input *UpdateRuleSetCommandInput
func (s *RulesetsService) UpdateRuleSetCommand(input *UpdateRuleSetCommandInput) (result *RuleSetView, resp *http.Response, err error) {
	path := "/rulesets/{id}"

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

type UpdateRuleSetCommandInput struct {
	Body RuleSetView
	Path struct {
		Id string
	}
}
