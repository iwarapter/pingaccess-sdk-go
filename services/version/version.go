package version

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
	ServiceName = "Version"
)

type VersionService struct {
	*client.Client
}

func New(cfg *config.Config) *VersionService {

	return &VersionService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Version operation
func (c *VersionService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//VersionCommand - Get the PingAccess version number
//RequestType: GET
//Input:
func (s *VersionService) VersionCommand() (output *models.VersionDocClass, resp *http.Response, err error) {
	path := "/version"
	op := &request.Operation{
		Name:       "VersionCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.VersionDocClass{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}
