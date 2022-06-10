package integration_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	. "github.com/Eun/go-hit"
	"github.com/Netflix-Clone-MicFlix/User-Service/internal"
	"github.com/Netflix-Clone-MicFlix/User-Service/internal/repositories"
	"github.com/Netflix-Clone-MicFlix/User-Service/pkg/mongodb"
)

const (
	// Attempts connection
	host       = "micflix-api-gateway.com:8080/user-service"
	healthPath = "http://" + host + "/healthz"
	attempts   = 20
	id         = "8c665ed0-72dc-490f-9968-afca1d087191" //<-- test guid

	// HTTP REST
	basePath        = "http://" + host + "/v1"
	KeycloakService = "http://micflix-keycloak.com:8080/auth/realms/micflix/protocol/openid-connect/token"

	testAuthSecret = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAiMoKmSYAn0yOooKktMQT8rC6gRlcic3HpvhnNk7X/u8n3GfdlyzVAOWhzGDE2MPm/tIduR8g/qX4ZVcFy2Vf9bIf4GdMYndITBnumloQkH8D+yqdwhlxSsjxwLLAvhHWEXUnigGgfMu6ylf325yLnAsYpNPkvuIb191Mn4vTGiq4qKq0/+kOVvzYLO2x5WwIDhd5DeyAFSmCOCWb7qVZ7crH17IMeWgvAo/K7coZFI4TNQl6c6J6gpqFfhzFolIdNXZr+asdgKwA2beaCRjmc9wgDkRaA9o56P69ZcVAotiboQFkhlRQf1bMd3tPgQIdM0gu5y8XljU5S/5xwAQWtwIDAQAB"
)

var authToken string
var UserRepo internal.UserRepo
var ProfileRepo internal.ProfileRepo
var MovieTagRepo internal.MovieTagRepo

func TestMain(m *testing.M) {

	err := healthCheck(attempts)
	if err != nil {
		log.Fatalf("Integration tests: host %s is not available: %s", host, err)
	}

	mdb, err := mongodb.New("IntegrationTest", "IntegrationTest", "localhost", "Users")
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - Mongodb.New: %w", err))
	}
	UserRepo = repositories.NewUserRepo(mdb)
	ProfileRepo = repositories.NewProfileRepo(mdb)
	MovieTagRepo = repositories.NewMovieTagRepo(mdb)

	authToken, err = getToken()
	if err != nil || authToken == "" {
		log.Fatal("Integration tests: cannot get token", err)
	}

	log.Printf("Integration tests: host %s is available", host)

	code := m.Run()
	os.Exit(code)
}

func healthCheck(attempts int) error {
	var err error

	for attempts > 0 {
		err = Do(Get(healthPath), Expect().Status().Equal(http.StatusOK))
		if err == nil {
			return nil
		}

		log.Printf("Integration tests: url %s is not available, attempts left: %d", healthPath, attempts)

		time.Sleep(time.Second)

		attempts--
	}

	return err
}

func getToken() (string, error) {
	// get the access_token from the response
	form := url.Values{}
	form.Add("grant_type", "password")
	form.Add("username", "IntegrationTest")
	form.Add("password", "IntegrationTest")
	form.Add("client_id", "micflix-angular-client")
	form.Add("client_secret", "")

	resp, err := http.PostForm(KeycloakService, form)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// get the access_token from the response
	var token struct {
		AccessToken string `json:"access_token"`
	}
	err = json.Unmarshal(body, &token)
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

// Http GET: GetAllUsers
func TestGetAllUsers(t *testing.T) {
	Do(Get(basePath+"/users"),
		Send().Headers("authorization").Add("Bearer "+authToken),
		Expect().Status().Equal(http.StatusOK),
	)
}

// Http GET: GetById
func TestGetUserById(t *testing.T) {
	TestCreateUser(t) //whitout adding the user, the test will fail
	Do(Get(basePath+"/users/"+id),
		Send().Headers("authorization").Add("Bearer "+authToken),
		Expect().Status().Equal(http.StatusOK),
	)
}

// Http GET: GetAllProfilesById
func TestGetAllProfilesById(t *testing.T) {
	TestCreateUser(t) //whitout adding the user, the test will fail
	Do(Get(basePath+"/users/profile/"+id),
		Send().Headers("authorization").Add("Bearer "+authToken),
		Expect().Status().Equal(http.StatusOK),
	)
}

// Http POST: Create
func TestCreateUser(t *testing.T) {
	Do(Post(basePath+"/user/"+id),
		Send().Headers("authorization").Add("Bearer "+authToken),
		Expect().Status().Equal(http.StatusCreated),
	)
}
