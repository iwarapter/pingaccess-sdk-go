package authenticationChallengePolicies

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
)

type AuthenticationChallengePoliciesAPI interface {
	GetAuthenticationChallengePoliciesCommand(input *GetAuthenticationChallengePoliciesCommandInput) (output *models.AuthenticationChallengePoliciesView, resp *http.Response, err error)
	AddAuthenticationChallengePolicyCommand(input *AddAuthenticationChallengePolicyCommandInput) (output *models.AuthenticationChallengePolicyView, resp *http.Response, err error)
	GetRequestMatcherDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error)
	GetRequestMatcherDescriptorCommand(input *GetRequestMatcherDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error)
	GetChallengeResponseFilterDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error)
	GetChallengeResponseFilterDescriptorCommand(input *GetChallengeResponseFilterDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error)
	GetChallengeResponseGeneratorDescriptorsCommand() (output *models.DescriptorsView, resp *http.Response, err error)
	GetChallengeResponseGeneratorDescriptorCommand(input *GetChallengeResponseGeneratorDescriptorCommandInput) (output *models.DescriptorView, resp *http.Response, err error)
	DeleteAuthenticationChallengePolicyCommand(input *DeleteAuthenticationChallengePolicyCommandInput) (resp *http.Response, err error)
	GetAuthenticationChallengePolicyCommand(input *GetAuthenticationChallengePolicyCommandInput) (output *models.AuthenticationChallengePolicyView, resp *http.Response, err error)
	UpdateAuthenticationChallengePolicyCommand(input *UpdateAuthenticationChallengePolicyCommandInput) (output *models.AuthenticationChallengePolicyView, resp *http.Response, err error)
}
