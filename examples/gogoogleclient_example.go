package main

import (
  "fmt"
  "github.com/agileproducts/gogoogleclient"
)

func main() {
  client := gogoogleclient.Client("https://www.googleapis.com/auth/analytics") 
  properties, error := client.Get("https://www.googleapis.com/analytics/v3/management/accounts/~all/webproperties")
  if error != nil {
    fmt.Println(error)
  }
  fmt.Println(properties)
}

