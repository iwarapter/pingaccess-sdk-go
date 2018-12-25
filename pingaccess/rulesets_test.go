package pingaccess

import (
	"strconv"
	"testing"
)

func TestRuleSetElementTypes(t *testing.T) {
	svc := config()

	result1, resp1, err1 := svc.Rulesets.GetRuleSetElementTypesCommand()
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if len(result1.Items) != 2 {
		t.Errorf("Unable the marshall results")
	}
}

func TestRuleSetSuccessCriteria(t *testing.T) {
	svc := config()

	result, resp, err := svc.Rulesets.GetRuleSetSuccessCriteriaCommand()
	if err != nil {
		t.Errorf("Unable to execute command: %s", err.Error())
	}
	if resp.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp.StatusCode)
	}
	if len(result.Items) != 2 {
		t.Errorf("Unable the marshall results")
	}
}

func TestRuleSetMethods(t *testing.T) {
	svc := config()

	input1 := AddRuleSetCommandInput{
		Body: RuleSetView{
			Name:            "new-rule-set-test",
			SuccessCriteria: "SuccessIfAllSucceed",
			ElementType:     "Rule",
			Policy:          []int{2},
		}}

	result1, resp1, err1 := svc.Rulesets.AddRuleSetCommand(&input1)
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if result1 == nil {
		t.Errorf("Unable the marshall results")
	}

	//do a get on all rules
	input2 := GetRuleSetsCommandInput{}
	result2, resp2, err2 := svc.Rulesets.GetRuleSetsCommand(&input2)
	if err2 != nil {
		t.Errorf("Unable to retrieve rules: %s", err2)
	}
	if resp2.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp2.StatusCode)
	}
	if result2 == nil {
		t.Errorf("Unable the marshall results")
	}
	id := strconv.Itoa(result1.Id)

	//update the rule
	input3 := UpdateRuleSetCommandInput{
		Id: id,
		Body: RuleSetView{
			Name:            "new-rule-set-test",
			SuccessCriteria: "SuccessIfAnyOneSucceeds",
			ElementType:     "Rule",
			Policy:          []int{2},
		}}
	result3, resp3, err3 := svc.Rulesets.UpdateRuleSetCommand(&input3)
	if err3 != nil {
		t.Errorf("Unable to update rule: %s", err3)
	}
	if resp3.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp3.StatusCode)
	}
	if result3.SuccessCriteria != input3.Body.SuccessCriteria {
		t.Errorf("Failed to update rule, expected: %s got: %s", result3.SuccessCriteria, input3.Body.SuccessCriteria)
	}

	//get the rule and check the update
	input4 := GetRuleSetCommandInput{
		Id: id,
	}
	result4, resp4, err4 := svc.Rulesets.GetRuleSetCommand(&input4)
	if err4 != nil {
		t.Errorf("Unable to get rule: %s", err4)
	}
	if resp4.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp4.StatusCode)
	}
	if result4.SuccessCriteria != input3.Body.SuccessCriteria {
		t.Errorf("Failed to get rule")
	}

	//delete our initial rule
	input5 := DeleteRuleSetCommandInput{
		Id: id,
	}
	resp5, err5 := svc.Rulesets.DeleteRuleSetCommand(&input5)
	if err5 != nil {
		t.Errorf("Unable to delete rule: %s", err5)
	}
	if resp5.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp5.StatusCode)
	}
}
