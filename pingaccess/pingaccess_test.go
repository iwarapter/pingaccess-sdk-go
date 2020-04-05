package pingaccess

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/ory/dockertest"
)

var paURL *url.URL

func TestMain(m *testing.M) {
	log.Println("Initializing docker container")
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	devOpsUser, devOpsUserExists := os.LookupEnv("PING_IDENTITY_DEVOPS_USER")
	devOpsKey, devOpsKeyExists := os.LookupEnv("PING_IDENTITY_DEVOPS_KEY")

	var tag = "6.0.1-edge"
	var options *dockertest.RunOptions

	if devOpsUserExists && devOpsKeyExists {
		options = &dockertest.RunOptions{
			Repository: "pingidentity/pingaccess",
			Env:        []string{"PING_IDENTITY_ACCEPT_EULA=YES",fmt.Sprintf("PING_IDENTITY_DEVOPS_USER=%s", devOpsUser), fmt.Sprintf("PING_IDENTITY_DEVOPS_KEY=%s", devOpsKey)},
			Tag:        tag,
		}
	} else {
		dir, _ := os.Getwd()
		options = &dockertest.RunOptions{
			Repository: "pingidentity/pingaccess",
			Env: []string{"PING_IDENTITY_ACCEPT_EULA=YES"},
			Mounts:     []string{dir + "/pingaccess.lic:/opt/in/instance/conf/pingaccess.lic"},
			Tag:        tag,
		}
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(options)
	resource.Expire(60)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	pool.MaxWait = time.Minute * 2
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		paURL, _ = url.Parse(fmt.Sprintf("https://localhost:%s", resource.GetPort("9000/tcp")))
		client := NewClient("administrator", "2FederateM0re", paURL, "/pa-admin-api/v3", nil)

		log.Println("Attempting to connect to PingAccess admin API....")
		_, _, err = client.Version.VersionCommand()
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	log.Println("Connected to PingAccess admin API....")
	code := m.Run()

	log.Println("Tests complete shutting down container")
	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func config(url *url.URL) *Client {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	return NewClient("administrator", "2FederateM0re", url, "/pa-admin-api/v3", nil)
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
