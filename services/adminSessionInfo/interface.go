package adminSessionInfo

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v60/pingaccess/models"
)

type AdminSessionInfoAPI interface {
	AdminSessionDeleteCommand() (resp *http.Response, err error)
	AdminSessionInfoCommand() (output *models.SessionInfo, resp *http.Response, err error)
	AdminSessionInfoCheckCommand() (output *models.SessionInfo, resp *http.Response, err error)
}
