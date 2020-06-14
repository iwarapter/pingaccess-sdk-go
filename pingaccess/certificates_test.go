package pingaccess

import (
	"encoding/base64"
	"io/ioutil"
	"strconv"
	"testing"
)

func TestCertificateMethods(t *testing.T) {
	svc := config(paURL)
	svc.LogDebug = true
	content, err := ioutil.ReadFile("test_data/amazon_root_ca1.pem")
	if err != nil {
		t.Errorf("Unable to load test certificate file")
	}
	// add a new identity mapping
	input1 := ImportTrustedCertInput{
		Body: X509FileImportDocView{
			Alias:    String("test1"),
			FileData: String(base64.StdEncoding.EncodeToString(content)),
		}}
	result1, resp1, err1 := svc.Certificates.ImportTrustedCert(&input1)
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1 == nil || resp1.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if result1 == nil {
		t.Fatalf("Unable the marshall results")
	}

	//do a get on all trusted certificates
	input2 := GetTrustedCertsInput{}
	result2, resp2, err2 := svc.Certificates.GetTrustedCerts(&input2)
	if err2 != nil {
		t.Errorf("Unable to retrieve trusted certificates: %s", err2)
	}
	if resp2.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp2.StatusCode)
	}
	if result2 == nil {
		t.Errorf("Unable the marshall results")
	}

	//update the identity mapping
	input3 := UpdateTrustedCertInput{
		Id: strconv.Itoa(*result1.Id),
		Body: X509FileImportDocView{
			Alias:    String("test1"),
			FileData: String(base64.StdEncoding.EncodeToString(content)),
		}}
	result3, resp3, err3 := svc.Certificates.UpdateTrustedCert(&input3)
	if err3 != nil {
		t.Errorf("Unable to update identity mapping: %s", err3)
	}
	if resp3.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp3.StatusCode)
	}
	if *result3.IssuerDn != "CN=Amazon Root CA 1, O=Amazon, C=US" {
		t.Errorf("Failed to get IssuerDn: %s", *result3.IssuerDn)
	}

	//get the identity mapping and check the update
	input4 := GetTrustedCertInput{
		Id: strconv.Itoa(*result1.Id),
	}
	result4, resp4, err4 := svc.Certificates.GetTrustedCert(&input4)
	if err4 != nil {
		t.Errorf("Unable to get identity mapping: %s", err4)
	}
	if resp4.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp4.StatusCode)
	}
	if *result4.IssuerDn != "CN=Amazon Root CA 1, O=Amazon, C=US" {
		t.Errorf("Failed to get IssuerDn: %s", *result4.IssuerDn)
	}

	//delete our initial identity mapping
	input5 := DeleteTrustedCertCommandInput{
		Id: strconv.Itoa(*result1.Id),
	}
	resp5, err5 := svc.Certificates.DeleteTrustedCertCommand(&input5)
	if err5 != nil {
		t.Errorf("Unable to delete identity mapping: %s", err5)
	}
	if resp5.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp5.StatusCode)
	}
}
