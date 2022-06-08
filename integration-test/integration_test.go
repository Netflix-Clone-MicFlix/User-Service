package integration_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	. "github.com/Eun/go-hit"
	"github.com/Netflix-Clone-MicFlix/User-Service/internal"
)

const (
	// Attempts connection
	host = "localhost:8080/user-service"
	// pgHost     = "postgres://user:pass@survey-service_postgres:5432/survey-service"
	healthPath = "http://" + host + "/healthz"
	attempts   = 20

	// HTTP REST
	basePath        = "http://" + host + "/v1"
	KeycloakService = "http://localhost:8080/auth/realms/micflix/protocol/openid-connect/token"

	testAuthSecret = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAkL+KfUTxVH+R4hG9jA5UR8/kqH2rAAtCpmDvcMJkDDZnFYM790fz/kVB3TlMP5oUChqTFY/dMGhtKjZ+JsC5r/pK6m5x1OX8MqsSLrfUL8Xkp9v1CCV0nVCNrAvVzET8t7UL4jXbADw9zrkwk9fsXdQLoY8ZTDVwOCtoNvWO0D7DyKY08SPxWBDqLPpPojBd5gXMqx33M+IY71801bsP8k7B7UjjkOina98jkKdBEyOhvt52b/t9TCoEPcnzsbOeDHG6C25Dx/azF70DtM4DKKa9wvWrFyZS8z45GwClhSCLb6ifPkrS4tanJ/gK85nT8cIcw7H2wWxqT4ZWlGmvwwIDAQAB"
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
	// cfg, err := config.NewIntergrationTestConfig("../config/config.yml")
	// if err != nil {
	// 	log.Fatalf("Config error: %s", err)
	// }
	// // Repository mongodb
	// mdb, err := mongodb.New(cfg.MDB.Username, cfg.MDB.Password, cfg.MDB.Cluster, cfg.MDB.Database)
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	// }
	// UserRepo = repositories.NewUserRepo(mdb)
	// ProfileRepo = repositories.NewProfileRepo(mdb)
	// MovieTagRepo = repositories.NewMovieTagRepo(mdb)

	authToken, err = getToken()
	if err != nil {
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
