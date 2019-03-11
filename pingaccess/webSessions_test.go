package pingaccess

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestWebSessionsRequestQueryParamsAreUsed(t *testing.T) {
	// Start a local HTTP server
	dat, _ := ioutil.ReadFile("test_data/example_websessions.json")
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		equals(t, req.URL.String(), "/pa-admin-api/v3/webSessions?filter=demo&name=demo2&numberPerPage=1&order=ASC&page=1&sortKey=audience")
		// Send response to be tested
		rw.Write(dat)
	}))
	// Close the server when test finishes
	defer server.Close()
	url, _ := url.Parse(server.URL)
	svc := config(url)

	input1 := GetWebSessionsCommandInput{
		Page:          "1",
		NumberPerPage: "1",
		Filter:        "demo",
		Name:          "demo2",
		SortKey:       "audience",
		Order:         "ASC",
	}

	result1, resp1, err1 := svc.WebSessions.GetWebSessionsCommand(&input1)
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

func TestInvalidInputsCreateErrors(t *testing.T) {
	// Start a local HTTP server
	dat := []byte(`
	{
		"form": {},
		"flash": [
			"The 'page' parameter is not an integer: asdas"
		]
	}`)
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		equals(t, req.URL.String(), "/pa-admin-api/v3/webSessions?page=asdas")
		// Send response to be tested
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(dat)
	}))
	// Close the server when test finishes
	defer server.Close()
	url, _ := url.Parse(server.URL)
	svc := config(url)

	input1 := GetWebSessionsCommandInput{
		Page: "asdas",
	}

	result1, resp1, err1 := svc.WebSessions.GetWebSessionsCommand(&input1)
	if err1 == nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1.StatusCode != 400 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if result1 != nil {
		t.Errorf("Unable the marshall results")
	}

}

func TestWebSessionsErrorHandling(t *testing.T) {
	url, _ := url.Parse("wrong")
	svc := config(url)

	_, _, err := svc.WebSessions.GetWebSessionsCommand(&GetWebSessionsCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.WebSessions.AddWebSessionCommand(&AddWebSessionCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, err = svc.WebSessions.DeleteWebSessionCommand(&DeleteWebSessionCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.WebSessions.GetWebSessionCommand(&GetWebSessionCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.WebSessions.UpdateWebSessionCommand(&UpdateWebSessionCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
}

func TestWebSessionsMethods(t *testing.T) {
	svc := config(paURL)

	// add a new websession
	input1 := AddWebSessionCommandInput{
		Body: WebSessionView{
			Audience: String("some_random_folk"),
			Name:     String("my_test_websession"),
			ClientCredentials: &OAuthClientCredentialsView{
				ClientId: String("my_client"),
				ClientSecret: &HiddenFieldView{
					Value: String("my_secret"),
				},
			},
		}}
	result1, resp1, err1 := svc.WebSessions.AddWebSessionCommand(&input1)
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if result1 == nil {
		t.Errorf("Unable the marshall results")
	}

	//do a get on all websessions
	input2 := GetWebSessionsCommandInput{}
	result2, resp2, err2 := svc.WebSessions.GetWebSessionsCommand(&input2)
	if err2 != nil {
		t.Errorf("Unable to retrieve websessions: %s", err2)
	}
	if resp2.StatusCode != 200 {
		defer resp2.Body.Close()
		body, _ := ioutil.ReadAll(resp2.Body)
		t.Errorf("Invalid response code: %s", body)
	}
	if result2 == nil {
		t.Errorf("Unable the marshall results")
	}

	//update the websession
	input3 := UpdateWebSessionCommandInput{
		Id: result1.Id.String(),
		Body: WebSessionView{
			Audience: String("some_new_random_folk"),
			Name:     String("my_test_websession"),
			ClientCredentials: &OAuthClientCredentialsView{
				ClientId: String("my_client"),
				ClientSecret: &HiddenFieldView{
					Value: String("my_secret"),
				},
			},
		}}
	result3, resp3, err3 := svc.WebSessions.UpdateWebSessionCommand(&input3)
	if err3 != nil {
		t.Errorf("Unable to delete websession: %s", err3)
	}
	if resp3.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp3.StatusCode)
	}
	if *result3.Audience != *input3.Body.Audience {
		t.Errorf("Failed to update websession")
	}

	//get the websession and check the update
	input4 := GetWebSessionCommandInput{
		Id: result1.Id.String(),
	}
	result4, resp4, err4 := svc.WebSessions.GetWebSessionCommand(&input4)
	if err4 != nil {
		t.Errorf("Unable to delete websession: %s", err4)
	}
	if resp4.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp4.StatusCode)
	}
	if *result4.Audience != *input3.Body.Audience {
		t.Errorf("Failed to update websession")
	}

	//delete our initial websession
	input5 := DeleteWebSessionCommandInput{
		Id: result1.Id.String(),
	}
	resp5, err5 := svc.WebSessions.DeleteWebSessionCommand(&input5)
	if err5 != nil {
		t.Errorf("Unable to delete websession: %s", err5)
	}
	if resp5.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp5.StatusCode)
	}
}
