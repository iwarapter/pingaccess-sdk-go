package license

import (
	"net/http"

	"github.com/iwarapter/pingaccess-sdk-go/v60/pingaccess/models"
)

type LicenseAPI interface {
	GetLicenseCommand() (output *models.LicenseView, resp *http.Response, err error)
	ImportLicenseCommand(input *ImportLicenseCommandInput) (output *models.LicenseView, resp *http.Response, err error)
}
