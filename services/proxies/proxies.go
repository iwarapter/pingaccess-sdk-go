package proxies

import (
	"net/http"
	"strings"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/client"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/client/metadata"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/config"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "Proxies"
)

type ProxiesService struct {
	*client.Client
}

func New(cfg *config.Config) *ProxiesService {

	return &ProxiesService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Proxies operation
func (c *ProxiesService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetProxiesCommand - Get all Proxies
//RequestType: GET
//Input: input *GetProxiesCommandInput
func (s *ProxiesService) GetProxiesCommand(input *GetProxiesCommandInput) (output *models.HttpClientProxyView, resp *http.Response, err error) {
	path := "/proxies"
	op := &request.Operation{
		Name:       "GetProxiesCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"filter":        input.Filter,
			"name":          input.Name,
			"sortKey":       input.SortKey,
			"order":         input.Order,
		},
	}
	output = &models.HttpClientProxyView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetProxiesCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddProxyCommand - Create a Proxy
//RequestType: POST
//Input: input *AddProxyCommandInput
func (s *ProxiesService) AddProxyCommand(input *AddProxyCommandInput) (output *models.HttpClientProxyView, resp *http.Response, err error) {
	path := "/proxies"
	op := &request.Operation{
		Name:        "AddProxyCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.HttpClientProxyView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type AddProxyCommandInput struct {
	Body models.HttpClientProxyView
}

//DeleteProxyCommand - Delete a Proxy
//RequestType: DELETE
//Input: input *DeleteProxyCommandInput
func (s *ProxiesService) DeleteProxyCommand(input *DeleteProxyCommandInput) (resp *http.Response, err error) {
	path := "/proxies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteProxyCommand",
		HTTPMethod:  "DELETE",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

type DeleteProxyCommandInput struct {
	Id string
}

//GetProxyCommand - Get a Proxy
//RequestType: GET
//Input: input *GetProxyCommandInput
func (s *ProxiesService) GetProxyCommand(input *GetProxyCommandInput) (output *models.HttpClientProxyView, resp *http.Response, err error) {
	path := "/proxies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetProxyCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.HttpClientProxyView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetProxyCommandInput struct {
	Id string
}

//UpdateProxyCommand - Update a Proxy
//RequestType: PUT
//Input: input *UpdateProxyCommandInput
func (s *ProxiesService) UpdateProxyCommand(input *UpdateProxyCommandInput) (output *models.HttpClientProxyView, resp *http.Response, err error) {
	path := "/proxies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateProxyCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.HttpClientProxyView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateProxyCommandInput struct {
	Body models.HttpClientProxyView
	Id   string
}
