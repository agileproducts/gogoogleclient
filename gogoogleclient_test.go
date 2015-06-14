package gogoogleclient

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "encoding/json"
)

func TestLoadConfigurationVariables(test *testing.T) {
  expected := "boo"
  jsonconfig := LoadConfigurationVariables()
  var actual map[string]interface{}
  json.Unmarshal(jsonconfig, &actual)
  assert.Equal(test,expected,actual,"It should return ba")
}