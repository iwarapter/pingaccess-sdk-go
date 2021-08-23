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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/iwarapter/pingaccess-sdk-go/v60/pingaccess"
	"github.com/iwarapter/pingaccess-sdk-go/v60/pingaccess/config"
	"github.com/iwarapter/pingaccess-sdk-go/v60/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v60/services/certificates"

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
	defer func() {
		if result1 != nil {
			_, _ = svc.DeleteTrustedCertCommand(&certificates.DeleteTrustedCertCommandInput{Id: strconv.Itoa(*result1.Id)})
		}
	}()
	require.Nil(t, err1)
	require.NotNil(t, resp1)
	assert.Equal(t, http.StatusOK, resp1.StatusCode)
	require.NotNil(t, result1)

	result2, resp2, err2 := svc.ExportTrustedCert(&certificates.ExportTrustedCertInput{Id: strconv.Itoa(*result1.Id)})
	require.Nil(t, err2)
	require.NotNil(t, resp2)
	assert.Equal(t, http.StatusOK, resp2.StatusCode)
	require.NotNil(t, result2)
	b, _ := pem.Decode([]byte(*result2))
	require.NotNil(t, b)
	_, err = x509.ParseCertificate(b.Bytes)
	assert.NoError(t, err)
}
