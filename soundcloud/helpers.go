package soundcloud

import (
	"net/url"
)

func Values(keyValues ...string) url.Values {
	v := url.Values{}
	for i := 0; i < len(keyValues)-1; i += 2 {
		v.Set(keyValues[i], keyValues[i+1])
	}
	return v
}

func (api *Api) Values(keyValues ...string) url.Values {
	return Values(keyValues...)
}