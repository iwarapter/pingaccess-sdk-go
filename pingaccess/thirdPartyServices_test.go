package pingaccess

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestThirdPartyServicesRequestQueryParamsAreUsed(t *testing.T) {
	// Start a local HTTP server
	dat, _ := ioutil.ReadFile("test_data/example_thirdPartyServices.json")
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		equals(t, req.URL.String(), "/pa-admin-api/v3/thirdPartyServices?filter=demo&name=demo2&numberPerPage=1&order=ASC&page=1&sortKey=audience")
		// Send response to be tested
		rw.Write(dat)
	}))
	// Close the server when test finishes
	defer server.Close()
	url, _ := url.Parse(server.URL)
	svc := config(url)

	input1 := GetThirdPartyServicesCommandInput{
		Page:          "1",
		NumberPerPage: "1",
		Filter:        "demo",
		Name:          "demo2",
		SortKey:       "audience",
		Order:         "ASC",
	}

	result1, resp1, err1 := svc.ThirdPartyServices.GetThirdPartyServicesCommand(&input1)
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
func TestVThirdPartyServicesErrorHandling(t *testing.T) {
	url, _ := url.Parse("wrong")
	svc := config(url)

	_, _, err := svc.ThirdPartyServices.GetThirdPartyServicesCommand(&GetThirdPartyServicesCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.ThirdPartyServices.AddThirdPartyServiceCommand(&AddThirdPartyServiceCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, err = svc.ThirdPartyServices.DeleteThirdPartyServiceCommand(&DeleteThirdPartyServiceCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.ThirdPartyServices.GetThirdPartyServiceCommand(&GetThirdPartyServiceCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.ThirdPartyServices.UpdateThirdPartyServiceCommand(&UpdateThirdPartyServiceCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
}

func TestThirdPartyServicesMethods(t *testing.T) {
	svc := config(paURL)

	// add a new ThirdPartyService
	input1 := AddThirdPartyServiceCommandInput{
		Body: ThirdPartyServiceView{
			Name:                  String("bar"),
			Targets:               &[]*string{String("localhost:1234")},
			AvailabilityProfileId: Int(1),
		}}
	result1, resp1, err1 := svc.ThirdPartyServices.AddThirdPartyServiceCommand(&input1)
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1 == nil || resp1.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if result1 == nil {
		t.Fatalf("Unable the marshall results")
	}

	//do a get on all ThirdPartyServices
	input2 := GetThirdPartyServicesCommandInput{}
	result2, resp2, err2 := svc.ThirdPartyServices.GetThirdPartyServicesCommand(&input2)
	if err2 != nil {
		t.Errorf("Unable to retrieve ThirdPartyServices: %s", err2)
	}
	if resp2.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp2.StatusCode)
	}
	if result2 == nil {
		t.Errorf("Unable the marshall results")
	}

	//update the ThirdPartyService
	input3 := UpdateThirdPartyServiceCommandInput{
		Id: *result1.Id,
		Body: ThirdPartyServiceView{
			Name:                  String("bar"),
			Targets:               &[]*string{String("localhost:1234"), String("localhost:1235")},
			AvailabilityProfileId: Int(1),
		}}
	result3, resp3, err3 := svc.ThirdPartyServices.UpdateThirdPartyServiceCommand(&input3)
	if err3 != nil {
		t.Errorf("Unable to update ThirdPartyService: %s", err3)
	}
	if resp3.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp3.StatusCode)
	}
	if len(*result3.Targets) != len(*input3.Body.Targets) {
		t.Errorf("Failed to update ThirdPartyService, expected: %d got: %d", len(*input3.Body.Targets), len(*result3.Targets))
	}

	//get the ThirdPartyService and check the update
	input4 := GetThirdPartyServiceCommandInput{
		Id: *result1.Id,
	}
	result4, resp4, err4 := svc.ThirdPartyServices.GetThirdPartyServiceCommand(&input4)
	if err4 != nil {
		t.Errorf("Unable to get ThirdPartyService: %s", err4)
	}
	if resp4.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp4.StatusCode)
	}
	if *(*result4.Targets)[0] != *(*input3.Body.Targets)[0] {
		t.Errorf("Failed to get ThirdPartyService")
	}

	//delete our initial ThirdPartyService
	input5 := DeleteThirdPartyServiceCommandInput{
		Id: *result1.Id,
	}
	resp5, err5 := svc.ThirdPartyServices.DeleteThirdPartyServiceCommand(&input5)
	if err5 != nil {
		t.Errorf("Unable to delete ThirdPartyService: %s", err5)
	}
	if resp5.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp5.StatusCode)
	}
}
