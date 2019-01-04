package pingaccess

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestSitesRequestQueryParamsAreUsed(t *testing.T) {
	// Start a local HTTP server
	dat, _ := ioutil.ReadFile("test_data/example_websessions.json")
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		equals(t, req.URL.String(), "/pa-admin-api/v3/sites?filter=demo&name=demo2&numberPerPage=1&order=ASC&page=1&sortKey=audience")
		// Send response to be tested
		rw.Write(dat)
	}))
	// Close the server when test finishes
	defer server.Close()
	url, _ := url.Parse(server.URL)
	svc := config(url)

	input1 := GetSitesCommandInput{
		Page:          "1",
		NumberPerPage: "1",
		Filter:        "demo",
		Name:          "demo2",
		SortKey:       "audience",
		Order:         "ASC",
	}

	result1, resp1, err1 := svc.Sites.GetSitesCommand(&input1)
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
func TestVSitesErrorHandling(t *testing.T) {
	url, _ := url.Parse("wrong")
	svc := config(url)

	_, _, err := svc.Sites.GetSitesCommand(&GetSitesCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.Sites.AddSiteCommand(&AddSiteCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, err = svc.Sites.DeleteSiteCommand(&DeleteSiteCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.Sites.GetSiteCommand(&GetSiteCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.Sites.UpdateSiteCommand(&UpdateSiteCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
}
func TestSitesMethods(t *testing.T) {
	url, _ := url.Parse("https://localhost:9000")
	svc := config(url)

	// add a new site
	input1 := AddSiteCommandInput{
		Body: SiteView{
			Name:                    String("bar"),
			Targets:                 &[]*string{String("localhost:1234")},
			MaxConnections:          Int(-1),
			MaxWebSocketConnections: Int(-1),
		}}
	result1, resp1, err1 := svc.Sites.AddSiteCommand(&input1)
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if result1 == nil {
		t.Errorf("Unable the marshall results")
	}

	//do a get on all sites
	input2 := GetSitesCommandInput{}
	result2, resp2, err2 := svc.Sites.GetSitesCommand(&input2)
	if err2 != nil {
		t.Errorf("Unable to retrieve sites: %s", err2)
	}
	if resp2.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp2.StatusCode)
	}
	if result2 == nil {
		t.Errorf("Unable the marshall results")
	}

	//update the site
	input3 := UpdateSiteCommandInput{
		Id: result1.Id.String(),
		Body: SiteView{
			Name:                    String("bar"),
			Targets:                 &[]*string{String("localhost:1234"), String("localhost:1235")},
			MaxConnections:          Int(-1),
			MaxWebSocketConnections: Int(-1),
		}}
	result3, resp3, err3 := svc.Sites.UpdateSiteCommand(&input3)
	if err3 != nil {
		t.Errorf("Unable to update site: %s", err3)
	}
	if resp3.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp3.StatusCode)
	}
	if len(*result3.Targets) != len(*input3.Body.Targets) {
		t.Errorf("Failed to update site, expected: %d got: %d", len(*input3.Body.Targets), len(*result3.Targets))
	}

	//get the site and check the update
	input4 := GetSiteCommandInput{
		Id: result1.Id.String(),
	}
	result4, resp4, err4 := svc.Sites.GetSiteCommand(&input4)
	if err4 != nil {
		t.Errorf("Unable to get site: %s", err4)
	}
	if resp4.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp4.StatusCode)
	}
	if *(*result4.Targets)[0] != *(*input3.Body.Targets)[0] {
		t.Errorf("Failed to get site")
	}

	//delete our initial site
	input5 := DeleteSiteCommandInput{
		Id: result1.Id.String(),
	}
	resp5, err5 := svc.Sites.DeleteSiteCommand(&input5)
	if err5 != nil {
		t.Errorf("Unable to delete site: %s", err5)
	}
	if resp5.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp5.StatusCode)
	}
}
