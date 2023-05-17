package utils

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

func DoRequestWithContext(ctx context.Context, caller *http.Client, url, method string, response interface{}, body io.Reader) (int, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return 0, err
	}

	resp, err := caller.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return statusCode, err
	}

	if statusCode == http.StatusOK {
		if err := json.Unmarshal(respBody, &response); err != nil {
			return 0, err
		}
		return statusCode, nil
	}

	return statusCode, errors.New(string(respBody))
}
