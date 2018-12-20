package pingaccess

import (
	"testing"
)

func TestGetIdentityMappingsCommand(t *testing.T) {
	svc := config()

	input := GetIdentityMappingsCommandInput{}
	results, _, _ := svc.IdentityMappings.GetIdentityMappingsCommand(&input)
	if len(results.Items) != 0 {
		t.Errorf("Marshalled object should not contain items")
	}
}
