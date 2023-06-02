package fourbytes

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

// This client gets the signature from the 4byte.directory API.
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

	// Returns a 200 response with a null value, even if
	// the retrieved value does not exist in the database.
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
