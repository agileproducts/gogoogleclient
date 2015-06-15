/*
Package gogoogleclient helps service apps connect to google apis
*/
package gogoogleclient

import (
  "os"
  "log"
  "fmt"
  "encoding/json"
  "net/http"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
)

// Client returns a google service client
func Client(scope string) *http.Client {
  fmt.Println(scope)
  configVariables := loadConfigurationVariables()
  jwtConfig, error := google.JWTConfigFromJSON(configVariables, scope)
  if error != nil {
    log.Fatal(error)
  }
  client := jwtConfig.Client(oauth2.NoContext)
  return client
}

// loadConfigurationVariables gets the service account info from environment variables and returns them as json
func loadConfigurationVariables() []byte {
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