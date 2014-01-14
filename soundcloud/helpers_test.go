package soundcloud

import (
	"fmt"
	"testing"
)

var doAuthorizedRequests bool
var api *Api
var ladygaga_id uint64 = 55213175
var thriftshop_id uint64 = 57886117
var macklemore_id uint64 = 557633
var joneisen_id uint64 = 557633

// -- helpers --

func init() {
	doAuthorizedRequests = (TestConfig["access_token"] != "")
	if !doAuthorizedRequests {
		fmt.Println("*** Authorized requests will not performed because no access_token was specified in config_test.go")
	}
	api = createApi()
}

func authorizedRequest(t *testing.T) {
	if !doAuthorizedRequests {
		t.Skip("Access Token not provided.")
	}
}

func createApi() *Api {
	api := &Api{
		ClientId:     TestConfig["client_id"].(string),
		ClientSecret: TestConfig["client_secret"].(string),
		AccessToken:  TestConfig["access_token"].(string),
		RefreshToken: TestConfig["refresh_token"].(string),
	}

	_ = api.Refresh()
	return api
}
