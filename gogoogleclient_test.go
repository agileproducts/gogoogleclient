package gogoogleclient

import (
	"github.com/stretchr/testify/assert"
	_ "net/http"
	"testing"
)

func TestClient(test *testing.T) {
	client := Client("https://www.googleapis.com/auth/analytics")
	//assert.IsType(test, client, *http.Client, "It should return an http client for a valid scope")
	assert.NotNil(test, client, "It should return an http client for a valid scope")
}

func TestLoadApplicationCredentials(test *testing.T) {
	jsonconfig := loadApplicationCredentials()
	assert.NotNil(test, jsonconfig, "It should extract the relevant environment variables as JSON")
}
