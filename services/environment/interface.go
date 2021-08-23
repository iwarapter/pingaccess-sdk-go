package environment

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

type EnvironmentAPI interface {
	DeleteEnvironmentCommand() (resp *http.Response, err error)
	GetEnvironmentCommand() (output *models.EnvironmentView, resp *http.Response, err error)
	UpdateEnvironmentCommand(input *UpdateEnvironmentCommandInput) (output *models.EnvironmentView, resp *http.Response, err error)
}
