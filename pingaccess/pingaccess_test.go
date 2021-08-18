package pingaccess_test

import (
	"crypto/tls"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/config"
	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/iwarapter/pingaccess-sdk-go/services/version"
	"github.com/iwarapter/pingaccess-sdk-go/services/virtualhosts"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

var paURL string

func TestMain(m *testing.M) {
	log.Println("Initializing docker container")
	//pool, err := dockertest.NewPool("")
	//if err != nil {
	//	log.Fatalf("Could not connect to docker: %s", err)
	//}
	//
	//devOpsUser, devOpsUserExists := os.LookupEnv("PING_IDENTITY_DEVOPS_USER")
	//devOpsKey, devOpsKeyExists := os.LookupEnv("PING_IDENTITY_DEVOPS_KEY")
	//
	//var tag = "6.0.2-edge"
	//var options *dockertest.RunOptions
	//
	//if devOpsUserExists && devOpsKeyExists {
	//	options = &dockertest.RunOptions{
	//		Repository: "pingidentity/pingaccess",
	//		Env:        []string{"PING_IDENTITY_ACCEPT_EULA=YES", fmt.Sprintf("PING_IDENTITY_DEVOPS_USER=%s", devOpsUser), fmt.Sprintf("PING_IDENTITY_DEVOPS_KEY=%s", devOpsKey)},
	//		Tag:        tag,
	//	}
	//} else {
	//	dir, _ := os.Getwd()
	//	options = &dockertest.RunOptions{
	//		Repository: "pingidentity/pingaccess",
	//		Env:        []string{"PING_IDENTITY_ACCEPT_EULA=YES"},
	//		Mounts:     []string{dir + "/pingaccess.lic:/opt/in/instance/conf/pingaccess.lic"},
	//		Tag:        tag,
	//	}
	//}
	//
	//// pulls an image, creates a container based on it and runs it
	//resource, err := pool.RunWithOptions(options)
	//resource.Expire(60)
	//if err != nil {
	//	log.Fatalf("Could not start resource: %s", err)
	//}
	//
	//pool.MaxWait = time.Minute * 2
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	//if err := pool.Retry(func() error {
	//	var err error
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	paURL = "https://localhost:9000/pa-admin-api/v3"
	client := version.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Access").WithEndpoint(paURL).WithDebug(false))

	log.Println("Attempting to connect to PingAccess admin API....")
	_, _, err := client.VersionCommand()
	if err != nil {
		log.Fatalf("Could not connect to PingAccess: %s", err)
	}

	log.Println("Connected to PingAccess admin API....")
	code := m.Run()
	//
	//log.Println("Tests complete shutting down container")
	//// You can't defer this because os.Exit doesn't care for defer
	//if err := pool.Purge(resource); err != nil {
	//	log.Fatalf("Could not purge resource: %s", err)
	//}

	os.Exit(code)
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
	if result != nil {
		t.Errorf("expected no response, got: %v", result)
	}
	if resp == nil {
		t.Fatalf("response body is nil")
	}
	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("expected response code %d, got %d", http.StatusUnprocessableEntity, resp.StatusCode)
	}
}
