// +build integration

package virtualhosts

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

func TestGetVirtualHostsCommand(t *testing.T) {
	svc := New(config())

	input := GetVirtualHostsCommandInput{}
	results, _ := svc.GetVirtualHostsCommand(&input)
	if len(results.Items) < 0 {
		t.Errorf("Marshelled object should contain items")
	}
}

func TestAddVirtualHostCommand(t *testing.T) {
	svc := New(config())

	input := AddVirtualHostCommandInput{
		Body: VirtualHostView{
			AgentResourceCacheTTL: 900,
			Host:      "localhost",
			KeyPairId: 0,
			Port:      3000,
			TrustedCertificateGroupId: 0,
		}}
	_, err := svc.AddVirtualHostCommand(&input)
	if err != nil {
		t.Errorf("Unable to created execute command: %s", err.Error())
	}
}
