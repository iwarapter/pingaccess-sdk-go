package pingaccess

import (
	"io/ioutil"
	"strconv"
	"testing"
)

func TestWebSessionsMethods(t *testing.T) {
	svc := config()

	// add a new websession
	input1 := AddWebSessionCommandInput{
		Body: WebSessionView{
			Audience: "some_random_folk",
			Name:     "my_test_websession",
			ClientCredentials: OAuthClientCredentialsView{
				ClientId: "my_client",
				ClientSecret: HiddenFieldView{
					Value: "my_secret",
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
	id := strconv.Itoa(result1.Id)

	//update the websession
	input3 := UpdateWebSessionCommandInput{
		Path: struct {
			Id string
		}{
			Id: id,
		},
		Body: WebSessionView{
			Audience: "some_new_random_folk",
			Name:     "my_test_websession",
			ClientCredentials: OAuthClientCredentialsView{
				ClientId: "my_client",
				ClientSecret: HiddenFieldView{
					Value: "my_secret",
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
	if result3.Audience != input3.Body.Audience {
		t.Errorf("Failed to update websession")
	}

	//get the websession and check the update
	input4 := GetWebSessionCommandInput{
		Path: struct {
			Id string
		}{
			Id: id,
		}}
	result4, resp4, err4 := svc.WebSessions.GetWebSessionCommand(&input4)
	if err4 != nil {
		t.Errorf("Unable to delete websession: %s", err4)
	}
	if resp4.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp4.StatusCode)
	}
	if result4.Audience != input3.Body.Audience {
		t.Errorf("Failed to update websession")
	}

	//delete our initial websession
	input5 := DeleteWebSessionCommandInput{
		Path: struct {
			Id string
		}{
			Id: id,
		}}
	resp5, err5 := svc.WebSessions.DeleteWebSessionCommand(&input5)
	if err5 != nil {
		t.Errorf("Unable to delete websession: %s", err5)
	}
	if resp5.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp5.StatusCode)
	}
}
