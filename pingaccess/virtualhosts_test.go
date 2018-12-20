package pingaccess

import (
	"io/ioutil"
	"strconv"
	"testing"
)

func TestVirtualHostsMethods(t *testing.T) {
	svc := config()

	// add a new virtualhost
	input1 := AddVirtualHostCommandInput{
		Body: VirtualHostView{
			AgentResourceCacheTTL:     900,
			Host:                      "localhost",
			KeyPairId:                 0,
			Port:                      3000,
			TrustedCertificateGroupId: 0,
		}}
	result1, resp1, err1 := svc.Virtualhosts.AddVirtualHostCommand(&input1)
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1.StatusCode != 200 {
		defer resp1.Body.Close()
		body, _ := ioutil.ReadAll(resp1.Body)
		t.Errorf("Response Code: %d", resp1.StatusCode)
		t.Errorf("Invalid response code: %s", body)
	}
	if result1 == nil {
		t.Errorf("Unable the marshall results")
	}

	//do a get on all virtualhosts
	input2 := GetVirtualHostsCommandInput{}
	result2, resp2, err2 := svc.Virtualhosts.GetVirtualHostsCommand(&input2)
	if err2 != nil {
		t.Errorf("Unable to retrieve virtualhosts: %s", err2)
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

	//update the virtualhost
	input3 := UpdateVirtualHostCommandInput{
		Path: struct {
			Id string
		}{
			Id: id,
		},
		Body: VirtualHostView{
			AgentResourceCacheTTL:     900,
			Host:                      "localhost",
			KeyPairId:                 0,
			Port:                      3001,
			TrustedCertificateGroupId: 0,
		}}
	result3, resp3, err3 := svc.Virtualhosts.UpdateVirtualHostCommand(&input3)
	if err3 != nil {
		t.Errorf("Unable to delete virtualhost: %s", err3)
	}
	if resp3.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp3.StatusCode)
	}
	if result3.Port != input3.Body.Port {
		t.Errorf("Failed to update virtualhost")
	}

	//get the virtualhost and check the update
	input4 := GetVirtualHostCommandInput{
		Path: struct {
			Id string
		}{
			Id: id,
		}}
	result4, resp4, err4 := svc.Virtualhosts.GetVirtualHostCommand(&input4)
	if err4 != nil {
		t.Errorf("Unable to delete virtualhost: %s", err4)
	}
	if resp4.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp4.StatusCode)
	}
	if result4.Port != input3.Body.Port {
		t.Errorf("Failed to update virtualhost")
	}

	//delete our initial virtualhost
	input5 := DeleteVirtualHostCommandInput{
		Path: struct {
			Id string
		}{
			Id: id,
		}}
	resp5, err5 := svc.Virtualhosts.DeleteVirtualHostCommand(&input5)
	if err5 != nil {
		t.Errorf("Unable to delete virtualhost: %s", err5)
	}
	if resp5.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp5.StatusCode)
	}
}
