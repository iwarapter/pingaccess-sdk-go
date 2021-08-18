package pingaccess_test

import (
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/config"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/siteAuthenticators"

	"testing"
)

func TestSiteAuthenticatorDescriptors(t *testing.T) {
	svc := siteAuthenticators.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL).WithDebug(false))

	result1, resp1, err1 := svc.GetSiteAuthenticatorDescriptorsCommand()
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if len(result1.Items) != 3 {
		t.Errorf("Unable the marshall results")
	}

	input2 := siteAuthenticators.GetSiteAuthenticatorDescriptorCommandInput{
		SiteAuthenticatorType: "mutualtlssiteauthenticator",
	}
	result2, resp2, err2 := svc.GetSiteAuthenticatorDescriptorCommand(&input2)
	if err2 != nil {
		t.Errorf("Unable to execute command: %s", err2.Error())
	}
	if resp2.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp2.StatusCode)
	}
	if *result2.ClassName != "com.pingidentity.pa.siteauthenticators.MutualTlsSiteAuthenticator" {
		t.Errorf("Unable the marshall results")
	}
}

func TestSiteAuthenticatorMethods(t *testing.T) {
	svc := siteAuthenticators.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL).WithDebug(false))

	// add a new site authenticator
	input1 := siteAuthenticators.AddSiteAuthenticatorCommandInput{
		Body: pa.SiteAuthenticatorView{
			ClassName: pingaccess.String("com.pingidentity.pa.siteauthenticators.MutualTlsSiteAuthenticator"),
			Name:      pingaccess.String("matls"),
			Configuration: map[string]interface{}{
				"keyPairId": "2",
			},
		}}
	result1, resp1, err1 := svc.AddSiteAuthenticatorCommand(&input1)
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1 == nil || resp1.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if result1 == nil {
		t.Fatalf("Unable the marshall results")
	}

	//do a get on all site authenticators
	input2 := siteAuthenticators.GetSiteAuthenticatorsCommandInput{}
	result2, resp2, err2 := svc.GetSiteAuthenticatorsCommand(&input2)
	if err2 != nil {
		t.Errorf("Unable to retrieve site authenticators: %s", err2)
	}
	if resp2.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp2.StatusCode)
	}
	if result2 == nil {
		t.Errorf("Unable the marshall results")
	}

	//update the site authenticator
	input3 := siteAuthenticators.UpdateSiteAuthenticatorCommandInput{
		Id: result1.Id.String(),
		Body: pa.SiteAuthenticatorView{
			ClassName: pingaccess.String("com.pingidentity.pa.siteauthenticators.MutualTlsSiteAuthenticator"),
			Name:      pingaccess.String("matls"),
			Configuration: map[string]interface{}{
				"keyPairId": "3",
			},
		}}
	result3, resp3, err3 := svc.UpdateSiteAuthenticatorCommand(&input3)
	if err3 != nil {
		t.Errorf("Unable to update site authenticator: %s", err3)
	}
	if resp3.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp3.StatusCode)
	}
	if result3.Configuration["KeyPairId"] != input3.Body.Configuration["KeyPairId"] {
		t.Errorf("Failed to update site authenticator, expected: %s got: %s", result3.Configuration["KeyPairId"], input3.Body.Configuration["KeyPairId"])
	}

	//get the site authenticator and check the update
	input4 := siteAuthenticators.GetSiteAuthenticatorCommandInput{
		Id: result1.Id.String(),
	}
	result4, resp4, err4 := svc.GetSiteAuthenticatorCommand(&input4)
	if err4 != nil {
		t.Errorf("Unable to get site authenticator: %s", err4)
	}
	if resp4.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp4.StatusCode)
	}
	if result4.Configuration["KeyPairId"] != input3.Body.Configuration["KeyPairId"] {
		t.Errorf("Failed to get site authenticator")
	}

	//delete our initial site authenticator
	input5 := siteAuthenticators.DeleteSiteAuthenticatorCommandInput{
		Id: result1.Id.String(),
	}
	resp5, err5 := svc.DeleteSiteAuthenticatorCommand(&input5)
	if err5 != nil {
		t.Errorf("Unable to delete site authenticator: %s", err5)
	}
	if resp5.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp5.StatusCode)
	}
}
