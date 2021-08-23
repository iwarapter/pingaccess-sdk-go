package authenticationChallengePolicies

import (
	"net/http"
	"strings"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/client"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/client/metadata"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/config"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "AuthenticationChallengePolicies"
)

//AuthenticationChallengePoliciesService provides the API operations for making requests to
// AuthenticationChallengePolicies endpoint.
type AuthenticationChallengePoliciesService struct {
	*client.Client
}

//New createa a new instance of the AuthenticationChallengePoliciesService client.
//
// Example:
//   cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL)
//
//   //Create a AuthenticationChallengePoliciesService from the configuration
//   svc := authenticationChallengePolicies.New(cfg)
//
func New(cfg *config.Config) *AuthenticationChallengePoliciesService {

	return &AuthenticationChallengePoliciesService{Client: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingaccess.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a AuthenticationChallengePolicies operation
func (s *AuthenticationChallengePoliciesService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := s.NewRequest(op, params, data)

	return req
}

//GetAuthenticationChallengePoliciesCommand - Get all Authentication Challenge Policies
//RequestType: GET
//Input: input *GetAuthenticationChallengePoliciesCommandInput
func (s *AuthenticationChallengePoliciesService) GetAuthenticationChallengePoliciesCommand(input *GetAuthenticationChallengePoliciesCommandInput) (output *models.AuthenticationChallengePoliciesView, resp *http.Response, err error) {
	path := "/authenticationChallengePolicies"
	op := &request.Operation{
		Name:       "GetAuthenticationChallengePoliciesCommand",
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
	output = &models.AuthenticationChallengePoliciesView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAuthenticationChallengePoliciesCommandInput - Inputs for GetAuthenticationChallengePoliciesCommand
type GetAuthenticationChallengePoliciesCommandInput struct {
	Page          string
	NumberPerPage string
	Filter        string
	Name          string
	SortKey       string
	Order         string
}

//AddAuthenticationChallengePolicyCommand - Create an Authentication Challenge Policy
//RequestType: POST
//Input: input *AddAuthenticationChallengePolicyCommandInput
func (s *AuthenticationChallengePoliciesService) AddAuthenticationChallengePolicyCommand(input *AddAuthenticationChallengePolicyCommandInput) (output *models.AuthenticationChallengePolicyView, resp *http.Response, err error) {
	path := "/authenticationChallengePolicies"
	op := &request.Operation{
		Name:        "AddAuthenticationChallengePolicyCommand",
		HTTPMethod:  "POST",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AuthenticationChallengePolicyView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// AddAuthenticationChallengePolicyCommandInput - Inputs for AddAuthenticationChallengePolicyCommand
type AddAuthenticationChallengePolicyCommandInput struct {
	Body models.AuthenticationChallengePolicyView
}

//GetRequestMatcherDescriptorsCommand - Get the descriptors for all the Authentication Challenge Policy Request Matchers
//RequestType: GET
//Input:
func (s *AuthenticationChallengePoliciesService) GetRequestMatcherDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error) {
	path := "/authenticationChallengePolicies/requestMatchers/descriptors"
	op := &request.Operation{
		Name:       "GetRequestMatcherDescriptorsCommand",
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

//GetRequestMatcherDescriptorCommand - Get the descriptor for an Authentication Challenge Policy Request Matcher type
//RequestType: GET
//Input: input *GetRequestMatcherDescriptorCommandInput
func (s *AuthenticationChallengePoliciesService) GetRequestMatcherDescriptorCommand(input *GetRequestMatcherDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error) {
	path := "/authenticationChallengePolicies/requestMatchers/descriptors/{requestMatcherType}"
	path = strings.Replace(path, "{requestMatcherType}", input.RequestMatcherType, -1)

	op := &request.Operation{
		Name:        "GetRequestMatcherDescriptorCommand",
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

// GetRequestMatcherDescriptorCommandInput - Inputs for GetRequestMatcherDescriptorCommand
type GetRequestMatcherDescriptorCommandInput struct {
	RequestMatcherType string
}

//GetChallengeResponseFilterDescriptorsCommand - Get the descriptors for all the Authentication Challenge Policy Response Filtersr
//RequestType: GET
//Input:
func (s *AuthenticationChallengePoliciesService) GetChallengeResponseFilterDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error) {
	path := "/authenticationChallengePolicies/responseFilters/descriptors"
	op := &request.Operation{
		Name:       "GetChallengeResponseFilterDescriptorsCommand",
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

//GetChallengeResponseFilterDescriptorCommand - Get the descriptor for an Authentication Challenge Policy Response Filter type
//RequestType: GET
//Input: input *GetChallengeResponseFilterDescriptorCommandInput
func (s *AuthenticationChallengePoliciesService) GetChallengeResponseFilterDescriptorCommand(input *GetChallengeResponseFilterDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error) {
	path := "/authenticationChallengePolicies/responseFilters/descriptors/{responseFilterType}"
	path = strings.Replace(path, "{responseFilterType}", input.ResponseFilterType, -1)

	op := &request.Operation{
		Name:        "GetChallengeResponseFilterDescriptorCommand",
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

// GetChallengeResponseFilterDescriptorCommandInput - Inputs for GetChallengeResponseFilterDescriptorCommand
type GetChallengeResponseFilterDescriptorCommandInput struct {
	ResponseFilterType string
}

//GetChallengeResponseGeneratorDescriptorsCommand - Get the descriptors for all the Authentication Challenge Policy Response Generators
//RequestType: GET
//Input:
func (s *AuthenticationChallengePoliciesService) GetChallengeResponseGeneratorDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error) {
	path := "/authenticationChallengePolicies/responseGenerators/descriptors"
	op := &request.Operation{
		Name:       "GetChallengeResponseGeneratorDescriptorsCommand",
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

//GetChallengeResponseGeneratorDescriptorCommand - Get the descriptor for an Authentication Challenge Policy Response Generator type
//RequestType: GET
//Input: input *GetChallengeResponseGeneratorDescriptorCommandInput
func (s *AuthenticationChallengePoliciesService) GetChallengeResponseGeneratorDescriptorCommand(input *GetChallengeResponseGeneratorDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error) {
	path := "/authenticationChallengePolicies/responseGenerators/descriptors/{responseGeneratorType}"
	path = strings.Replace(path, "{responseGeneratorType}", input.ResponseGeneratorType, -1)

	op := &request.Operation{
		Name:        "GetChallengeResponseGeneratorDescriptorCommand",
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

// GetChallengeResponseGeneratorDescriptorCommandInput - Inputs for GetChallengeResponseGeneratorDescriptorCommand
type GetChallengeResponseGeneratorDescriptorCommandInput struct {
	ResponseGeneratorType string
}

//DeleteAuthenticationChallengePolicyCommand - Delete an Authentication Challenge Policy
//RequestType: DELETE
//Input: input *DeleteAuthenticationChallengePolicyCommandInput
func (s *AuthenticationChallengePoliciesService) DeleteAuthenticationChallengePolicyCommand(input *DeleteAuthenticationChallengePolicyCommandInput) (resp *http.Response, err error) {
	path := "/authenticationChallengePolicies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "DeleteAuthenticationChallengePolicyCommand",
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

// DeleteAuthenticationChallengePolicyCommandInput - Inputs for DeleteAuthenticationChallengePolicyCommand
type DeleteAuthenticationChallengePolicyCommandInput struct {
	Id string
}

//GetAuthenticationChallengePolicyCommand - Get an Authentication Challenge Policy
//RequestType: GET
//Input: input *GetAuthenticationChallengePolicyCommandInput
func (s *AuthenticationChallengePoliciesService) GetAuthenticationChallengePolicyCommand(input *GetAuthenticationChallengePolicyCommandInput) (output *models.AuthenticationChallengePolicyView, resp *http.Response, err error) {
	path := "/authenticationChallengePolicies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "GetAuthenticationChallengePolicyCommand",
		HTTPMethod:  "GET",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AuthenticationChallengePolicyView{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// GetAuthenticationChallengePolicyCommandInput - Inputs for GetAuthenticationChallengePolicyCommand
type GetAuthenticationChallengePolicyCommandInput struct {
	Id string
}

//UpdateAuthenticationChallengePolicyCommand - Update an Authentication Challenge Policy
//RequestType: PUT
//Input: input *UpdateAuthenticationChallengePolicyCommandInput
func (s *AuthenticationChallengePoliciesService) UpdateAuthenticationChallengePolicyCommand(input *UpdateAuthenticationChallengePolicyCommandInput) (output *models.AuthenticationChallengePolicyView, resp *http.Response, err error) {
	path := "/authenticationChallengePolicies/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:        "UpdateAuthenticationChallengePolicyCommand",
		HTTPMethod:  "PUT",
		HTTPPath:    path,
		QueryParams: map[string]string{},
	}
	output = &models.AuthenticationChallengePolicyView{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

// UpdateAuthenticationChallengePolicyCommandInput - Inputs for UpdateAuthenticationChallengePolicyCommand
type UpdateAuthenticationChallengePolicyCommandInput struct {
	Body models.AuthenticationChallengePolicyView
	Id   string
}
