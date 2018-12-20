package pingaccess

import (
	"testing"
)

func TestGetVirtualHostsCommand(t *testing.T) {
	svc := config()

	input := GetVirtualHostsCommandInput{}
	results, _, _ := svc.Virtualhosts.GetVirtualHostsCommand(&input)
	if len(results.Items) == 0 {
		t.Errorf("Marshelled object should contain items")
	}
}

func TestAddVirtualHostCommand(t *testing.T) {
	svc := config()

	input := AddVirtualHostCommandInput{
		Body: VirtualHostView{
			AgentResourceCacheTTL: 900,
			Host:      "localhost",
			KeyPairId: 0,
			Port:      3000,
			TrustedCertificateGroupId: 0,
		}}
	_, _, err := svc.Virtualhosts.AddVirtualHostCommand(&input)
	if err != nil {
		t.Errorf("Unable to created execute command: %s", err.Error())
	}
}
