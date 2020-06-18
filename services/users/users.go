package users

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
	ServiceName = "Users"
)

type UsersService struct {
	*client.Client
}

func New(cfg *config.Config) *UsersService {

	return &UsersService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Users operation
func (c *UsersService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetUsersCommand - Get all Users
//RequestType: GET
//Input: input *GetUsersCommandInput
func (s *UsersService) GetUsersCommand(input *GetUsersCommandInput) (output *models.UsersView, resp *http.Response, err error) {
	path := "/users"
	op := &request.Operation{
		Name:       "GetUsersCommand",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"filter":        input.Filter,
			"username":      input.Username,
			"sortKey":       input.SortKey,
			"order":         input.Order,
		},
	}
	output = &models.UsersView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetUsersCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Username      string
	SortKey       string
	Order         string
}

//GetUserCommand - Get a User
//RequestType: GET
//Input: input *GetUserCommandInput
func (s *UsersService) GetUserCommand(input *GetUserCommandInput) (output *models.UserView, resp *http.Response, err error) {
	path := "/users/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetUserCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.UserView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetUserCommandInput struct {
	Id string
}

//UpdateUserCommand - Update a User
//RequestType: PUT
//Input: input *UpdateUserCommandInput
func (s *UsersService) UpdateUserCommand(input *UpdateUserCommandInput) (output *models.UserView, resp *http.Response, err error) {
	path := "/users/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateUserCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.UserView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateUserCommandInput struct {
	Body models.UserView
	Id   string
}

//UpdateUserPasswordCommand - Update a User's Password
//RequestType: PUT
//Input: input *UpdateUserPasswordCommandInput
func (s *UsersService) UpdateUserPasswordCommand(input *UpdateUserPasswordCommandInput) (output *models.UserPasswordView, resp *http.Response, err error) {
	path := "/users/{id}/password"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateUserPasswordCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.UserPasswordView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateUserPasswordCommandInput struct {
	Body models.UserPasswordView
	Id   string
}
