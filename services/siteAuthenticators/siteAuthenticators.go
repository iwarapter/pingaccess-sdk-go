package siteAuthenticators

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
	ServiceName = "SiteAuthenticators"
)

type SiteAuthenticatorsService struct {
	*client.Client
}

func New(cfg *config.Config) *SiteAuthenticatorsService {

	return &SiteAuthenticatorsService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a SiteAuthenticators operation
func (c *SiteAuthenticatorsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetSiteAuthenticatorsCommand - Get all Site Authenticators
//RequestType: GET
//Input: input *GetSiteAuthenticatorsCommandInput
func (s *SiteAuthenticatorsService) GetSiteAuthenticatorsCommand(input *GetSiteAuthenticatorsCommandInput) (output *models.SiteAuthenticatorsView, resp *http.Response, err error) {
	path := "/siteAuthenticators"
	op := &request.Operation{
		Name:       "GetSiteAuthenticatorsCommand",
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
	output = &models.SiteAuthenticatorsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetSiteAuthenticatorsCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddSiteAuthenticatorCommand - Create a Site Authenticator
//RequestType: POST
//Input: input *AddSiteAuthenticatorCommandInput
func (s *SiteAuthenticatorsService) AddSiteAuthenticatorCommand(input *AddSiteAuthenticatorCommandInput) (output *models.SiteAuthenticatorView, resp *http.Response, err error) {
	path := "/siteAuthenticators"
	op := &request.Operation{
		Name:        "AddSiteAuthenticatorCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.SiteAuthenticatorView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type AddSiteAuthenticatorCommandInput struct {
	Body models.SiteAuthenticatorView
}

//GetSiteAuthenticatorDescriptorsCommand - Get descriptors for all supported Site Authenticator types
//RequestType: GET
//Input:
func (s *SiteAuthenticatorsService) GetSiteAuthenticatorDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error) {
	path := "/siteAuthenticators/descriptors"
	op := &request.Operation{
		Name:       "GetSiteAuthenticatorDescriptorsCommand",
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

//GetSiteAuthenticatorDescriptorCommand - Get descriptor for a Site Authenticator type
//RequestType: GET
//Input: input *GetSiteAuthenticatorDescriptorCommandInput
func (s *SiteAuthenticatorsService) GetSiteAuthenticatorDescriptorCommand(input *GetSiteAuthenticatorDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error) {
	path := "/siteAuthenticators/descriptors/{siteAuthenticatorType}"
	path = strings.Replace(path, "{siteAuthenticatorType}", input.SiteAuthenticatorType, -1)

	op := &request.Operation{
		Name:        "GetSiteAuthenticatorDescriptorCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.DescriptorView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetSiteAuthenticatorDescriptorCommandInput struct {
	SiteAuthenticatorType string
}

//DeleteSiteAuthenticatorCommand - Delete a Site Authenticator
//RequestType: DELETE
//Input: input *DeleteSiteAuthenticatorCommandInput
func (s *SiteAuthenticatorsService) DeleteSiteAuthenticatorCommand(input *DeleteSiteAuthenticatorCommandInput) (resp *http.Response, err error) {
	path := "/siteAuthenticators/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteSiteAuthenticatorCommand",
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

type DeleteSiteAuthenticatorCommandInput struct {
	Id string
}

//GetSiteAuthenticatorCommand - Get a Site Authenticator
//RequestType: GET
//Input: input *GetSiteAuthenticatorCommandInput
func (s *SiteAuthenticatorsService) GetSiteAuthenticatorCommand(input *GetSiteAuthenticatorCommandInput) (output *models.SiteAuthenticatorView, resp *http.Response, err error) {
	path := "/siteAuthenticators/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetSiteAuthenticatorCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.SiteAuthenticatorView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetSiteAuthenticatorCommandInput struct {
	Id string
}

//UpdateSiteAuthenticatorCommand - Update a Site Authenticator
//RequestType: PUT
//Input: input *UpdateSiteAuthenticatorCommandInput
func (s *SiteAuthenticatorsService) UpdateSiteAuthenticatorCommand(input *UpdateSiteAuthenticatorCommandInput) (output *models.SiteAuthenticatorView, resp *http.Response, err error) {
	path := "/siteAuthenticators/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateSiteAuthenticatorCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.SiteAuthenticatorView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateSiteAuthenticatorCommandInput struct {
	Body models.SiteAuthenticatorView
	Id   string
}
