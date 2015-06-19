/* 
A simple example which calls the Google analytics API and lists out the properties
the given user credentials have access to
*/

package main

import (
  "fmt"
  "io/ioutil"
  "github.com/agileproducts/gogoogleclient"
)

func main() {
  client := gogoogleclient.Client("https://www.googleapis.com/auth/analytics") 
  properties, error := client.Get("https://www.googleapis.com/analytics/v3/management/accounts/~all/webproperties")
  if error != nil {
    fmt.Println(error)
  }
  defer properties.Body.Close()
  body, err := ioutil.ReadAll(properties.Body)
  if err != nil {
    fmt.Print(error)
  }
  fmt.Println(string(body))
}

