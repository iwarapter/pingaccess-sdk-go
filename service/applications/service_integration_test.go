// +build integration

package applications

import (
	"crypto/tls"
	"net/http"
	"os"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func config() *pingaccess.Config {
	return &pingaccess.Config{
		Username: "Administrator",
		Password: "2Access2",
		BaseURL:  "https://localhost:9000/pa-admin-api/v3",
	}
}

func TestMain(m *testing.M) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	retCode := m.Run()
	os.Exit(retCode)
}

func TestGetApplicationsCommand(t *testing.T) {
	svc := New(config())

	input := GetApplicationsCommandInput{}
	results, _ := svc.GetApplicationsCommand(&input)
	if len(results.Items) != 1 {
		t.Errorf("Marshelled object should contain items")
	}
}
