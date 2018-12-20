package pingaccess

import (
	"net/http"
	"strings"
)

type SiteAuthenticatorsService service

//GetSiteAuthenticatorsCommand - Get all Site Authenticators
//RequestType: GET
//Input: input *GetSiteAuthenticatorsCommandInput
func (s *SiteAuthenticatorsService) GetSiteAuthenticatorsCommand(input *GetSiteAuthenticatorsCommandInput) (result *SiteAuthenticatorsView, resp *http.Response, err error) {
	path := "/siteAuthenticators"

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
func (s *SiteAuthenticatorsService) AddSiteAuthenticatorCommand(input *AddSiteAuthenticatorCommandInput) (result *SiteAuthenticatorView, resp *http.Response, err error) {
	path := "/siteAuthenticators"

	req, err := s.client.newRequest("POST", path, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type AddSiteAuthenticatorCommandInput struct {
	Body SiteAuthenticatorView
}

//GetSiteAuthenticatorDescriptorsCommand - Get descriptors for all supported Site Authenticator types
//RequestType: GET
//Input:
func (s *SiteAuthenticatorsService) GetSiteAuthenticatorDescriptorsCommand() (result *SiteAuthenticatorDescriptorsView, resp *http.Response, err error) {
	path := "/siteAuthenticators/descriptors"

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

//GetSiteAuthenticatorDescriptorCommand - Get descriptor for a Site Authenticator type
//RequestType: GET
//Input: input *GetSiteAuthenticatorDescriptorCommandInput
func (s *SiteAuthenticatorsService) GetSiteAuthenticatorDescriptorCommand(input *GetSiteAuthenticatorDescriptorCommandInput) (result *SiteAuthenticatorDescriptor, resp *http.Response, err error) {
	path := "/siteAuthenticators/descriptors/{siteAuthenticatorType}"

	path = strings.Replace(path, "{siteAuthenticatorType}", input.Path.SiteAuthenticatorType, -1)

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

type GetSiteAuthenticatorDescriptorCommandInput struct {
	Path struct {
		SiteAuthenticatorType string
	}
}

//DeleteSiteAuthenticatorCommand - Delete a Site Authenticator
//RequestType: DELETE
//Input: input *DeleteSiteAuthenticatorCommandInput
func (s *SiteAuthenticatorsService) DeleteSiteAuthenticatorCommand(input *DeleteSiteAuthenticatorCommandInput) (resp *http.Response, err error) {
	path := "/siteAuthenticators/{id}"

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

type DeleteSiteAuthenticatorCommandInput struct {
	Path struct {
		Id string
	}
}

//GetSiteAuthenticatorCommand - Get a Site Authenticator
//RequestType: GET
//Input: input *GetSiteAuthenticatorCommandInput
func (s *SiteAuthenticatorsService) GetSiteAuthenticatorCommand(input *GetSiteAuthenticatorCommandInput) (result *SiteAuthenticatorView, resp *http.Response, err error) {
	path := "/siteAuthenticators/{id}"

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

type GetSiteAuthenticatorCommandInput struct {
	Path struct {
		Id string
	}
}

//UpdateSiteAuthenticatorCommand - Update a Site Authenticator
//RequestType: PUT
//Input: input *UpdateSiteAuthenticatorCommandInput
func (s *SiteAuthenticatorsService) UpdateSiteAuthenticatorCommand(input *UpdateSiteAuthenticatorCommandInput) (result *SiteAuthenticatorView, resp *http.Response, err error) {
	path := "/siteAuthenticators/{id}"

	path = strings.Replace(path, "{id}", input.Path.Id, -1)

	req, err := s.client.newRequest("PUT", path, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type UpdateSiteAuthenticatorCommandInput struct {
	Body SiteAuthenticatorView
	Path struct {
		Id string
	}
}
