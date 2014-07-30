package soundcloud

import (
	"testing"
)

func TestMe(t *testing.T) {
	authorizedRequest(t)
	res, err := api.Me().Get(nil)
	if err != nil {
		t.Error(err)
	} else if res.Id != TestConfig["my_id"].(uint64) {
		t.Error("Idcheck failed", res)
	}
}
