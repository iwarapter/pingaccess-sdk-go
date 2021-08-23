package pingaccess_test

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/config"
	pa "github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/version"
	"github.com/iwarapter/pingaccess-sdk-go/v62/services/virtualhosts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var paURL string

func TestMain(m *testing.M) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	paURL = "https://localhost:9000/pa-admin-api/v3"
	client := version.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL).WithDebug(false))

	log.Println("Attempting to connect to PingAccess admin API....")
	_, _, err := client.VersionCommand()
	if err != nil {
		log.Fatalf("Could not connect to PingAccess: %s", err)
	}
	log.Println("Connected to PingAccess admin API....")
	os.Exit(m.Run())
}

func TestClientParsesApiErrorViewCorrect(t *testing.T) {
	svc := virtualhosts.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL).WithDebug(false))

	result, resp, err := svc.AddVirtualHostCommand(&virtualhosts.AddVirtualHostCommandInput{Body: pa.VirtualHostView{}})
	if err == nil {
		t.Fatalf("expected err, got nil")
	}
	if strings.Contains(err.Error(), "{") || strings.Contains(err.Error(), "}") {
		t.Error("unexpected json in message")
	}
	if !strings.Contains(err.Error(), "host contains 2 validation failures:") {
		t.Error("expected error message missing")
	}
	if !strings.Contains(err.Error(), "The Host Name must not contain special characters, spaces, or the URI scheme. A wildcard hostname may be specified.") {
		t.Error("expected error message missing")
	}
	if !strings.Contains(err.Error(), "must not be null") {
		t.Error("expected error message missing")
	}
	if !strings.Contains(err.Error(), "port contains 1 validation failures:") {
		t.Error("expected error message missing")
	}
	if !strings.Contains(err.Error(), "Port must be in the range 1 to 65535") {
		t.Error("expected error message missing")
	}
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
}
