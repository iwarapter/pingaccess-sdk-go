package environment

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/client"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/client/metadata"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/config"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "Environment"
)

//EnvironmentService provides the API operations for making requests to
// Environment endpoint.
type EnvironmentService struct {
	*client.Client
}

//New createa a new instance of the EnvironmentService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a EnvironmentService from the configuration
//   svc := environment.New(cfg)
//
func New(cfg *config.Config) *EnvironmentService {

	return &EnvironmentService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Environment operation
func (s *EnvironmentService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//DeleteEnvironmentCommand - Resets the Environment configuration to default values
//RequestType: DELETE
//Input:
func (s *EnvironmentService) DeleteEnvironmentCommand() (resp *http.Response, err error) {
	path := "/environment"
	op := &request.Operation{
		Name:       "DeleteEnvironmentCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//GetEnvironmentCommand - Get the Environment
//RequestType: GET
//Input:
func (s *EnvironmentService) GetEnvironmentCommand() (output *models.EnvironmentView, resp *http.Response, err error) {
	path := "/environment"
	op := &request.Operation{
		Name:       "GetEnvironmentCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.EnvironmentView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateEnvironmentCommand - Update the Environment
//RequestType: PUT
//Input: input *UpdateEnvironmentCommandInput
func (s *EnvironmentService) UpdateEnvironmentCommand(input *UpdateEnvironmentCommandInput) (output *models.EnvironmentView, resp *http.Response, err error) {
	path := "/environment"
	op := &request.Operation{
		Name:        "UpdateEnvironmentCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.EnvironmentView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateEnvironmentCommandInput - Inputs for UpdateEnvironmentCommand
type UpdateEnvironmentCommandInput struct {
	Body models.EnvironmentView
}
