package pingaccess

import (
	"net/http"
	"strings"
)

type ApplicationsService service

//GetApplicationsCommand - Get all Applications
//RequestType: GET
//Input: input *GetApplicationsCommandInput
func (s *ApplicationsService) GetApplicationsCommand(input *GetApplicationsCommandInput) (result *ApplicationsView, resp *http.Response, err error) {
	path := "/applications"

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
func (s *ApplicationsService) AddApplicationCommand(input *AddApplicationCommandInput) (result *ApplicationView, resp *http.Response, err error) {
	path := "/applications"

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

type AddApplicationCommandInput struct {
	Body ApplicationView
}

//GetResourcesCommand - Get all Resources
//RequestType: GET
//Input: input *GetResourcesCommandInput
func (s *ApplicationsService) GetResourcesCommand(input *GetResourcesCommandInput) (result *ResourcesView, resp *http.Response, err error) {
	path := "/applications/resources"

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

type GetResourcesCommandInput struct {
	Params struct {
		Page          int
		NumberPerPage int
	}
}

//GetApplicationsResourcesMethodsCommand - Get Application Resource Methods
//RequestType: GET
//Input:
func (s *ApplicationsService) GetApplicationsResourcesMethodsCommand() (result *MethodsView, resp *http.Response, err error) {
	path := "/applications/resources/methods"

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

//DeleteApplicationResourceCommand - Delete an Application Resource
//RequestType: DELETE
//Input: input *DeleteApplicationResourceCommandInput
func (s *ApplicationsService) DeleteApplicationResourceCommand(input *DeleteApplicationResourceCommandInput) (resp *http.Response, err error) {
	path := "/applications/{applicationId}/resources/{resourceId}"

	path = strings.Replace(path, "{applicationId}", input.Path.ApplicationId, -1)

	path = strings.Replace(path, "{resourceId}", input.Path.ResourceId, -1)

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

type DeleteApplicationResourceCommandInput struct {
	Path struct {
		ApplicationId string
		ResourceId    string
	}
}

//GetApplicationResourceCommand - Get an Application Resource
//RequestType: GET
//Input: input *GetApplicationResourceCommandInput
func (s *ApplicationsService) GetApplicationResourceCommand(input *GetApplicationResourceCommandInput) (result *ResourceView, resp *http.Response, err error) {
	path := "/applications/{applicationId}/resources/{resourceId}"

	path = strings.Replace(path, "{applicationId}", input.Path.ApplicationId, -1)

	path = strings.Replace(path, "{resourceId}", input.Path.ResourceId, -1)

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

type GetApplicationResourceCommandInput struct {
	Path struct {
		ApplicationId string
		ResourceId    string
	}
}

//UpdateApplicationResourceCommand - Update an Application Resource
//RequestType: PUT
//Input: input *UpdateApplicationResourceCommandInput
func (s *ApplicationsService) UpdateApplicationResourceCommand(input *UpdateApplicationResourceCommandInput) (result *ResourceView, resp *http.Response, err error) {
	path := "/applications/{applicationId}/resources/{resourceId}"

	path = strings.Replace(path, "{applicationId}", input.Path.ApplicationId, -1)

	path = strings.Replace(path, "{resourceId}", input.Path.ResourceId, -1)

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

type UpdateApplicationResourceCommandInput struct {
	Body ResourceView
	Path struct {
		ApplicationId string
		ResourceId    string
	}
}

//DeleteApplicationCommand - Delete an Application
//RequestType: DELETE
//Input: input *DeleteApplicationCommandInput
func (s *ApplicationsService) DeleteApplicationCommand(input *DeleteApplicationCommandInput) (result *ApplicationView, resp *http.Response, err error) {
	path := "/applications/{id}"

	path = strings.Replace(path, "{id}", input.Path.Id, -1)

	req, err := s.client.newRequest("DELETE", path, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type DeleteApplicationCommandInput struct {
	Path struct {
		Id string
	}
}

//GetApplicationCommand - Get an Application
//RequestType: GET
//Input: input *GetApplicationCommandInput
func (s *ApplicationsService) GetApplicationCommand(input *GetApplicationCommandInput) (result *ApplicationView, resp *http.Response, err error) {
	path := "/applications/{id}"

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

type GetApplicationCommandInput struct {
	Path struct {
		Id string
	}
}

//UpdateApplicationCommand - Update an Application
//RequestType: PUT
//Input: input *UpdateApplicationCommandInput
func (s *ApplicationsService) UpdateApplicationCommand(input *UpdateApplicationCommandInput) (result *ApplicationView, resp *http.Response, err error) {
	path := "/applications/{id}"

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

type UpdateApplicationCommandInput struct {
	Body ApplicationView
	Path struct {
		Id string
	}
}

//GetApplicationResourcesCommand - Get Resources for an Application
//RequestType: GET
//Input: input *GetApplicationResourcesCommandInput
func (s *ApplicationsService) GetApplicationResourcesCommand(input *GetApplicationResourcesCommandInput) (result *ResourceView, resp *http.Response, err error) {
	path := "/applications/{id}/resources"

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
		Id string
	}
}

//AddApplicationResourceCommand - Add Resource to an Application
//RequestType: POST
//Input: input *AddApplicationResourceCommandInput
func (s *ApplicationsService) AddApplicationResourceCommand(input *AddApplicationResourceCommandInput) (result *ResourceView, resp *http.Response, err error) {
	path := "/applications/{id}/resources"

	path = strings.Replace(path, "{id}", input.Path.Id, -1)

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

type AddApplicationResourceCommandInput struct {
	Body ResourceView
	Path struct {
		Id string
	}
}