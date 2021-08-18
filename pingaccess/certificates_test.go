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
	"net/http"
	"strconv"
	"time"

	"github.com/iwarapter/pingaccess-sdk-go/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/config"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/certificates"

	"testing"
)

func TestCertificatesMethods(t *testing.T) {
	svc := certificates.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL).WithDebug(false))

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
	signedCert := caBuf.String()

	input2 := certificates.ImportTrustedCertInput{
		Body: models.X509FileImportDocView{
			FileData: pingaccess.String(base64.StdEncoding.EncodeToString([]byte(signedCert))),
		},
	}
	result1, resp1, err1 := svc.ImportTrustedCert(&input2)
	if err1 != nil {
		t.Errorf("Unable to execute command: %s", err1.Error())
	}
	if resp1 == nil || resp1.StatusCode != http.StatusOK {
		t.Errorf("Invalid response code: %d", resp1.StatusCode)
	}
	if result1 == nil {
		t.Fatalf("Unable the marshall results")
	}

	result2, resp2, err2 := svc.ExportTrustedCert(&certificates.ExportTrustedCertInput{Id: strconv.Itoa(*result1.Id)})
	if err2 != nil {
		t.Errorf("Unable to execute command: %s", err2.Error())
	}
	if resp2 == nil || resp2.StatusCode != http.StatusOK {
		t.Errorf("Invalid response code: %d", resp2.StatusCode)
	}
	if result2 == nil {
		t.Fatalf("Unable the marshall results")
	}
	b, _ := pem.Decode([]byte(*result2))
	if _, err := x509.ParseCertificate(b.Bytes); err != nil {
		t.Fatalf("response is not a valid certificate: %s", *result2)
	}
}
