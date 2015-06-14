/*
Package gogoogleclient helps service apps connect to google apis
*/
package gogoogleclient

import (
  "os"
  "log"
  "encoding/json"
  _ "golang.org/x/oauth2"
  _ "golang.org/x/oauth2/google"
)


// LoadConfigurationVariables gets the service account info from environment variables and returns them as json
func LoadConfigurationVariables() []byte {
  keyMap := map[string]string {
    "Email" : os.Getenv("GOOGLE_CLIENT_EMAIL"),
    "PrivateKey" : os.Getenv("GOOGLE_PRIVATE_KEY"),
  }
  for key, value := range keyMap {
    if value == "" {
      log.Fatal(key + " environment variable not set")
    }
  }
  jsonMap,err := json.Marshal(keyMap)
  if err != nil {
    log.Fatal(err)
  }
  return jsonMap
}