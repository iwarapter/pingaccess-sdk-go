package pingaccess

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestGetApplicationsCommand(t *testing.T) {
	svc := config()

	input := GetApplicationsCommandInput{}
	results, _, _ := svc.Applications.GetApplicationsCommand(&input)
	if results == nil {
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
