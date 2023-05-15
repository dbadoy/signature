package openchain

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dbadoy/signature"
	"github.com/dbadoy/signature/pkg/option"
	"github.com/ethereum/go-ethereum/common"
)

var (
	_ = signature.Caller(&Client{})
)

// This client gets the signature from the openchainxyz API.
type Client struct {
	version string
	caller  *http.Client
}

// timeout is in seconds, where 0 means no timeout.
func New(version string /* This parameter has no meaning until a later version */, timeout time.Duration) (*Client, error) {
	return &Client{
		version: Version, /* Fixed value: 'v1/' */
		caller: &http.Client{
			Timeout: timeout,
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

	if items, ok := resp.Result.Event[id]; ok {
		for _, item := range items {
			sigs = append(sigs, item.Name)
		}
	}

	if items, ok := resp.Result.Method[id]; ok {
		for _, item := range items {
			sigs = append(sigs, item.Name)
		}
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

func (c *Client) doRequest(ctx context.Context, api, method string, response interface{}, body io.Reader, opt option.Option) (int, error) {
	var (
		query string
		err   error
		url   = fmt.Sprintf("%s%s%s", BaseURL, c.version, api)
	)

	if opt != nil {
		query, err = opt.Encode()
		if err != nil {
			return 0, err
		}
		url += fmt.Sprintf("?%s", query)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return 0, err
	}

	resp, err := c.caller.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return statusCode, err
	}

	if statusCode == 200 {
		if err := json.Unmarshal(respBody, &response); err != nil {
			return 0, err
		}
		return statusCode, nil
	}

	return statusCode, errors.New(string(respBody))
}
