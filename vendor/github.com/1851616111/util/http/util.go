package http

import "net/http"

func Send(spec *HttpSpec) (*http.Response, error) {
	req, err := NewRequest(spec)
	if err != nil {
		return nil, err
	}

	cli := http.Client{}
	return cli.Do(req)
}
