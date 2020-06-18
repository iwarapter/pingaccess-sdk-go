package adminSessionInfo

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
	ServiceName = "AdminSessionInfo"
)

type AdminSessionInfoService struct {
	*client.Client
}

func New(cfg *config.Config) *AdminSessionInfoService {

	return &AdminSessionInfoService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a AdminSessionInfo operation
func (c *AdminSessionInfoService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//AdminSessionDeleteCommand - Invalidate the Admin session information
//RequestType: DELETE
//Input:
func (s *AdminSessionInfoService) AdminSessionDeleteCommand() (resp *http.Response, err error) {
	path := "/adminSessionInfo"
	op := &request.Operation{
		Name:       "AdminSessionDeleteCommand",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

//AdminSessionInfoCommand - Return the Admin session information
//RequestType: GET
//Input:
func (s *AdminSessionInfoService) AdminSessionInfoCommand() (output *models.SessionInfo, resp *http.Response, err error) {
	path := "/adminSessionInfo"
	op := &request.Operation{
		Name:       "AdminSessionInfoCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SessionInfo{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//AdminSessionInfoCheckCommand - Return the Admin session information without affecting session expiration
//RequestType: GET
//Input:
func (s *AdminSessionInfoService) AdminSessionInfoCheckCommand() (output *models.SessionInfo, resp *http.Response, err error) {
	path := "/adminSessionInfo/checkOnly"
	op := &request.Operation{
		Name:       "AdminSessionInfoCheckCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SessionInfo{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}
