package pingaccess_test

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"log"
	"math/big"
	"time"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/config"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/keyPairs"

	"testing"
)

func TestKeyPairsMethods(t *testing.T) {
	svc := keyPairs.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL).WithDebug(false))

	// add a new site authenticator
	input1 := keyPairs.GenerateCsrCommandInput{Id: "1"}
	result1, resp1, err1 := svc.GenerateCsrCommand(&input1)
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1 == nil || resp1.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if result1 == nil {
		t.Fatalf("Unable the marshall results")
	}
	b, _ := pem.Decode([]byte(*result1))
	if _, err := x509.ParseCertificateRequest(b.Bytes); err != nil {
		t.Fatalf("response is not a valid CSR: %s", *result1)
	}

	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization: []string{"Testing"},
			Country:      []string{"GB"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &key.PublicKey, key)
	if err != nil {
		log.Fatalf("Failed to create certificateCa: %s", err)
	}
	caBuf := new(bytes.Buffer)
	_ = pem.Encode(caBuf, &pem.Block{Type: "CERTIFICATE", Bytes: caBytes})

	//*csrPem = strings.ReplaceAll(*csrPem, " NEW ", " ")
	//b, _ := pem.Decode([]byte(*csrPem))
	csr, err := x509.ParseCertificateRequest(b.Bytes)
	if err != nil {
		t.Fatalf("unable to parse csr: %s", err)
	}
	template := x509.Certificate{
		Signature:          csr.Signature,
		SignatureAlgorithm: csr.SignatureAlgorithm,

		PublicKeyAlgorithm: csr.PublicKeyAlgorithm,
		PublicKey:          csr.PublicKey,

		SerialNumber: big.NewInt(2),
		Issuer:       ca.Subject,
		Subject:      csr.Subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(24 * time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	certBytes, err := x509.CreateCertificate(rand.Reader, &template, ca, csr.PublicKey, key)
	if err != nil {
		t.Fatalf("unable to sign certificate request: %s", err)
	}
	buf := new(bytes.Buffer)
	err = pem.Encode(buf, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	if err != nil {
		t.Fatalf("unable to encode certificate: %s", err)
	}
	signedCert := buf.String()

	input2 := keyPairs.ImportCSRResponseCommandInput{
		Id: "1",
		Body: models.CSRResponseImportDocView{
			FileData: pingaccess.String(base64.StdEncoding.EncodeToString([]byte(signedCert))),
		},
	}
	result2, resp2, err2 := svc.ImportCSRResponseCommand(&input2)
	if err2 != nil {
		t.Errorf("Unable to execute command: %s", err2.Error())
	}
	if resp2 == nil || resp2.StatusCode != 200 {
		t.Errorf("Invalid response code: %d", resp2.StatusCode)
	}
	if result2 == nil {
		t.Fatalf("Unable the marshall results")
	}

}
