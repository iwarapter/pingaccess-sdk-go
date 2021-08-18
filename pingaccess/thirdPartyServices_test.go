package pingaccess_test

import (
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/config"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/thirdPartyServices"
)

func TestThirdPartyServicesRequestQueryParamsAreUsed(t *testing.T) {
	svc := thirdPartyServices.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL).WithDebug(true))

	input1 := thirdPartyServices.GetThirdPartyServicesCommandInput{
		Page:          "1",
		NumberPerPage: "1",
		Filter:        "demo",
		Name:          "demo2",
		SortKey:       "hostValue",
		Order:         "ASC",
	}

	result1, resp1, err1 := svc.GetThirdPartyServicesCommand(&input1)
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
	svc := thirdPartyServices.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint("wrong").WithDebug(false))

	_, _, err := svc.GetThirdPartyServicesCommand(&thirdPartyServices.GetThirdPartyServicesCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.AddThirdPartyServiceCommand(&thirdPartyServices.AddThirdPartyServiceCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, err = svc.DeleteThirdPartyServiceCommand(&thirdPartyServices.DeleteThirdPartyServiceCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.GetThirdPartyServiceCommand(&thirdPartyServices.GetThirdPartyServiceCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
	_, _, err = svc.UpdateThirdPartyServiceCommand(&thirdPartyServices.UpdateThirdPartyServiceCommandInput{})
	if err == nil {
		t.Errorf("This should go bang")
	}
}

func TestThirdPartyServicesMethods(t *testing.T) {
	svc := thirdPartyServices.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL).WithDebug(false))

	// add a new ThirdPartyService
	input1 := thirdPartyServices.AddThirdPartyServiceCommandInput{
		Body: pa.ThirdPartyServiceView{
			Name:                  pingaccess.String("bar"),
			Targets:               &[]*string{pingaccess.String("localhost:1234")},
			AvailabilityProfileId: pingaccess.Int(1),
		}}
	result1, resp1, err1 := svc.AddThirdPartyServiceCommand(&input1)
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
	input2 := thirdPartyServices.GetThirdPartyServicesCommandInput{}
	result2, resp2, err2 := svc.GetThirdPartyServicesCommand(&input2)
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
	input3 := thirdPartyServices.UpdateThirdPartyServiceCommandInput{
		Id: *result1.Id,
		Body: pa.ThirdPartyServiceView{
			Name:                  pingaccess.String("bar"),
			Targets:               &[]*string{pingaccess.String("localhost:1234"), pingaccess.String("localhost:1235")},
			AvailabilityProfileId: pingaccess.Int(1),
		}}
	result3, resp3, err3 := svc.UpdateThirdPartyServiceCommand(&input3)
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
	input4 := thirdPartyServices.GetThirdPartyServiceCommandInput{
		Id: *result1.Id,
	}
	result4, resp4, err4 := svc.GetThirdPartyServiceCommand(&input4)
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
	input5 := thirdPartyServices.DeleteThirdPartyServiceCommandInput{
		Id: *result1.Id,
	}
	resp5, err5 := svc.DeleteThirdPartyServiceCommand(&input5)
	if err5 != nil {
		t.Errorf("Unable to delete ThirdPartyService: %s", err5)
	}
	if resp5.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp5.StatusCode)
	}
}
