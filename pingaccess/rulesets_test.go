package pingaccess

import (
	"testing"
)

func TestRuleSetElementTypes(t *testing.T) {
	svc := config(paURL)

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
	svc := config(paURL)

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
	svc := config(paURL)

	demoRule := AddRuleCommandInput{
		Body: RuleView{
			ClassName:             String("com.pingidentity.pa.policy.CIDRPolicyInterceptor"),
			Name:                  String("another_test"),
			SupportedDestinations: &[]*string{String("localhost:1234")},
			Configuration: map[string]interface{}{
				"cidrNotation":              "127.0.0.1/32",
				"negate":                    false,
				"overrideIpSource":          false,
				"headers":                   []interface{}{},
				"headerValueLocation":       "LAST",
				"fallbackToLastHopIp":       true,
				"errorResponseCode":         403,
				"errorResponseStatusMsg":    "Forbidden",
				"errorResponseTemplateFile": "policy.error.page.template.html",
				"errorResponseContentType":  "text/html;charset=UTF-8",
			},
		},
	}

	ruleResult, ruleResp, ruleErr := svc.Rules.AddRuleCommand(&demoRule)
	if ruleErr != nil {
		t.Errorf("Unable to execute command: %s", ruleErr.Error())
	}
	if ruleResp == nil || ruleResp.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", ruleResp.StatusCode)
	}
	if ruleResult == nil {
		t.Fatalf("Unable the marshall results")
	}
	id, _ := ruleResult.Id.Int64()

	input1 := AddRuleSetCommandInput{
		Body: RuleSetView{
			Name:            String("new-rule-set-test"),
			SuccessCriteria: String("SuccessIfAllSucceed"),
			ElementType:     String("Rule"),
			Policy:          &[]*int{Int(int(id))},
		}}

	result1, resp1, err1 := svc.Rulesets.AddRuleSetCommand(&input1)
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1 == nil || resp1.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if result1 == nil {
		t.Fatalf("Unable the marshall results")
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

	//update the rule
	input3 := UpdateRuleSetCommandInput{
		Id: result1.Id.String(),
		Body: RuleSetView{
			Name:            String("new-rule-set-test"),
			SuccessCriteria: String("SuccessIfAnyOneSucceeds"),
			ElementType:     String("Rule"),
			Policy:          &[]*int{Int(int(id))},
		}}
	result3, resp3, err3 := svc.Rulesets.UpdateRuleSetCommand(&input3)
	if err3 != nil {
		t.Errorf("Unable to update rule: %s", err3)
	}
	if resp3.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp3.StatusCode)
	}
	if *result3.SuccessCriteria != *input3.Body.SuccessCriteria {
		t.Errorf("Failed to update rule, expected: %s got: %s", *result3.SuccessCriteria, *input3.Body.SuccessCriteria)
	}

	//get the rule and check the update
	input4 := GetRuleSetCommandInput{
		Id: result1.Id.String(),
	}
	result4, resp4, err4 := svc.Rulesets.GetRuleSetCommand(&input4)
	if err4 != nil {
		t.Errorf("Unable to get rule: %s", err4)
	}
	if resp4.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp4.StatusCode)
	}
	if *result4.SuccessCriteria != *input3.Body.SuccessCriteria {
		t.Errorf("Failed to get rule")
	}

	//delete our initial rule
	input5 := DeleteRuleSetCommandInput{
		Id: result1.Id.String(),
	}
	resp5, err5 := svc.Rulesets.DeleteRuleSetCommand(&input5)
	if err5 != nil {
		t.Errorf("Unable to delete rule: %s", err5)
	}
	if resp5.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp5.StatusCode)
	}
}
