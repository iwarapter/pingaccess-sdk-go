package thirdPartyServices

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
	ServiceName = "ThirdPartyServices"
)

type ThirdPartyServicesService struct {
	*client.Client
}

func New(cfg *config.Config) *ThirdPartyServicesService {

	return &ThirdPartyServicesService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a ThirdPartyServices operation
func (c *ThirdPartyServicesService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetThirdPartyServicesCommand - Get all Third-Party Services
//RequestType: GET
//Input: input *GetThirdPartyServicesCommandInput
func (s *ThirdPartyServicesService) GetThirdPartyServicesCommand(input *GetThirdPartyServicesCommandInput) (output *models.ThirdPartyServicesView, resp *http.Response, err error) {
	path := "/thirdPartyServices"
	op := &request.Operation{
		Name:       "GetThirdPartyServicesCommand",
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
	output = &models.ThirdPartyServicesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetThirdPartyServicesCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddThirdPartyServiceCommand - Create a Third-Party Service
//RequestType: POST
//Input: input *AddThirdPartyServiceCommandInput
func (s *ThirdPartyServicesService) AddThirdPartyServiceCommand(input *AddThirdPartyServiceCommandInput) (output *models.ThirdPartyServiceView, resp *http.Response, err error) {
	path := "/thirdPartyServices"
	op := &request.Operation{
		Name:        "AddThirdPartyServiceCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ThirdPartyServiceView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type AddThirdPartyServiceCommandInput struct {
	Body models.ThirdPartyServiceView
}

//DeleteThirdPartyServiceCommand - Delete a Third-Party Service
//RequestType: DELETE
//Input: input *DeleteThirdPartyServiceCommandInput
func (s *ThirdPartyServicesService) DeleteThirdPartyServiceCommand(input *DeleteThirdPartyServiceCommandInput) (resp *http.Response, err error) {
	path := "/thirdPartyServices/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteThirdPartyServiceCommand",
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

type DeleteThirdPartyServiceCommandInput struct {
	Id string
}

//GetThirdPartyServiceCommand - Get a Third-Party Service
//RequestType: GET
//Input: input *GetThirdPartyServiceCommandInput
func (s *ThirdPartyServicesService) GetThirdPartyServiceCommand(input *GetThirdPartyServiceCommandInput) (output *models.ThirdPartyServiceView, resp *http.Response, err error) {
	path := "/thirdPartyServices/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetThirdPartyServiceCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ThirdPartyServiceView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetThirdPartyServiceCommandInput struct {
	Id string
}

//UpdateThirdPartyServiceCommand - Update a Third-Party Service
//RequestType: PUT
//Input: input *UpdateThirdPartyServiceCommandInput
func (s *ThirdPartyServicesService) UpdateThirdPartyServiceCommand(input *UpdateThirdPartyServiceCommandInput) (output *models.ThirdPartyServiceView, resp *http.Response, err error) {
	path := "/thirdPartyServices/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateThirdPartyServiceCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ThirdPartyServiceView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateThirdPartyServiceCommandInput struct {
	Body models.ThirdPartyServiceView
	Id   string
}
