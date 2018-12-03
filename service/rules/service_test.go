package rules

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestWeCanMarshallRules(t *testing.T) {
	results := RulesView{}
	dat, err := ioutil.ReadFile("test_data/example_rules.json")
	if err != nil {
		t.Errorf("Unable to read test file")
		t.FailNow()
	}
	err = json.Unmarshal(dat, &results)
	if err != nil {
		t.Errorf("Unable to marshall example response")
	}

	if len(results.Items) != 3 {
		t.Errorf("Incorrect result length")
	}

	if len(results.Items) != 3 {
		t.Errorf("Incorrect results length")
	}

	if len(results.Items[0].SupportedDestinations) != 2 {
		t.Errorf("Incorrect result supported destinations length")
	}

	if len(results.Items[0].Configuration) != 10 {
		t.Errorf("Incorrect result configuration length")
	}
}
