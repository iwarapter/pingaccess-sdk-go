package applications

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
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
	if len(results.Items) == 0 {
		t.Errorf("Marshelled object should contain items")
	}
}

func TestWeCanMarshallApplications(t *testing.T) {
	results := ApplicationsView{}
	dat, err := ioutil.ReadFile("test_data/example_applications.json")
	if err != nil {
		t.Errorf("Unable to read test file")
		t.FailNow()
	}
	err = json.Unmarshal(dat, &results)
	if err != nil {
		t.Errorf("Unable to marshall example response")
	}

	if len(results.Items) != 1 {
		t.Errorf("Incorrect results length")
	}

	if len(results.Items[0].VirtualHostIds) != 1 {
		t.Errorf("Incorrect result virtualHostIds length")
	}

	if results.Items[0].IdentityMappingIds["Web"] == 2 {
		t.Errorf("Incorrect result identityMappingIds[Web]")
	}
}
