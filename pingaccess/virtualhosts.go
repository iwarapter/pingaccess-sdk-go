package pingaccess

import (
	"net/http"
	"strings"
)

type VirtualhostsService service

//GetVirtualHostsCommand - Get all Virtual Hosts
//RequestType: GET
//Input: input *GetVirtualHostsCommandInput
func (s *VirtualhostsService) GetVirtualHostsCommand(input *GetVirtualHostsCommandInput) (result *VirtualHostsView, resp *http.Response, err error) {
	path := "/virtualhosts"

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

type GetVirtualHostsCommandInput struct {
	Params struct {
		Page          int
		NumberPerPage int
		Filter        string
		VirtualHost   string
		SortKey       string
		Order         string
	}
}

//AddVirtualHostCommand - Create a Virtual Host
//RequestType: POST
//Input: input *AddVirtualHostCommandInput
func (s *VirtualhostsService) AddVirtualHostCommand(input *AddVirtualHostCommandInput) (result *VirtualHostView, resp *http.Response, err error) {
	path := "/virtualhosts"

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

type AddVirtualHostCommandInput struct {
	Body VirtualHostView
}

//DeleteVirtualHostCommand - Delete a Virtual Host
//RequestType: DELETE
//Input: input *DeleteVirtualHostCommandInput
func (s *VirtualhostsService) DeleteVirtualHostCommand(input *DeleteVirtualHostCommandInput) (resp *http.Response, err error) {
	path := "/virtualhosts/{id}"

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

type DeleteVirtualHostCommandInput struct {
	Path struct {
		Id string
	}
}

//GetVirtualHostCommand - Get a Virtual Host
//RequestType: GET
//Input: input *GetVirtualHostCommandInput
func (s *VirtualhostsService) GetVirtualHostCommand(input *GetVirtualHostCommandInput) (result *VirtualHost, resp *http.Response, err error) {
	path := "/virtualhosts/{id}"

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

type GetVirtualHostCommandInput struct {
	Path struct {
		Id string
	}
}

//UpdateVirtualHostCommand - Update a Virtual Host
//RequestType: PUT
//Input: input *UpdateVirtualHostCommandInput
func (s *VirtualhostsService) UpdateVirtualHostCommand(input *UpdateVirtualHostCommandInput) (result *VirtualHostView, resp *http.Response, err error) {
	path := "/virtualhosts/{id}"

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

type UpdateVirtualHostCommandInput struct {
	Body VirtualHostView
	Path struct {
		Id string
	}
}
