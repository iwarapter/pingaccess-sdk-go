package pingaccess

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestVirtualhostsRequestQueryParamsAreUsed(t *testing.T) {
	// Start a local HTTP server
	dat, _ := ioutil.ReadFile("test_data/example_websessions.json")
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		equals(t, req.URL.String(), "/pa-admin-api/v3/virtualhosts?filter=demo&numberPerPage=1&order=ASC&page=1&sortKey=audience&virtualHost=demo2")
		// Send response to be tested
		rw.Write(dat)
	}))
	// Close the server when test finishes
	defer server.Close()
	url, _ := url.Parse(server.URL)
	svc := NewClient("Administrator", "2Access2", url, nil)

	input1 := GetVirtualHostsCommandInput{
		Page:          "1",
		NumberPerPage: "1",
		Filter:        "demo",
		VirtualHost:   "demo2",
		SortKey:       "audience",
		Order:         "ASC",
	}

	result1, resp1, err1 := svc.Virtualhosts.GetVirtualHostsCommand(&input1)
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
func TestVirtualhostsErrorHandling(t *testing.T) {
	url, _ := url.Parse("wrong")
	svc := NewClient("Administrator", "2Access2", url, nil)

	_, _, err := svc.Virtualhosts.GetVirtualHostsCommand(&GetVirtualHostsCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.Virtualhosts.AddVirtualHostCommand(&AddVirtualHostCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, err = svc.Virtualhosts.DeleteVirtualHostCommand(&DeleteVirtualHostCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.Virtualhosts.GetVirtualHostCommand(&GetVirtualHostCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.Virtualhosts.UpdateVirtualHostCommand(&UpdateVirtualHostCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
}

func TestVirtualHostsMethods(t *testing.T) {
	svc := config()

	// add a new virtualhost
	input1 := AddVirtualHostCommandInput{
		Body: VirtualHostView{
			AgentResourceCacheTTL: 900,
			Host:      "localhost",
			KeyPairId: 0,
			Port:      3000,
			TrustedCertificateGroupId: 0,
		}}
	result1, resp1, err1 := svc.Virtualhosts.AddVirtualHostCommand(&input1)
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if result1 == nil {
		t.Errorf("Unable the marshall results")
	}

	//do a get on all virtualhosts
	input2 := GetVirtualHostsCommandInput{}
	result2, resp2, err2 := svc.Virtualhosts.GetVirtualHostsCommand(&input2)
	if err2 != nil {
		t.Errorf("Unable to retrieve virtualhosts: %s", err2)
	}
	if resp2.StatusCode != 200 {
		defer resp2.Body.Close()
		body, _ := ioutil.ReadAll(resp2.Body)
		t.Errorf("Invalid response code: %s", body)
	}
	if result2 == nil {
		t.Errorf("Unable the marshall results")
	}
	id := strconv.Itoa(result1.Id)

	//update the virtualhost
	input3 := UpdateVirtualHostCommandInput{
		Id: id,
		Body: VirtualHostView{
			AgentResourceCacheTTL: 900,
			Host:      "localhost",
			KeyPairId: 0,
			Port:      3001,
			TrustedCertificateGroupId: 0,
		}}
	result3, resp3, err3 := svc.Virtualhosts.UpdateVirtualHostCommand(&input3)
	if err3 != nil {
		t.Errorf("Unable to delete virtualhost: %s", err3)
	}
	if resp3.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp3.StatusCode)
	}
	if result3.Port != input3.Body.Port {
		t.Errorf("Failed to update virtualhost")
	}

	//get the virtualhost and check the update
	input4 := GetVirtualHostCommandInput{
		Id: id,
	}
	result4, resp4, err4 := svc.Virtualhosts.GetVirtualHostCommand(&input4)
	if err4 != nil {
		t.Errorf("Unable to delete virtualhost: %s", err4)
	}
	if resp4.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp4.StatusCode)
	}
	if result4.Port != input3.Body.Port {
		t.Errorf("Failed to update virtualhost")
	}

	//delete our initial virtualhost
	input5 := DeleteVirtualHostCommandInput{
		Id: id,
	}
	resp5, err5 := svc.Virtualhosts.DeleteVirtualHostCommand(&input5)
	if err5 != nil {
		t.Errorf("Unable to delete virtualhost: %s", err5)
	}
	if resp5.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp5.StatusCode)
	}
}
