package openchain

import (
	"context"
	"net/http"

	"github.com/dbadoy/signature/internal/option"
)

type LookupV1Options struct {
	Method string `json:"function"`
	Event  string `json:"event"`
	Filter bool   `json:"filter"`
}

func (s *LookupV1Options) Encode() (string, error) {
	return option.EncodeQueryParam(s), nil
}

// Lookup returns a 200 response with a null value, even
// if the retrieved value does not exist in the database,
// and MUST include '0x' prefix.
//
// Recommend using the 'Signature' method.
func (c *Client) LookupV1(ctx context.Context, opts LookupV1Options) (*SignatureResponse, int, error) {
	var (
		res        SignatureResponse
		statusCode int
		err        error
	)

	statusCode, err = c.doRequest(ctx, "lookup", http.MethodGet, &res, nil, &opts)
	return &res, statusCode, err
}
