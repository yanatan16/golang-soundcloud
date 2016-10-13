package soundcloud

import (
	"testing"
)

func TestResolve(t *testing.T) {
	ret, err := api.Resolve("http://soundcloud.com/matas/hobnotropic")
	if err != nil {
		t.Error(err)
	} else if ret.Path != "/tracks/49931.json" {
		t.Error("resolve didn't work " + ret.String())
	}
}

func TestResolveEmptyUrl(t *testing.T) {
	_, err := api.Resolve("http://soundcloud.com/xxx/yyy")
	if err == nil {
		t.Error("err is nil")
	}

	if err.Error() != "empty location" {
		t.Error("wrong error")
	}
}
