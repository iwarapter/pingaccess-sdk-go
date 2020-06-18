package hsmProviders

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
	ServiceName = "HsmProviders"
)

type HsmProvidersService struct {
	*client.Client
}

func New(cfg *config.Config) *HsmProvidersService {

	return &HsmProvidersService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a HsmProviders operation
func (c *HsmProvidersService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetHsmProvidersCommand - Get all HSM Providers
//RequestType: GET
//Input: input *GetHsmProvidersCommandInput
func (s *HsmProvidersService) GetHsmProvidersCommand(input *GetHsmProvidersCommandInput) (output *models.HsmProviderView, resp *http.Response, err error) {
	path := "/hsmProviders"
	op := &request.Operation{
		Name:       "GetHsmProvidersCommand",
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
	output = &models.HsmProviderView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetHsmProvidersCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddHsmProviderCommand - Create an HSM Provider
//RequestType: POST
//Input: input *AddHsmProviderCommandInput
func (s *HsmProvidersService) AddHsmProviderCommand(input *AddHsmProviderCommandInput) (output *models.HsmProviderView, resp *http.Response, err error) {
	path := "/hsmProviders"
	op := &request.Operation{
		Name:        "AddHsmProviderCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.HsmProviderView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type AddHsmProviderCommandInput struct {
	Body models.HsmProviderView
}

//GetHsmProviderDescriptorsCommand - Get descriptors for all HSM Providers
//RequestType: GET
//Input:
func (s *HsmProvidersService) GetHsmProviderDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error) {
	path := "/hsmProviders/descriptors"
	op := &request.Operation{
		Name:       "GetHsmProviderDescriptorsCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.DescriptorsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteHsmProviderCommand - Delete an HSM Provider
//RequestType: DELETE
//Input: input *DeleteHsmProviderCommandInput
func (s *HsmProvidersService) DeleteHsmProviderCommand(input *DeleteHsmProviderCommandInput) (resp *http.Response, err error) {
	path := "/hsmProviders/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteHsmProviderCommand",
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

type DeleteHsmProviderCommandInput struct {
	Id string
}

//GetHsmProviderCommand - Get an HSM Provider
//RequestType: GET
//Input: input *GetHsmProviderCommandInput
func (s *HsmProvidersService) GetHsmProviderCommand(input *GetHsmProviderCommandInput) (output *models.HsmProviderView, resp *http.Response, err error) {
	path := "/hsmProviders/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetHsmProviderCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.HsmProviderView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetHsmProviderCommandInput struct {
	Id string
}

//UpdateHsmProviderCommand - Update an HSM Provider
//RequestType: PUT
//Input: input *UpdateHsmProviderCommandInput
func (s *HsmProvidersService) UpdateHsmProviderCommand(input *UpdateHsmProviderCommandInput) (output *models.HsmProviderView, resp *http.Response, err error) {
	path := "/hsmProviders/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateHsmProviderCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.HsmProviderView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateHsmProviderCommandInput struct {
	Body models.HsmProviderView
	Id   string
}
