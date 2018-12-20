package pingaccess

import (
	"net/http"
	"strings"
)

type IdentityMappingsService service

//GetIdentityMappingsCommand - Get all Identity Mappings
//RequestType: GET
//Input: input *GetIdentityMappingsCommandInput
func (s *IdentityMappingsService) GetIdentityMappingsCommand(input *GetIdentityMappingsCommandInput) (result *IdentityMappingsView, resp *http.Response, err error) {
	path := "/identityMappings"

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
func (s *IdentityMappingsService) AddIdentityMappingCommand(input *AddIdentityMappingCommandInput) (result *IdentityMappingView, resp *http.Response, err error) {
	path := "/identityMappings"

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

type AddIdentityMappingCommandInput struct {
	Body IdentityMappingView
}

//GetIdentityMappingDescriptorsCommand - Get descriptors for all supported Identity Mapping types
//RequestType: GET
//Input:
func (s *IdentityMappingsService) GetIdentityMappingDescriptorsCommand() (result *IdentityMappingDescriptorsView, resp *http.Response, err error) {
	path := "/identityMappings/descriptors"

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

//GetIdentityMappingDescriptorCommand - Get descriptor for an Identity Mapping type
//RequestType: GET
//Input: input *GetIdentityMappingDescriptorCommandInput
func (s *IdentityMappingsService) GetIdentityMappingDescriptorCommand(input *GetIdentityMappingDescriptorCommandInput) (result *IdentityMappingDescriptor, resp *http.Response, err error) {
	path := "/identityMappings/descriptors/{identityMappingType}"

	path = strings.Replace(path, "{identityMappingType}", input.Path.IdentityMappingType, -1)

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

type GetIdentityMappingDescriptorCommandInput struct {
	Path struct {
		IdentityMappingType string
	}
}

//DeleteIdentityMappingCommand - Delete an Identity Mapping
//RequestType: DELETE
//Input: input *DeleteIdentityMappingCommandInput
func (s *IdentityMappingsService) DeleteIdentityMappingCommand(input *DeleteIdentityMappingCommandInput) (resp *http.Response, err error) {
	path := "/identityMappings/{id}"

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

type DeleteIdentityMappingCommandInput struct {
	Path struct {
		Id string
	}
}

//GetIdentityMappingCommand - Get an Identity Mapping
//RequestType: GET
//Input: input *GetIdentityMappingCommandInput
func (s *IdentityMappingsService) GetIdentityMappingCommand(input *GetIdentityMappingCommandInput) (result *IdentityMappingView, resp *http.Response, err error) {
	path := "/identityMappings/{id}"

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

type GetIdentityMappingCommandInput struct {
	Path struct {
		Id string
	}
}

//UpdateIdentityMappingCommand - Update an Identity Mapping
//RequestType: PUT
//Input: input *UpdateIdentityMappingCommandInput
func (s *IdentityMappingsService) UpdateIdentityMappingCommand(input *UpdateIdentityMappingCommandInput) (result *IdentityMappingView, resp *http.Response, err error) {
	path := "/identityMappings/{id}"

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

type UpdateIdentityMappingCommandInput struct {
	Body IdentityMappingView
	Path struct {
		Id string
	}
}
