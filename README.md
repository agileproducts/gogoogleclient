# GoGoogleClient

A golang package to help with creating service accounts to google api services.

## Environment

The authentication information needed to connect to a Google service account should be stored in two environment variables

```
GOOGLE_CLIENT_EMAIL
GOOGLE_PRIVATE_KEY
```

These can be taken from the corresponding values in the json file which can be downloaded from the Google developer console.

## Dependencies

* golang.org/x/oauth2
* golang.org/x/oauth2/google
* github.com/stretchr/testify/assert