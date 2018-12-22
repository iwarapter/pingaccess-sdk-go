package pingaccess

import (
	"strconv"
	"testing"
)

func TestSitesMethods(t *testing.T) {
	svc := config()

	// add a new site
	input1 := AddSiteCommandInput{
		Body: SiteView{
			Name:                    "bar",
			Targets:                 []string{"localhost:1234"},
			MaxConnections:          -1,
			MaxWebSocketConnections: -1,
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
	id := strconv.Itoa(result1.Id)

	//update the site
	input3 := UpdateSiteCommandInput{
		Path: struct {
			Id string
		}{
			Id: id,
		},
		Body: SiteView{
			Name:                    "bar",
			Targets:                 []string{"localhost:1234", "localhost:1235"},
			MaxConnections:          -1,
			MaxWebSocketConnections: -1,
		}}
	result3, resp3, err3 := svc.Sites.UpdateSiteCommand(&input3)
	if err3 != nil {
		t.Errorf("Unable to update site: %s", err3)
	}
	if resp3.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp3.StatusCode)
	}
	if len(result3.Targets) != len(input3.Body.Targets) {
		t.Errorf("Failed to update site, expected: %d got: %d", len(input3.Body.Targets), len(result3.Targets))
	}

	//get the site and check the update
	input4 := GetSiteCommandInput{
		Path: struct {
			Id string
		}{
			Id: id,
		}}
	result4, resp4, err4 := svc.Sites.GetSiteCommand(&input4)
	if err4 != nil {
		t.Errorf("Unable to get site: %s", err4)
	}
	if resp4.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp4.StatusCode)
	}
	if result4.Targets[0] != input3.Body.Targets[0] {
		t.Errorf("Failed to get site")
	}

	//delete our initial site
	input5 := DeleteSiteCommandInput{
		Path: struct {
			Id string
		}{
			Id: id,
		}}
	resp5, err5 := svc.Sites.DeleteSiteCommand(&input5)
	if err5 != nil {
		t.Errorf("Unable to delete site: %s", err5)
	}
	if resp5.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp5.StatusCode)
	}
}
