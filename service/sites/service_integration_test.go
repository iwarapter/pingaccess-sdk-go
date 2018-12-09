// +build integration

package sites

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

func TestGetSitesCommand(t *testing.T) {
	svc := New(config())

	input := GetSitesCommandInput{}
	results, _ := svc.GetSitesCommand(&input)
	if len(results.Items) < 0 {
		t.Errorf("Marshelled object should contain items")
	}
}

func TestAddSiteCommand(t *testing.T) {
	svc := New(config())

	targets := []*string{}
	str := "localhost:1234"
	targets = append(targets, &str)
	input := AddSiteCommandInput{
		Body: SiteView{
			Name:                    "bar",
			Targets:                 targets,
			MaxConnections:          -1,
			MaxWebSocketConnections: -1,
		},
	}
	_, err := svc.AddSiteCommand(&input)
	if err != nil {
		t.Errorf("Unable to created execute command: %s", err.Error())
	}
}
