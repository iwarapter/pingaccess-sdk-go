package config

import (
	"net/http"
	"strings"

	"github.com/iwarapter/pingaccess-sdk-go/v60/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/v60/pingaccess/client"
	"github.com/iwarapter/pingaccess-sdk-go/v60/pingaccess/client/metadata"
	"github.com/iwarapter/pingaccess-sdk-go/v60/pingaccess/config"
	"github.com/iwarapter/pingaccess-sdk-go/v60/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v60/pingaccess/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "Config"
)

//ConfigService provides the API operations for making requests to
// Config endpoint.
type ConfigService struct {
	*client.Client
}

//New creates a new instance of the ConfigService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint(paURL.String())
//
//   //Create a ConfigService from the configuration
//   svc := config.New(cfg)
//
func New(cfg *config.Config) *ConfigService {

	return &ConfigService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Config operation
func (s *ConfigService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//ConfigExportCommand - [Attention: The endpoint "/config/export" is deprecated. The "config/export/workflows" endpoint should be used instead.]    Export a JSON backup of the entire system
//RequestType: GET
//Input:
func (s *ConfigService) ConfigExportCommand() (output *models.ExportData, resp *http.Response, err error) {
	path := "/config/export"
	op := &request.Operation{
		Name:       "ConfigExportCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ExportData{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetConfigExportWorkflowsCommand - Get the status of pending Exports
//RequestType: GET
//Input:
func (s *ConfigService) GetConfigExportWorkflowsCommand() (output *models.ConfigStatusesView, resp *http.Response, err error) {
	path := "/config/export/workflows"
	op := &request.Operation{
		Name:       "GetConfigExportWorkflowsCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ConfigStatusesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//AddConfigExportWorkflowCommand - Start a JSON backup of the entire system for export
//RequestType: POST
//Input:
func (s *ConfigService) AddConfigExportWorkflowCommand() (output *models.ConfigStatusView, resp *http.Response, err error) {
	path := "/config/export/workflows"
	op := &request.Operation{
		Name:       "AddConfigExportWorkflowCommand",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.ConfigStatusView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetConfigExportWorkflowCommand - Check the status of an Export
//RequestType: GET
//Input: input *GetConfigExportWorkflowCommandInput
func (s *ConfigService) GetConfigExportWorkflowCommand(input *GetConfigExportWorkflowCommandInput) (output *models.ConfigStatusView, resp *http.Response, err error) {
	path := "/config/export/workflows/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetConfigExportWorkflowCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ConfigStatusView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetConfigExportWorkflowCommandInput - Inputs for GetConfigExportWorkflowCommand
type GetConfigExportWorkflowCommandInput struct {
	Id string
}

//GetConfigExportWorkflowDataCommand - Export a JSON backup of the entire system
//RequestType: GET
//Input: input *GetConfigExportWorkflowDataCommandInput
func (s *ConfigService) GetConfigExportWorkflowDataCommand(input *GetConfigExportWorkflowDataCommandInput) (output *models.ExportData, resp *http.Response, err error) {
	path := "/config/export/workflows/{id}/data"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetConfigExportWorkflowDataCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ExportData{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetConfigExportWorkflowDataCommandInput - Inputs for GetConfigExportWorkflowDataCommand
type GetConfigExportWorkflowDataCommandInput struct {
	Id string
}

//ConfigImportCommand - [Attention: The endpoint "/config/import" is deprecated. The "config/import/workflows" endpoint should be used instead.]   Import JSON backup.
//RequestType: POST
//Input: input *ConfigImportCommandInput
func (s *ConfigService) ConfigImportCommand(input *ConfigImportCommandInput) (resp *http.Response, err error) {
	path := "/config/import"
	op := &request.Operation{
		Name:        "ConfigImportCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}

	req := s.newRequest(op, input.Body, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

// ConfigImportCommandInput - Inputs for ConfigImportCommand
type ConfigImportCommandInput struct {
	Body string
}

//GetConfigImportWorkflowsCommand - Get the status of pending imports
//RequestType: GET
//Input:
func (s *ConfigService) GetConfigImportWorkflowsCommand() (output *models.ConfigStatusesView, resp *http.Response, err error) {
	path := "/config/import/workflows"
	op := &request.Operation{
		Name:       "GetConfigImportWorkflowsCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ConfigStatusesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//AddConfigImportWorkflowCommand - Start an Import of a JSON backup
//RequestType: POST
//Input: input *AddConfigImportWorkflowCommandInput
func (s *ConfigService) AddConfigImportWorkflowCommand(input *AddConfigImportWorkflowCommandInput) (resp *http.Response, err error) {
	path := "/config/import/workflows"
	op := &request.Operation{
		Name:        "AddConfigImportWorkflowCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}

	req := s.newRequest(op, input.Body, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

// AddConfigImportWorkflowCommandInput - Inputs for AddConfigImportWorkflowCommand
type AddConfigImportWorkflowCommandInput struct {
	Body models.ExportData
}

//GetConfigImportWorkflowCommand - Check the status of an Export
//RequestType: GET
//Input: input *GetConfigImportWorkflowCommandInput
func (s *ConfigService) GetConfigImportWorkflowCommand(input *GetConfigImportWorkflowCommandInput) (output *models.ConfigStatusView, resp *http.Response, err error) {
	path := "/config/import/workflows/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetConfigImportWorkflowCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.ConfigStatusView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetConfigImportWorkflowCommandInput - Inputs for GetConfigImportWorkflowCommand
type GetConfigImportWorkflowCommandInput struct {
	Id string
}
