package httpsListeners

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
	ServiceName = "HttpsListeners"
)

type HttpsListenersService struct {
	*client.Client
}

func New(cfg *config.Config) *HttpsListenersService {

	return &HttpsListenersService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a HttpsListeners operation
func (c *HttpsListenersService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetHttpsListenersCommand - Get all HTTPS Listeners
//RequestType: GET
//Input: input *GetHttpsListenersCommandInput
func (s *HttpsListenersService) GetHttpsListenersCommand(input *GetHttpsListenersCommandInput) (output *models.HttpsListenersView, resp *http.Response, err error) {
	path := "/httpsListeners"
	op := &request.Operation{
		Name:       "GetHttpsListenersCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"sortKey": input.SortKey,
			"order":   input.Order,
		},
	}
	output = &models.HttpsListenersView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetHttpsListenersCommandInput struct {
	SortKey string
	Order   string
}

//GetHttpsListenerCommand - Get an HTTPS Listener
//RequestType: GET
//Input: input *GetHttpsListenerCommandInput
func (s *HttpsListenersService) GetHttpsListenerCommand(input *GetHttpsListenerCommandInput) (output *models.HttpsListenerView, resp *http.Response, err error) {
	path := "/httpsListeners/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetHttpsListenerCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.HttpsListenerView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetHttpsListenerCommandInput struct {
	Id string
}

//UpdateHttpsListener - Update an HTTPS Listener
//RequestType: PUT
//Input: input *UpdateHttpsListenerInput
func (s *HttpsListenersService) UpdateHttpsListener(input *UpdateHttpsListenerInput) (output *models.HttpsListenerView, resp *http.Response, err error) {
	path := "/httpsListeners/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateHttpsListener",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.HttpsListenerView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateHttpsListenerInput struct {
	Body models.HttpsListenerView
	Id   string
}
