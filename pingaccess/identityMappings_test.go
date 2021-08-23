package pingaccess_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/iwarapter/pingaccess-sdk-go/v60/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/v60/pingaccess/config"
	pa "github.com/iwarapter/pingaccess-sdk-go/v60/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v60/services/identityMappings"
)

func TestIdentityMappingDescriptors(t *testing.T) {
	svc := identityMappings.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL).WithDebug(false))
	result1, resp1, err1 := svc.GetIdentityMappingDescriptorsCommand()
	require.NoError(t, err1)
	assert.Equal(t, http.StatusOK, resp1.StatusCode)
	assert.Len(t, result1.Items, 2)

	input2 := identityMappings.GetIdentityMappingDescriptorCommandInput{
		IdentityMappingType: "jwtidentitymapping",
	}
	result2, resp2, err2 := svc.GetIdentityMappingDescriptorCommand(&input2)
	if err2 != nil {
		t.Errorf("Unable to execute command: %s", err2.Error())
	}
	if resp2.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp2.StatusCode)
	}
	if *result2.ClassName != "com.pingidentity.pa.identitymappings.JwtIdentityMapping" {
		t.Errorf("Unable the marshall results")
	}
}

func TestIdentityMappingMethods(t *testing.T) {
	svc := identityMappings.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL).WithDebug(false))

	// add a new identity mapping
	input1 := identityMappings.AddIdentityMappingCommandInput{
		Body: pa.IdentityMappingView{
			ClassName: pingaccess.String("com.pingidentity.pa.identitymappings.HeaderIdentityMapping"),
			Name:      pingaccess.String("woot"),
			Configuration: map[string]interface{}{
				"attributeHeaderMappings": []interface{}{
					map[string]interface{}{
						"subject":       true,
						"attributeName": "USER",
						"headerName":    "USER",
					},
				},
				"headerClientCertificateMappings": []interface{}{},
			},
		}}
	result1, resp1, err1 := svc.AddIdentityMappingCommand(&input1)
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1 == nil || resp1.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if result1 == nil {
		t.Fatalf("Unable the marshall results")
	}

	//do a get on all identity mappings
	input2 := identityMappings.GetIdentityMappingsCommandInput{}
	result2, resp2, err2 := svc.GetIdentityMappingsCommand(&input2)
	if err2 != nil {
		t.Errorf("Unable to retrieve identity mappings: %s", err2)
	}
	if resp2.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp2.StatusCode)
	}
	if result2 == nil {
		t.Errorf("Unable the marshall results")
	}

	//update the identity mapping
	input3 := identityMappings.UpdateIdentityMappingCommandInput{
		Id: result1.Id.String(),
		Body: pa.IdentityMappingView{
			ClassName: pingaccess.String("com.pingidentity.pa.identitymappings.HeaderIdentityMapping"),
			Name:      pingaccess.String("woot"),
			Configuration: map[string]interface{}{
				"attributeHeaderMappings": []interface{}{
					map[string]interface{}{
						"subject":       true,
						"attributeName": "USER",
						"headerName":    "NEWUSER",
					},
				},
				"headerClientCertificateMappings": []interface{}{},
			},
		}}
	result3, resp3, err3 := svc.UpdateIdentityMappingCommand(&input3)
	if err3 != nil {
		t.Errorf("Unable to update identity mapping: %s", err3)
	}
	if resp3.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp3.StatusCode)
	}
	if result3.Configuration["KeyPairId"] != input3.Body.Configuration["KeyPairId"] {
		t.Errorf("Failed to update identity mapping, expected: %s got: %s", result3.Configuration["KeyPairId"], input3.Body.Configuration["KeyPairId"])
	}

	//get the identity mapping and check the update
	input4 := identityMappings.GetIdentityMappingCommandInput{
		Id: result1.Id.String(),
	}
	result4, resp4, err4 := svc.GetIdentityMappingCommand(&input4)
	if err4 != nil {
		t.Errorf("Unable to get identity mapping: %s", err4)
	}
	if resp4.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp4.StatusCode)
	}
	if result4.Configuration["KeyPairId"] != input3.Body.Configuration["KeyPairId"] {
		t.Errorf("Failed to get identity mapping")
	}

	//delete our initial identity mapping
	input5 := identityMappings.DeleteIdentityMappingCommandInput{
		Id: result1.Id.String(),
	}
	resp5, err5 := svc.DeleteIdentityMappingCommand(&input5)
	if err5 != nil {
		t.Errorf("Unable to delete identity mapping: %s", err5)
	}
	if resp5.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp5.StatusCode)
	}
}
