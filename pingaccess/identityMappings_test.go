package pingaccess

import (
	"testing"
)

func TestGetIdentityMappingsCommand(t *testing.T) {
	svc := config()

	input := GetIdentityMappingsCommandInput{}
	results, _, _ := svc.IdentityMappings.GetIdentityMappingsCommand(&input)
	if results == nil {
		t.Errorf("Marshalled object should not contain items")
	}
}
