package openchain

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/dbadoy/signature"
	"github.com/dbadoy/signature/internal/option"
	"github.com/dbadoy/signature/internal/utils"
	"github.com/ethereum/go-ethereum/common"
)

var (
	_ = signature.Caller(&Client{})
)

// This client gets the signature from the openchainxyz API.
type Client struct {
	cfg    *Config
	caller *http.Client
}

// timeout is in seconds, where 0 means no timeout.
func New(cfg *Config) (*Client, error) {
	if cfg == nil {
		cfg = DefaultConfig()
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &Client{
		cfg: cfg,
		caller: &http.Client{
			Timeout: cfg.Timeout,
		},
	}, nil

}

// Signature doesn't distinguish between events and methods
// as parameters, it just takes them in and routes them.
//
// It doesn't matter include '0x' or not.
func (c *Client) Signature(id string) ([]string, error) {
	var (
		resp *SignatureResponse
		err  error

		sigs []string
	)

	// Must contain '0x'
	if !strings.HasPrefix(id, "0x") {
		id = "0x" + id
	}

	switch len(id) {
	case 8 + 2 /* include '0x' */ :
		resp, _, err = c.LookupV1(context.Background(), LookupV1Options{Method: id, Filter: true})
	case 64 + 2:
		resp, _, err = c.LookupV1(context.Background(), LookupV1Options{Event: id, Filter: true})
	default:
		err = signature.ErrInvalidID
	}

	if err != nil {
		return nil, err
	}

	var items []Item
	e, ok := resp.Result.Event[id]
	if ok {
		items = e
	}
	m, ok := resp.Result.Method[id]
	if ok {
		items = m
	}

	for _, item := range items {
		sigs = append(sigs, item.Name)
	}

	// Lookup returns a 200 response with a null value, even
	// if the retrieved value does not exist in the database.
	//
	// Signature defines this as an error.
	if len(sigs) == 0 {
		return nil, signature.ErrSignatureNotFound
	}

	return sigs, nil
}

// SignatureWithBytes returns the signature for methodID bytes.
func (c *Client) SignatureWithBytes(id []byte) ([]string, error) {
	return c.Signature(common.Bytes2Hex(id))
}

func (c *Client) doRequest(ctx context.Context, version, api, method string, response interface{}, body io.Reader, opts option.Option) (int, error) {
	var url = fmt.Sprintf("%s%s%s", BaseURL, version, api)
	if opts != nil {
		query, err := opts.Encode()
		if err != nil {
			return 0, err
		}
		url += fmt.Sprintf("?%s", query)
	}

	return utils.DoRequestWithContext(ctx, c.caller, url, method, response, body)
}
