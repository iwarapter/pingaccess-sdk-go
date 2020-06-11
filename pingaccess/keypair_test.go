package pingaccess

import (
	"encoding/base64"
	"io/ioutil"
	"strconv"
	"testing"
)

func TestKeyPairMethods(t *testing.T) {
	svc := config(paURL)
	content, err := ioutil.ReadFile("test_data/provider.p12")
	if err != nil {
		t.Errorf("Unable to load test p12 file")
	}
	// add a new key pair
	input1 := ImportKeyPairCommandInput{
		Body: PKCS12FileImportDocView{
			Alias:    String("test1"),
			FileData: String(base64.StdEncoding.EncodeToString(content)),
			Password: String("password"),
		}}
	result1, resp1, err1 := svc.KeyPairs.ImportKeyPairCommand(&input1)
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if result1 == nil {
		t.Errorf("Unable the marshall results")
	}

	//do a get on all key pairs
	input2 := GetKeyPairsCommandInput{}
	result2, resp2, err2 := svc.KeyPairs.GetKeyPairsCommand(&input2)
	if err2 != nil {
		t.Errorf("Unable to retrieve key pairs: %s", err2)
	}
	if resp2.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp2.StatusCode)
	}
	if result2 == nil {
		t.Errorf("Unable the marshall results")
	}

	//update the key pair
	input3 := UpdateKeyPairCommandInput{
		Id: strconv.Itoa(*result1.Id),
		Body: PKCS12FileImportDocView{
			Alias:    String("test1"),
			FileData: String(base64.StdEncoding.EncodeToString(content)),
			Password: String("password"),
		}}
	result3, resp3, err3 := svc.KeyPairs.UpdateKeyPairCommand(&input3)
	if err3 != nil {
		t.Errorf("Unable to update key pair: %s", err3)
	}
	if resp3.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp3.StatusCode)
	}
	if *result3.IssuerDn != "CN=(LOCAL) CA, OU=LOCAL, O=ORGANISATION, L=LOCALITY, ST=STATE, C=--" {
		t.Errorf("Failed to get IssuerDn: %s", *result3.IssuerDn)
	}

	//get the key pair and check the update
	input4 := GetKeyPairCommandInput{
		Id: strconv.Itoa(*result1.Id),
	}
	result4, resp4, err4 := svc.KeyPairs.GetKeyPairCommand(&input4)
	if err4 != nil {
		t.Errorf("Unable to get key pair: %s", err4)
	}
	if resp4.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp4.StatusCode)
	}
	if *result4.IssuerDn != "CN=(LOCAL) CA, OU=LOCAL, O=ORGANISATION, L=LOCALITY, ST=STATE, C=--" {
		t.Errorf("Failed to get IssuerDn: %s", *result4.IssuerDn)
	}

	//delete our initial key pair
	input5 := DeleteKeyPairCommandInput{
		Id: strconv.Itoa(*result1.Id),
	}
	resp5, err5 := svc.KeyPairs.DeleteKeyPairCommand(&input5)
	if err5 != nil {
		t.Errorf("Unable to delete key pair: %s", err5)
	}
	if resp5.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp5.StatusCode)
	}
}
