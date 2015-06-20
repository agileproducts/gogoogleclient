/*
Package gogoogleclient helps service apps connect to google apis
*/
package gogoogleclient

import (
  "os"
  "log"
  "io/ioutil"
  "net/http"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
)

// Client returns a google service client
func Client(scope string) *http.Client {
  configVariables := loadApplicationCredentials()
  jwtConfig, error := google.JWTConfigFromJSON(configVariables, scope)
  if error != nil {
    log.Fatal(error)
  }
  client := jwtConfig.Client(oauth2.NoContext)
  return client
}

// loadConfigurationVariables gets the service account info from environment variables and returns them as json
func loadApplicationCredentials() []byte {
  credentialFile,error := ioutil.ReadFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
  if error != nil {
    log.Fatal(error)
  }
  return credentialFile
}
