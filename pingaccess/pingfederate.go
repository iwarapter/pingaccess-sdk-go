package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
)

type PingfederateService service

type PingfederateAPI interface {
	DeletePingFederateCommand() (resp *http.Response, err error)
	GetPingFederateCommand() (result *PingFederateRuntimeView, resp *http.Response, err error)
	UpdatePingFederateCommand(input *UpdatePingFederateCommandInput) (result *PingFederateRuntimeView, resp *http.Response, err error)
	DeletePingFederateAccessTokensCommand() (resp *http.Response, err error)
	GetPingFederateAccessTokensCommand() (result *PingFederateAccessTokenView, resp *http.Response, err error)
	UpdatePingFederateAccessTokensCommand(input *UpdatePingFederateAccessTokensCommandInput) (result *PingFederateAccessTokenView, resp *http.Response, err error)
	DeletePingFederateAdminCommand() (resp *http.Response, err error)
	GetPingFederateAdminCommand() (result *PingFederateAdminView, resp *http.Response, err error)
	UpdatePingFederateAdminCommand(input *UpdatePingFederateAdminCommandInput) (result *PingFederateAdminView, resp *http.Response, err error)
	GetLegacyPingFederateMetadataCommand() (result *OIDCProviderMetadata, resp *http.Response, err error)
	DeletePingFederateRuntimeCommand() (resp *http.Response, err error)
	GetPingFederateRuntimeCommand() (result *PingFederateMetadataRuntimeView, resp *http.Response, err error)
	UpdatePingFederateRuntimeCommand(input *UpdatePingFederateRuntimeCommandInput) (result *PingFederateMetadataRuntimeView, resp *http.Response, err error)
	GetPingFederateMetadataCommand() (result *OIDCProviderMetadata, resp *http.Response, err error)
}

//DeletePingFederateCommand - [Attention: This endpoint "/pingfederate" is deprecated. Please use /pingfederate/runtime to configure PingFederate instead] Resets the PingFederate configuration to default values
//RequestType: DELETE
//Input:
func (s *PingfederateService) DeletePingFederateCommand() (resp *http.Response, err error) {
	path := "/pingfederate"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

//GetPingFederateCommand - [Attention: This endpoint "/pingfederate" is deprecated. Please use /pingfederate/runtime to configure PingFederate instead] Get the PingFederate configuration
//RequestType: GET
//Input:
func (s *PingfederateService) GetPingFederateCommand() (result *PingFederateRuntimeView, resp *http.Response, err error) {
	path := "/pingfederate"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdatePingFederateCommand - [Attention: This endpoint "/pingfederate" is deprecated. Please use /pingfederate/runtime to configure PingFederate instead] Update the PingFederate configuration
//RequestType: PUT
//Input: input *UpdatePingFederateCommandInput
func (s *PingfederateService) UpdatePingFederateCommand(input *UpdatePingFederateCommandInput) (result *PingFederateRuntimeView, resp *http.Response, err error) {
	path := "/pingfederate"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type UpdatePingFederateCommandInput struct {
	Body PingFederateRuntimeView
}

//DeletePingFederateAccessTokensCommand - Resets the PingAccess OAuth Client configuration to default values
//RequestType: DELETE
//Input:
func (s *PingfederateService) DeletePingFederateAccessTokensCommand() (resp *http.Response, err error) {
	path := "/pingfederate/accessTokens"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

//GetPingFederateAccessTokensCommand - Get the PingAccess OAuth Client configuration
//RequestType: GET
//Input:
func (s *PingfederateService) GetPingFederateAccessTokensCommand() (result *PingFederateAccessTokenView, resp *http.Response, err error) {
	path := "/pingfederate/accessTokens"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdatePingFederateAccessTokensCommand - Update the PingFederate OAuth Client configuration
//RequestType: PUT
//Input: input *UpdatePingFederateAccessTokensCommandInput
func (s *PingfederateService) UpdatePingFederateAccessTokensCommand(input *UpdatePingFederateAccessTokensCommandInput) (result *PingFederateAccessTokenView, resp *http.Response, err error) {
	path := "/pingfederate/accessTokens"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type UpdatePingFederateAccessTokensCommandInput struct {
	Body PingFederateAccessTokenView
}

//DeletePingFederateAdminCommand - Resets the PingFederate Admin configuration to default values
//RequestType: DELETE
//Input:
func (s *PingfederateService) DeletePingFederateAdminCommand() (resp *http.Response, err error) {
	path := "/pingfederate/admin"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

//GetPingFederateAdminCommand - Get the PingFederate Admin configuration
//RequestType: GET
//Input:
func (s *PingfederateService) GetPingFederateAdminCommand() (result *PingFederateAdminView, resp *http.Response, err error) {
	path := "/pingfederate/admin"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdatePingFederateAdminCommand - Update the PingFederate Admin configuration
//RequestType: PUT
//Input: input *UpdatePingFederateAdminCommandInput
func (s *PingfederateService) UpdatePingFederateAdminCommand(input *UpdatePingFederateAdminCommandInput) (result *PingFederateAdminView, resp *http.Response, err error) {
	path := "/pingfederate/admin"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type UpdatePingFederateAdminCommandInput struct {
	Body PingFederateAdminView
}

//GetLegacyPingFederateMetadataCommand - [Attention: The endpoint "/pingfederate" is deprecated. This metadata corresponds to that configuration."/pingfederate/runtime" and "/pingfederate/runtime/metadata" should be used instead.] Get the PingFederate metadata
//RequestType: GET
//Input:
func (s *PingfederateService) GetLegacyPingFederateMetadataCommand() (result *OIDCProviderMetadata, resp *http.Response, err error) {
	path := "/pingfederate/metadata"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeletePingFederateRuntimeCommand - Resets the PingFederate configuration
//RequestType: DELETE
//Input:
func (s *PingfederateService) DeletePingFederateRuntimeCommand() (resp *http.Response, err error) {
	path := "/pingfederate/runtime"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

//GetPingFederateRuntimeCommand - Get the PingFederate configuration
//RequestType: GET
//Input:
func (s *PingfederateService) GetPingFederateRuntimeCommand() (result *PingFederateMetadataRuntimeView, resp *http.Response, err error) {
	path := "/pingfederate/runtime"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdatePingFederateRuntimeCommand - Update the PingFederate configuration
//RequestType: PUT
//Input: input *UpdatePingFederateRuntimeCommandInput
func (s *PingfederateService) UpdatePingFederateRuntimeCommand(input *UpdatePingFederateRuntimeCommandInput) (result *PingFederateMetadataRuntimeView, resp *http.Response, err error) {
	path := "/pingfederate/runtime"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type UpdatePingFederateRuntimeCommandInput struct {
	Body PingFederateMetadataRuntimeView
}

//GetPingFederateMetadataCommand - Get the PingFederate Runtime metadata
//RequestType: GET
//Input:
func (s *PingfederateService) GetPingFederateMetadataCommand() (result *OIDCProviderMetadata, resp *http.Response, err error) {
	path := "/pingfederate/runtime/metadata"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}
