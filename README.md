# golang-soundcloud

A [soundcloud](http://soundcloud.com) [API](http://developers.soundcloud.com) wrapper.

## Features

Implemented:

- Some GET requests are implemented
- Both authenticated and unauthenticated requests can be made
- No `interface{}` data types!

Todo:

- The full set of GET requests
- Authentication
- POST / DELETE requests

## Contributing

This API wrapper is a good start, but is no where near completion. If you're using Go and SoundCloud, please consider contributing by implementing more methods (see [Issues](https://github.com/yanatan16/golang-soundcloud/issues)) and making a pull request. As for style, just use `go fmt` before you pull!

## Documentation

[Documentation on godoc.org](http://godoc.org/github.com/yanatan16/golang-soundcloud/soundcloud)

## Install

```
go get github.com/yanatan16/golang-soundcloud/soundcloud
```

## Creation

```go
import (
  "github.com/yanatan16/golang-soundcloud/soundcloud"
)

unauthenticatedApi := &soundcloud.Api{
  ClientId: "my-client-id",
}
// Or
unauthenticatedApi2 := soundcloud.New("my-client-id","")

authenticatedApi := &soundcloud.Api{
  AccessToken: "my-access-token",
}
// or
authenticatedApi2 := soundcloud.New("", "my-access-token")
```

## Usage

```go
import (
  "fmt"
  "github.com/yanatan16/golang-soundcloud/soundcloud"
  "net/url"
)

func DoSomeSoundCloudApiStuff(accessToken string) {
  api := New("", accessToken)

  var myId string

  // Get yourself!
  if me, err := api.Me().Get(); err != nil {
    panic(err)
  } else {
    fmt.Printf("My userid is %s, username is %s, and I have %d followers\n", me.Id, me.Username, me.FollowerCount)
  }
}
```

## Tests

To run the tests, you'll need at least a `ClientId` (which you can get from [here](http://soundcloud.com/developer/clients/manage/)), and preferably an authenticated users' `AccessToken`, which is a bit harder to get (involves authenticating for an app and getting that auth token)

First, fill in `config_test.go.example` and save it as `config_test.go`. Then run `go test`

## Notes

- Certain methods require an access token so check the official documentation before using an unauthenticated `Api`. This package will use it if it is given.

## License

MIT-style. See LICENSE file.