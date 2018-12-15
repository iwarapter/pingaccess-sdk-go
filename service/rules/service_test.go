package rules

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func TestWeCanMarshallRules(t *testing.T) {
	results := RulesView{}
	dat, err := ioutil.ReadFile("test_data/example_rules.json")
	if err != nil {
		t.Errorf("Unable to read test file")
		t.FailNow()
	}
	err = json.Unmarshal(dat, &results)
	if err != nil {
		t.Errorf("Unable to marshall example response")
	}

	if len(results.Items) != 3 {
		t.Errorf("Incorrect result length")
	}

	if len(results.Items) != 3 {
		t.Errorf("Incorrect results length")
	}

	if len(results.Items[0].SupportedDestinations) != 2 {
		t.Errorf("Incorrect result supported destinations length")
	}

	if len(results.Items[0].Configuration) != 10 {
		t.Errorf("Incorrect result configuration length")
	}
}

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

func TestGetRulesCommand(t *testing.T) {
	svc := New(config())

	input := GetRulesCommandInput{}
	results, _ := svc.GetRulesCommand(&input)
	if len(results.Items) != 5 {
		t.Errorf("Marshelled object should contain items")
	}
}
