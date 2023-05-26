package openchain

import (
	"context"
	"net/http"
)

func (c *Client) StatsV1(ctx context.Context) (*StatsResponse, int, error) {
	var (
		res        StatsResponse
		statusCode int
		err        error
	)

	statusCode, err = c.doRequest(ctx, v1, "/stats", http.MethodGet, &res, nil, nil)
	return &res, statusCode, err
}
