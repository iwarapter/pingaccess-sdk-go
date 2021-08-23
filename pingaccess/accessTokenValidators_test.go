package pingaccess_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/config"
	pa "github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/accessTokenValidators"
)

func TestAccessTokenValidatorMethods(t *testing.T) {
	svc := accessTokenValidators.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL).WithDebug(false))
	input1 := accessTokenValidators.AddAccessTokenValidatorCommandInput{
		Body: pa.AccessTokenValidatorView{
			Name:      pingaccess.String("foo"),
			ClassName: pingaccess.String("com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"),
			Configuration: map[string]interface{}{
				"description":          "null",
				"path":                 "/foo",
				"subjectAttributeName": "foo",
				"issuer":               "null",
				"audience":             "null",
			},
		}}

	result1, resp1, err1 := svc.AddAccessTokenValidatorCommand(&input1)
	require.NoError(t, err1)
	assert.Equal(t, http.StatusOK, resp1.StatusCode)
	assert.Equal(t, "foo", *result1.Name)

	//do a get on all access token validator
	input2 := accessTokenValidators.GetAccessTokenValidatorCommandInput{}
	result2, resp2, err2 := svc.GetAccessTokenValidatorCommand(&input2)
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
	input3 := accessTokenValidators.UpdateAccessTokenValidatorCommandInput{
		Id: result1.Id.String(),
		Body: pa.AccessTokenValidatorView{
			Name:      pingaccess.String("foo"),
			ClassName: pingaccess.String("com.pingidentity.pa.accesstokenvalidators.JwksEndpoint"),
			Configuration: map[string]interface{}{
				"description":          "null",
				"path":                 "/foo/bar",
				"subjectAttributeName": "foo",
				"issuer":               "null",
				"audience":             "null",
			},
		},
	}
	result3, resp3, err3 := svc.UpdateAccessTokenValidatorCommand(&input3)
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
	input4 := accessTokenValidators.GetAccessTokenValidatorCommandInput{
		Id: result1.Id.String(),
	}
	result4, resp4, err4 := svc.GetAccessTokenValidatorCommand(&input4)
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
	input5 := accessTokenValidators.DeleteAccessTokenValidatorCommandInput{
		Id: result1.Id.String(),
	}
	resp5, err5 := svc.DeleteAccessTokenValidatorCommand(&input5)
	if err5 != nil {
		t.Errorf("Unable to delete access token validator: %s", err5)
	}
	if resp5.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp5.StatusCode)
	}
}
