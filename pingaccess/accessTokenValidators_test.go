package pingaccess

import (
	"testing"
)

func TestAccessTokenValidatorMethods(t *testing.T) {
	svc := config(paURL)
	input1 := AddAccessTokenValidatorCommandInput{
		Body: AccessTokenValidatorView{
			Name:      String("foo"),
			ClassName: String("com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"),
			Configuration: map[string]interface{}{
				"description":          "null",
				"path":                 "/foo",
				"subjectAttributeName": "foo",
				"issuer":               "null",
				"audience":             "null",
			},
		}}

	result1, resp1, err1 := svc.AccessTokenValidators.AddAccessTokenValidatorCommand(&input1)
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if result1 == nil {
		t.Errorf("Unable the marshall results")
	}

	//do a get on all access token validator
	input2 := GetAccessTokenValidatorCommandInput{}
	result2, resp2, err2 := svc.AccessTokenValidators.GetAccessTokenValidatorCommand(&input2)
	if err2 != nil {
		t.Errorf("Unable to retrieve access token validator: %s", err2)
	}
	if resp2.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp2.StatusCode)
	}
	if result2 == nil {
		t.Errorf("Unable the marshall results")
	}

	//update the access token validator
	input3 := UpdateAccessTokenValidatorCommandInput{
		Id: result1.Id.String(),
		Body: AccessTokenValidatorView{
			Name:      String("foo"),
			ClassName: String("com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"),
			Configuration: map[string]interface{}{
				"description":          "null",
				"path":                 "/foo/bar",
				"subjectAttributeName": "foo",
				"issuer":               "null",
				"audience":             "null",
			},
		},
	}
	result3, resp3, err3 := svc.AccessTokenValidators.UpdateAccessTokenValidatorCommand(&input3)
	if err3 != nil {
		t.Errorf("Unable to update access token validator: %s", err3)
	}
	if resp3.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp3.StatusCode)
	}
	if result3.Configuration["path"].(string) != "/foo/bar" {
		t.Errorf("Failed to get configuration[\"path\"]: %s", result3.Configuration["path"].(string))
	}

	//get the access token validator and check the update
	input4 := GetAccessTokenValidatorCommandInput{
		Id: result1.Id.String(),
	}
	result4, resp4, err4 := svc.AccessTokenValidators.GetAccessTokenValidatorCommand(&input4)
	if err4 != nil {
		t.Errorf("Unable to get access token validator: %s", err4)
	}
	if resp4.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp4.StatusCode)
	}
	if result4.Configuration["path"].(string) != "/foo/bar" {
		t.Errorf("Failed to get configuration[\"path\"]: %s", result4.Configuration["path"].(string))
	}

	//delete our initial access token validator
	input5 := DeleteAccessTokenValidatorCommandInput{
		Id: result1.Id.String(),
	}
	resp5, err5 := svc.AccessTokenValidators.DeleteAccessTokenValidatorCommand(&input5)
	if err5 != nil {
		t.Errorf("Unable to delete access token validator: %s", err5)
	}
	if resp5.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp5.StatusCode)
	}
}
