package gogoogleanalytics

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/agileproducts/gogoogleclient"
  "fmt"
  "encoding/json"
)

func TestAccountsList(test *testing.T) {
  client := gogoogleclient.Client("https://www.googleapis.com/auth/analytics")
  actual := AccountsList(client)
  blah, _ := json.Marshal(actual)
  fmt.Println(string(blah))
  assert.NotNil(test, actual, "It should return a list of accounts")
}



