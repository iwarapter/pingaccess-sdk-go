package applications

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
	ServiceName = "Applications"
)

type ApplicationsService struct {
	*client.Client
}

func New(cfg *config.Config) *ApplicationsService {

	return &ApplicationsService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Applications operation
func (c *ApplicationsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetApplicationsCommand - Get all Applications
//RequestType: GET
//Input: input *GetApplicationsCommandInput
func (s *ApplicationsService) GetApplicationsCommand(input *GetApplicationsCommandInput) (output *models.ApplicationsView, resp *http.Response, err error) {
	path := "/applications"
	op := &request.Operation{
		Name:       "GetApplicationsCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"page":          input.Page,
			"siteId":        input.SiteId,
			"numberPerPage": input.NumberPerPage,
			"agentId":       input.AgentId,
			"virtualHostId": input.VirtualHostId,
			"ruleId":        input.RuleId,
			"rulesetId":     input.RulesetId,
			"filter":        input.Filter,
			"name":          input.Name,
			"sortKey":       input.SortKey,
			"order":         input.Order,
		},
	}
	output = &models.ApplicationsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetApplicationsCommandInput struct {
	Page          string
	SiteId        string
	NumberPerPage string
	AgentId       string
	VirtualHostId string
	RuleId        string
	RulesetId     string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddApplicationCommand - Add an Application
//RequestType: POST
//Input: input *AddApplicationCommandInput
func (s *ApplicationsService) AddApplicationCommand(input *AddApplicationCommandInput) (output *models.ApplicationView, resp *http.Response, err error) {
	path := "/applications"
	op := &request.Operation{
		Name:        "AddApplicationCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ApplicationView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type AddApplicationCommandInput struct {
	Body models.ApplicationView
}

//DeleteReservedApplicationCommand - Resets the Reserved Application configuration to default values
//RequestType: DELETE
//Input:
func (s *ApplicationsService) DeleteReservedApplicationCommand() (resp *http.Response, err error) {
	path := "/applications/reserved"
	op := &request.Operation{
		Name:       "DeleteReservedApplicationCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetReservedApplicationCommand - Get Reserved Application configuration
//RequestType: GET
//Input:
func (s *ApplicationsService) GetReservedApplicationCommand() (output *models.ReservedApplicationView, resp *http.Response, err error) {
	path := "/applications/reserved"
	op := &request.Operation{
		Name:       "GetReservedApplicationCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ReservedApplicationView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateReservedApplicationCommand - Update Reserved Application configuration
//RequestType: PUT
//Input: input *UpdateReservedApplicationCommandInput
func (s *ApplicationsService) UpdateReservedApplicationCommand(input *UpdateReservedApplicationCommandInput) (output *models.ReservedApplicationView, resp *http.Response, err error) {
	path := "/applications/reserved"
	op := &request.Operation{
		Name:        "UpdateReservedApplicationCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ReservedApplicationView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateReservedApplicationCommandInput struct {
	Body models.ReservedApplicationView
}

//GetResourcesCommand - Get all Resources
//RequestType: GET
//Input: input *GetResourcesCommandInput
func (s *ApplicationsService) GetResourcesCommand(input *GetResourcesCommandInput) (output *models.ResourcesView, resp *http.Response, err error) {
	path := "/applications/resources"
	op := &request.Operation{
		Name:       "GetResourcesCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"ruleId":        input.RuleId,
			"rulesetId":     input.RulesetId,
			"name":          input.Name,
			"filter":        input.Filter,
			"sortKey":       input.SortKey,
			"order":         input.Order,
		},
	}
	output = &models.ResourcesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetResourcesCommandInput struct {
	Page          string
	NumberPerPage string
	RuleId        string
	RulesetId     string
	Name          string
	Filter        string
	SortKey       string
	Order         string
}

//GetApplicationsResourcesMethodsCommand - Get Application Resource Methods
//RequestType: GET
//Input:
func (s *ApplicationsService) GetApplicationsResourcesMethodsCommand() (output *models.MethodsView, resp *http.Response, err error) {
	path := "/applications/resources/methods"
	op := &request.Operation{
		Name:       "GetApplicationsResourcesMethodsCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.MethodsView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteApplicationResourceCommand - Delete an Application Resource
//RequestType: DELETE
//Input: input *DeleteApplicationResourceCommandInput
func (s *ApplicationsService) DeleteApplicationResourceCommand(input *DeleteApplicationResourceCommandInput) (resp *http.Response, err error) {
	path := "/applications/{applicationId}/resources/{resourceId}"
	path = strings.Replace(path, "{applicationId}", input.ApplicationId, -1)

	path = strings.Replace(path, "{resourceId}", input.ResourceId, -1)

	op := &request.Operation{
		Name:        "DeleteApplicationResourceCommand",
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

type DeleteApplicationResourceCommandInput struct {
	ApplicationId string
	ResourceId    string
}

//GetApplicationResourceCommand - Get an Application Resource
//RequestType: GET
//Input: input *GetApplicationResourceCommandInput
func (s *ApplicationsService) GetApplicationResourceCommand(input *GetApplicationResourceCommandInput) (output *models.ResourceView, resp *http.Response, err error) {
	path := "/applications/{applicationId}/resources/{resourceId}"
	path = strings.Replace(path, "{applicationId}", input.ApplicationId, -1)

	path = strings.Replace(path, "{resourceId}", input.ResourceId, -1)

	op := &request.Operation{
		Name:        "GetApplicationResourceCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ResourceView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetApplicationResourceCommandInput struct {
	ApplicationId string
	ResourceId    string
}

//UpdateApplicationResourceCommand - Update an Application Resource
//RequestType: PUT
//Input: input *UpdateApplicationResourceCommandInput
func (s *ApplicationsService) UpdateApplicationResourceCommand(input *UpdateApplicationResourceCommandInput) (output *models.ResourceView, resp *http.Response, err error) {
	path := "/applications/{applicationId}/resources/{resourceId}"
	path = strings.Replace(path, "{applicationId}", input.ApplicationId, -1)

	path = strings.Replace(path, "{resourceId}", input.ResourceId, -1)

	op := &request.Operation{
		Name:        "UpdateApplicationResourceCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ResourceView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateApplicationResourceCommandInput struct {
	Body          models.ResourceView
	ApplicationId string
	ResourceId    string
}

//DeleteApplicationCommand - Delete an Application
//RequestType: DELETE
//Input: input *DeleteApplicationCommandInput
func (s *ApplicationsService) DeleteApplicationCommand(input *DeleteApplicationCommandInput) (output *models.ApplicationView, resp *http.Response, err error) {
	path := "/applications/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteApplicationCommand",
		HTTPMethod:  "DELETE",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ApplicationView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type DeleteApplicationCommandInput struct {
	Id string
}

//GetApplicationCommand - Get an Application
//RequestType: GET
//Input: input *GetApplicationCommandInput
func (s *ApplicationsService) GetApplicationCommand(input *GetApplicationCommandInput) (output *models.ApplicationView, resp *http.Response, err error) {
	path := "/applications/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetApplicationCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ApplicationView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetApplicationCommandInput struct {
	Id string
}

//UpdateApplicationCommand - Update an Application
//RequestType: PUT
//Input: input *UpdateApplicationCommandInput
func (s *ApplicationsService) UpdateApplicationCommand(input *UpdateApplicationCommandInput) (output *models.ApplicationView, resp *http.Response, err error) {
	path := "/applications/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateApplicationCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ApplicationView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateApplicationCommandInput struct {
	Body models.ApplicationView
	Id   string
}

//GetResourceMatchingEvaluationOrderCommand - Get the resource path ordering for an Application
//RequestType: GET
//Input: input *GetResourceMatchingEvaluationOrderCommandInput
func (s *ApplicationsService) GetResourceMatchingEvaluationOrderCommand(input *GetResourceMatchingEvaluationOrderCommandInput) (output *models.ResourceMatchingEvaluationOrderView, resp *http.Response, err error) {
	path := "/applications/{id}/resourceMatchingEvaluationOrder"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetResourceMatchingEvaluationOrderCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ResourceMatchingEvaluationOrderView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetResourceMatchingEvaluationOrderCommandInput struct {
	Id string
}

//GetApplicationResourcesCommand - Get Resources for an Application
//RequestType: GET
//Input: input *GetApplicationResourcesCommandInput
func (s *ApplicationsService) GetApplicationResourcesCommand(input *GetApplicationResourcesCommandInput) (output *models.ResourcesView, resp *http.Response, err error) {
	path := "/applications/{id}/resources"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetApplicationResourcesCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"name":          input.Name,
			"filter":        input.Filter,
			"sortKey":       input.SortKey,
			"order":         input.Order,
		},
	}
	output = &models.ResourcesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetApplicationResourcesCommandInput struct {
	Page          string
	NumberPerPage string
	Name          string
	Filter        string
	SortKey       string
	Order         string

	Id string
}

//AddApplicationResourceCommand - Add Resource to an Application
//RequestType: POST
//Input: input *AddApplicationResourceCommandInput
func (s *ApplicationsService) AddApplicationResourceCommand(input *AddApplicationResourceCommandInput) (output *models.ResourceView, resp *http.Response, err error) {
	path := "/applications/{id}/resources"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "AddApplicationResourceCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ResourceView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type AddApplicationResourceCommandInput struct {
	Body models.ResourceView
	Id   string
}

//GetResourceAutoOrderCommand - Computes an automatic, intelligent resource ordering for an Application.
//RequestType: GET
//Input: input *GetResourceAutoOrderCommandInput
func (s *ApplicationsService) GetResourceAutoOrderCommand(input *GetResourceAutoOrderCommandInput) (output *models.ResourceOrderView, resp *http.Response, err error) {
	path := "/applications/{id}/resources/autoOrder"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetResourceAutoOrderCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ResourceOrderView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetResourceAutoOrderCommandInput struct {
	Id string
}
