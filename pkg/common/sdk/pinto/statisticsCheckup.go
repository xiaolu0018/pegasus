package pinto

import (
	"io"

	"github.com/1851616111/util/http"
)

func (p *PintoSDK) GetStatisticsCheckup(codes []string, ip string) (io.ReadCloser, error) {
	rsp, err := http.Send(&http.HttpSpec{
		URL:         ip + "/api/appoint/statistics",
		Method:      "GET",
		ContentType: http.ContentType_JSON,
		BodyParams:  http.NewBody().Add("salecodes", &codes),
	})
	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()
	return rsp.Body, nil
}
