package siteAuthenticators

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

func TestGetSiteAuthenticatorsCommand(t *testing.T) {
	svc := New(config())

	input := GetSiteAuthenticatorsCommandInput{}
	results, _ := svc.GetSiteAuthenticatorsCommand(&input)
	if len(results.Items) != 0 {
		t.Errorf("Marshelled object should not contain items")
	}
}
