package pingaccess

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type WebSessionsService service

//GetWebSessionsCommand - Get all WebSessions
//RequestType: GET
//Input: input *GetWebSessionsCommandInput
func (s *WebSessionsService) GetWebSessionsCommand(input *GetWebSessionsCommandInput) (result *WebSessionsView, resp *http.Response, err error) {
	path := "/webSessions"

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

type GetWebSessionsCommandInput struct {
	Params struct {
		Page          int
		NumberPerPage int
		Filter        string
		Name          string
		SortKey       string
		Order         string
	}
}

//AddWebSessionCommand - Create a WebSession
//RequestType: POST
//Input: input *AddWebSessionCommandInput
func (s *WebSessionsService) AddWebSessionCommand(input *AddWebSessionCommandInput) (result *WebSessionView, resp *http.Response, err error) {
	path := "/webSessions"

	b, _ := json.Marshal(input.Body)
	log.Printf("[CLIENT] RequestBody: %s", b)
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

type AddWebSessionCommandInput struct {
	Body WebSessionView
}

//DeleteWebSessionCommand - Delete a WebSession
//RequestType: DELETE
//Input: input *DeleteWebSessionCommandInput
func (s *WebSessionsService) DeleteWebSessionCommand(input *DeleteWebSessionCommandInput) (resp *http.Response, err error) {
	path := "/webSessions/{id}"

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

type DeleteWebSessionCommandInput struct {
	Path struct {
		Id string
	}
}

//GetWebSessionCommand - Get a WebSession
//RequestType: GET
//Input: input *GetWebSessionCommandInput
func (s *WebSessionsService) GetWebSessionCommand(input *GetWebSessionCommandInput) (result *WebSessionView, resp *http.Response, err error) {
	path := "/webSessions/{id}"

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

type GetWebSessionCommandInput struct {
	Path struct {
		Id string
	}
}

//UpdateWebSessionCommand - Update a WebSession
//RequestType: PUT
//Input: input *UpdateWebSessionCommandInput
func (s *WebSessionsService) UpdateWebSessionCommand(input *UpdateWebSessionCommandInput) (result *WebSessionView, resp *http.Response, err error) {
	path := "/webSessions/{id}"

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

type UpdateWebSessionCommandInput struct {
	Body WebSessionView
	Path struct {
		Id string
	}
}
