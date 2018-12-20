package pingaccess

import (
	"net/http"
	"strings"
)

type SitesService service

//GetSitesCommand - Get all Sites
//RequestType: GET
//Input: input *GetSitesCommandInput
func (s *SitesService) GetSitesCommand(input *GetSitesCommandInput) (result *SitesView, resp *http.Response, err error) {
	path := "/sites"

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

type GetSitesCommandInput struct {
	Params struct {
		Page          int
		NumberPerPage int
		Filter        string
		Name          string
		SortKey       string
		Order         string
	}
}

//AddSiteCommand - Create a Site
//RequestType: POST
//Input: input *AddSiteCommandInput
func (s *SitesService) AddSiteCommand(input *AddSiteCommandInput) (result *SiteView, resp *http.Response, err error) {
	path := "/sites"

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

type AddSiteCommandInput struct {
	Body SiteView
}

//DeleteSiteCommand - Delete a Site
//RequestType: DELETE
//Input: input *DeleteSiteCommandInput
func (s *SitesService) DeleteSiteCommand(input *DeleteSiteCommandInput) (resp *http.Response, err error) {
	path := "/sites/{id}"

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

type DeleteSiteCommandInput struct {
	Path struct {
		Id string
	}
}

//GetSiteCommand - Get a Site
//RequestType: GET
//Input: input *GetSiteCommandInput
func (s *SitesService) GetSiteCommand(input *GetSiteCommandInput) (result *SiteView, resp *http.Response, err error) {
	path := "/sites/{id}"

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

type GetSiteCommandInput struct {
	Path struct {
		Id string
	}
}

//UpdateSiteCommand - Update a Site
//RequestType: PUT
//Input: input *UpdateSiteCommandInput
func (s *SitesService) UpdateSiteCommand(input *UpdateSiteCommandInput) (result *SiteView, resp *http.Response, err error) {
	path := "/sites/{id}"

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

type UpdateSiteCommandInput struct {
	Body SiteView
	Path struct {
		Id string
	}
}
