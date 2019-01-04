package pingaccess

import (
	"net/url"
	"testing"
)

func TestRuleDescriptors(t *testing.T) {
	url, _ := url.Parse("https://localhost:9000")
	svc := config(url)

	result1, resp1, err1 := svc.Rules.GetRuleDescriptorsCommand()
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if len(result1.Items) == 0 {
		t.Errorf("Unable the marshall results")
	}

	input2 := GetRuleDescriptorCommandInput{
		RuleType: "authnreq",
	}
	result2, resp2, err2 := svc.Rules.GetRuleDescriptorCommand(&input2)
	if err2 != nil {
		t.Errorf("Unable to execute command: %s", err2.Error())
	}
	if resp2.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp2.StatusCode)
	}
	if *result2.ClassName != "com.pingidentity.pa.policy.authnreq.AuthenticationRequirementsPolicyInterceptor" {
		t.Errorf("Unable the marshall results")
	}
}

func TestRuleMethods(t *testing.T) {
	url, _ := url.Parse("https://localhost:9000")
	svc := config(url)

	// add a new rule
	input1 := AddRuleCommandInput{
		Body: RuleView{
			ClassName:             String("com.pingidentity.pa.policy.CIDRPolicyInterceptor"),
			Name:                  String("woottest"),
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
		}}

	result1, resp1, err1 := svc.Rules.AddRuleCommand(&input1)
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
	input2 := GetRulesCommandInput{}
	result2, resp2, err2 := svc.Rules.GetRulesCommand(&input2)
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
	input3 := UpdateRuleCommandInput{
		Id: result1.Id.String(),
		Body: RuleView{
			ClassName:             String("com.pingidentity.pa.policy.CIDRPolicyInterceptor"),
			Name:                  String("woottest"),
			SupportedDestinations: &[]*string{String("localhost:1234")},
			Configuration: map[string]interface{}{
				"cidrNotation":              "127.0.0.1/32",
				"negate":                    false,
				"overrideIpSource":          false,
				"headers":                   []interface{}{},
				"headerValueLocation":       "LAST",
				"fallbackToLastHopIp":       true,
				"errorResponseCode":         403,
				"errorResponseStatusMsg":    "Not Allowed",
				"errorResponseTemplateFile": "policy.error.page.template.html",
				"errorResponseContentType":  "text/html;charset=UTF-8",
			},
		}}
	result3, resp3, err3 := svc.Rules.UpdateRuleCommand(&input3)
	if err3 != nil {
		t.Errorf("Unable to update rule: %s", err3)
	}
	if resp3.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp3.StatusCode)
	}
	if result3.Configuration["errorResponseStatusMsg"] != input3.Body.Configuration["errorResponseStatusMsg"] {
		t.Errorf("Failed to update rule, expected: %s got: %s", result3.Configuration["errorResponseStatusMsg"], input3.Body.Configuration["errorResponseStatusMsg"])
	}

	//get the rule and check the update
	input4 := GetRuleCommandInput{
		Id: result1.Id.String(),
	}
	result4, resp4, err4 := svc.Rules.GetRuleCommand(&input4)
	if err4 != nil {
		t.Errorf("Unable to get rule: %s", err4)
	}
	if resp4.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp4.StatusCode)
	}
	if result4.Configuration["errorResponseStatusMsg"] != input3.Body.Configuration["errorResponseStatusMsg"] {
		t.Errorf("Failed to get rule")
	}

	//delete our initial rule
	input5 := DeleteRuleCommandInput{
		Id: result1.Id.String(),
	}
	resp5, err5 := svc.Rules.DeleteRuleCommand(&input5)
	if err5 != nil {
		t.Errorf("Unable to delete rule: %s", err5)
	}
	if resp5.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp5.StatusCode)
	}
}
