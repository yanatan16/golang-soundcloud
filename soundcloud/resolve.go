package soundcloud

import (
	"net/url"
)

func (api *Api) Resolve(scurl string) (uri *url.URL, err error) {
	params, err := api.extendParams(Values("url", scurl))
	if err != nil {
		return nil, err
	}

	req, err := buildGetRequest(urlify("/resolve"), params)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if urlerr, ok := err.(*url.Error); ok && urlerr.Err.Error() == "No redirects!" {
		err = nil
	} else if err != nil {
		return nil, err
	}

	loc, err := resp.Location()
	return loc, nil
}
