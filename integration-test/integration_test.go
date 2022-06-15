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

// run the tests
// step 1: docker compose up -d
// step 2: go clean -testcache | go test -v ./integration-test/...
// step 3: view the result in the console
// step 4: docker-compose down

const (
	// Attempts connection
	host       = "localhost:80"
	healthPath = "http://" + host + "/healthz"
	attempts   = 20
	id         = "1a6db844-2426-42a9-8770-2302714bcc2c" //<-- test guid

	// HTTP REST
	basePath        = "http://" + host + "/v1"
	KeycloakService = "https://keycloak.krekels-server.com/auth/realms/micflix/protocol/openid-connect/token"
	testAuthSecret  = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4t/mew6nUASz0MyRlDKKXE7YcPd90hR8Thao315S6S//diWxRPJnPb8InrfKHvncPGYjevCyah9XPN8cxwEP74f7FauXLkvDRsXomAYl17drW/54fQmzj3ZCIQByWkdv7SnAHHvjFzp1pTRGJF1OGPZyqDHuS0AttWUbOaaECm8tT3qzFlQYzvyhCWSRrZ1MJL1oYV4UqQCMcenKlL8R0zeRql/XuJ20AkkFio1UzDcTdHu5OSoQ16kB9nxi2QkCdeQz7le2x63OLoHUS7MpCCV0PwDmDNyj1VFS/sukAc5RFF5XF2lYa85GO+zCIEFe7/d6PS+cY9mVwLvOzJC9dQIDAQAB"
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

func TestGetAllUsers_NoToken(t *testing.T) {
	Do(Get(basePath+"/users"),
		Send().Headers("authorization").Add("Bearer"),
		Expect().Status().Equal(http.StatusUnauthorized),
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

func TestGetUserById_NoID(t *testing.T) {
	TestCreateUser(t) //whitout adding the user, the test will fail
	Do(Get(basePath+"/users/"),
		Send().Headers("authorization").Add("Bearer "+authToken),
		Expect().Status().Equal(http.StatusBadRequest),
	)
}

func TestGetUser_NoToken(t *testing.T) {
	TestCreateUser(t) //whitout adding the user, the test will fail
	Do(Get(basePath+"/users/"+id),
		Send().Headers("authorization").Add("Bearer "),
		Expect().Status().Equal(http.StatusUnauthorized),
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

func TestGetAllProfilesById_NoToken(t *testing.T) {
	TestCreateUser(t) //whitout adding the user, the test will fail
	Do(Get(basePath+"/users/profile/"+id),
		Send().Headers("authorization").Add("Bearer "),
		Expect().Status().Equal(http.StatusUnauthorized),
	)
}
func TestGetAllProfilesById_NoID(t *testing.T) {
	TestCreateUser(t) //whitout adding the user, the test will fail
	Do(Get(basePath+"/users/profile/"),
		Send().Headers("authorization").Add("Bearer "+authToken),
		Expect().Status().Equal(http.StatusBadRequest),
	)
}

// Http POST: Create
func TestCreateUser(t *testing.T) {
	Do(Post(basePath+"/user/"+id),
		Send().Headers("authorization").Add("Bearer "+authToken),
		Expect().Status().Equal(http.StatusCreated),
	)
}

// Http POST: Create
func TestCreateUser_NoID(t *testing.T) {
	Do(Post(basePath+"/user/"),
		Send().Headers("authorization").Add("Bearer "+authToken),
		Expect().Status().Equal(http.StatusBadRequest),
	)
}

// Http POST: Create
func TestCreateUser_NoToken(t *testing.T) {
	Do(Post(basePath+"/user/"+id),
		Send().Headers("authorization").Add("Bearer "),
		Expect().Status().Equal(http.StatusUnauthorized),
	)
}
