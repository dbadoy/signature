package fourbytes

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

// This client gets the signature from the 4byte.directory API.
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

	// The 4byte.directory API handles the '0x' prefix on
	// its own, so '0xa9059cbb' and 'a9059cbb' result in
	// the same thing.
	// The reason this library truncates the '0x' prefix
	// is to check for legitimate requests (length
	// checking) and discard them.
	if strings.HasPrefix(id, "0x") {
		id = strings.Trim(id, "0x")
	}

	switch len(id) {
	case 8:
		resp, _, err = c.MethodSignatureV1(context.Background(), MethodSigV1Options{id})
	case 64:
		resp, _, err = c.EventSignatureV1(context.Background(), EventSigV1Options{id})
	default:
		err = signature.ErrInvalidID
	}

	if err != nil {
		return nil, err
	}

	for _, siginfo := range resp.Results {
		sigs = append(sigs, siginfo.TextSignature)
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
