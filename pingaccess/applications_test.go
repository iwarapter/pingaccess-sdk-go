package pingaccess

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestApplicationsRequestQueryParamsAreUsed(t *testing.T) {
	// Start a local HTTP server
	dat, _ := ioutil.ReadFile("test_data/example_applications.json")
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		equals(t, req.URL.String(), "/pa-admin-api/v3/applications?agentId=1&filter=demo&name=demo2&numberPerPage=1&order=ASC&page=1&siteId=1&sortKey=audience&virtualHostId=1")
		// Send response to be tested
		rw.Write(dat)
	}))
	// Close the server when test finishes
	defer server.Close()
	url, _ := url.Parse(server.URL)
	svc := NewClient("Administrator", "2Access2", url, nil)

	input1 := GetApplicationsCommandInput{
		Page:          "1",
		NumberPerPage: "1",
		Filter:        "demo",
		Name:          "demo2",
		SortKey:       "audience",
		Order:         "ASC",
		SiteId:        "1",
		VirtualHostId: "1",
		AgentId:       "1",
	}

	result1, resp1, err1 := svc.Applications.GetApplicationsCommand(&input1)
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if result1 == nil {
		t.Errorf("Unable the marshall results")
	}

}

func TestGetApplicationResourcesCommandQueryParamsAreUsed(t *testing.T) {
	// Start a local HTTP server
	dat, _ := ioutil.ReadFile("test_data/example_applications.json")
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		equals(t, req.URL.String(), "/pa-admin-api/v3/applications/1/resources?filter=demo&name=demo2&numberPerPage=1&order=ASC&page=1&sortKey=audience")
		// Send response to be tested
		rw.Write(dat)
	}))
	// Close the server when test finishes
	defer server.Close()
	url, _ := url.Parse(server.URL)
	svc := NewClient("Administrator", "2Access2", url, nil)

	input1 := GetApplicationResourcesCommandInput{
		Page:          "1",
		NumberPerPage: "1",
		Filter:        "demo",
		Name:          "demo2",
		SortKey:       "audience",
		Order:         "ASC",
		Id:            "1",
	}

	result1, resp1, err1 := svc.Applications.GetApplicationResourcesCommand(&input1)
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if result1 == nil {
		t.Errorf("Unable the marshall results")
	}

}

func TestVApplicationsErrorHandling(t *testing.T) {
	url, _ := url.Parse("wrong")
	svc := NewClient("Administrator", "2Access2", url, nil)

	_, _, err := svc.Applications.GetApplicationsCommand(&GetApplicationsCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.Applications.AddApplicationCommand(&AddApplicationCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.Applications.DeleteApplicationCommand(&DeleteApplicationCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.Applications.GetApplicationCommand(&GetApplicationCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.Applications.UpdateApplicationCommand(&UpdateApplicationCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.Applications.GetResourcesCommand(&GetResourcesCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.Applications.GetApplicationsResourcesMethodsCommand()
	if err == nil {
		t.Errorf("This should go bang")
	}
}

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
