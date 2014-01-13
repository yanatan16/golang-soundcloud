package soundcloud

// Copy this to config_test.go (its ignored by git)
// And fill in your test clientID, clientSecret, authorization token and secret
// You can get the client information from http://developers.soundcloud.com/you/apps
// And you'll have to authorize with that client to get a token
// You don't NEED access tokens per se but some requests require it.

var TestConfig map[string]string = map[string]string{
	"client_id":    "f6212fc8e1e90501e7cece44f78b8f43",
	"access_token": "1-64086-73163832-9119bb3fe59014f",
	"my_id":        "73163832",
}
