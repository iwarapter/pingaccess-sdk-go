package pingaccess

import (
	"testing"
)

func TestAcmeServerMethods(t *testing.T) {
	svc := config(paURL)
	input1 := AddAcmeServerCommandInput{
		Body: AcmeServerView{
			Name: String("foo"),
			Url:  String("https://localhost:14000/dir"),
		}}

	result1, resp1, err1 := svc.Acme.AddAcmeServerCommand(&input1)
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1 == nil || resp1.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if result1 == nil {
		t.Fatalf("Unable the marshall results")
	}

	//do a get on all access token validator
	input2 := GetAcmeServerCommandInput{}
	result2, resp2, err2 := svc.Acme.GetAcmeServerCommand(&input2)
	if err2 != nil {
		t.Errorf("Unable to retrieve access token validator: %s", err2)
	}
	if resp2.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp2.StatusCode)
	}
	if result2 == nil {
		t.Errorf("Unable the marshall results")
	}

	//delete our initial access token validator
	input5 := DeleteAcmeServerCommandInput{
		AcmeServerId: *result1.Id,
	}
	_, resp5, err5 := svc.Acme.DeleteAcmeServerCommand(&input5)
	if err5 != nil {
		t.Errorf("Unable to delete access token validator: %s", err5)
	}
	if resp5.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp5.StatusCode)
	}
}
