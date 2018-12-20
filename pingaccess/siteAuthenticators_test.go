package pingaccess

import (
	"testing"
)

func TestGetSiteAuthenticatorsCommand(t *testing.T) {
	svc := config()

	input := GetSiteAuthenticatorsCommandInput{}
	results, _, _ := svc.SiteAuthenticators.GetSiteAuthenticatorsCommand(&input)
	if results == nil {
		t.Errorf("Marshelled object should not contain items")
	}
}
