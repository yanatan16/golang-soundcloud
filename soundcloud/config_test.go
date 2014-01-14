package soundcloud

// Copy this to config_test.go (its ignored by git)
// And fill in your test clientID, clientSecret, authorization token and secret
// You can get the client information from http://developers.soundcloud.com/you/apps
// And you'll have to authorize with that client to get a token
// You don't NEED access tokens per se but some requests require it.

var TestConfig map[string]interface{} = map[string]interface{}{
	"client_id":     "f6212fc8e1e90501e7cece44f78b8f43",
	"client_secret": "45d1d26f8399fc4c19abcf14d3830ff2",
	"access_token":  "1-64086-73163832-dc46009aebdb123",
	"refresh_token": "7b66f2a4a11add6d62fe25dd9d8ee785",
	"my_id":         uint64(73163832),
}
