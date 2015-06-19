# GoGoogleClient

A golang package to help with creating service accounts to google api services.

## Environment

The authentication information needed to connect to a Google service account should be downloaded as a json file from the Google 
[developer console](https://developers.google.com/console).

The location of this file should then be stored in the environment variable `GOOGLE_APPLICATION_CREDENTIALS`. 

## Usage

The Client method returns an authenticated http client for the given API scope:

```
client := gogoogleclient.Client("https://www.googleapis.com/auth/analytics") 
data, error := client.Get(appropriate http request to API)

```

## Dependencies

* golang.org/x/oauth2
* golang.org/x/oauth2/google
* github.com/stretchr/testify/assert