/*
Package gogoogleanalytics helps to extract data via the Google Analytics API
*/

package gogoogleanalytics

import (
  gapi "google.golang.org/api/analytics/v3"
  "net/http"
  "log"
  "io/ioutil"
)

// AccountsListHttp returns a list of accounts for a given authorised client using the raw http
func AccountsListHttp(client *http.Client) []byte {
  accounts, error := client.Get("https://www.googleapis.com/analytics/v3/management/accounts")
  if error != nil {
    log.Fatal(error)
  }
  data := responseToJson(accounts)
  return data
}

// AccountsListHttp returns a list of accounts for a given authorised client using the api
func AccountsList(client *http.Client) *gapi.Accounts{
  analyticsService, error := gapi.New(client)
  if error != nil {
    log.Fatal(error)
  } 
  function := analyticsService.Management.Accounts.List()
  accounts, error := function.Do()
  if error != nil {
    log.Fatal(error)
  } 
  return accounts
}

func responseToJson(response *http.Response) []byte {
  defer response.Body.Close()
  body, error := ioutil.ReadAll(response.Body)
  if error != nil {
    log.Fatal(error)
  }
  return body
}
