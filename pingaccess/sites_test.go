package pingaccess

import (
	"testing"
)

func TestGetSitesCommand(t *testing.T) {
	svc := config()

	input := GetSitesCommandInput{}
	results, _, _ := svc.Sites.GetSitesCommand(&input)
	if results == nil {
		t.Errorf("Marshelled object should contain items")
	}
}

func TestAddSiteCommand(t *testing.T) {
	svc := config()

	targets := []*string{}
	str := "localhost:1234"
	targets = append(targets, &str)
	input := AddSiteCommandInput{
		Body: SiteView{
			Name:                    "bar",
			Targets:                 targets,
			MaxConnections:          -1,
			MaxWebSocketConnections: -1,
		},
	}
	_, _, err := svc.Sites.AddSiteCommand(&input)
	if err != nil {
		t.Errorf("Unable to created execute command: %s", err.Error())
	}
}
