package file

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dbadoy/signature"
	"github.com/ethereum/go-ethereum/common"
)

var (
	BaseURL = "https://raw.githubusercontent.com/ethereum-lists/4bytes/master/signatures/"

	_ = signature.Caller(&Client{})
)

// This client gets the signature from the 'ethereum-lists' repository.
type Client struct {
	caller *http.Client
}

// timeout is in seconds, where 0 means no timeout.
func New(timeout time.Duration) (*Client, error) {
	return &Client{
		caller: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

// Signature returns the signature for methodID.
// It doesn't matter if it contains "0x" or not.
func (c *Client) Signature(id string) ([]string, error) {
	if id = parseID(id); id == "" {
		return nil, signature.ErrInvalidID
	}

	res, err := c.caller.Get(BaseURL + id)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		var err error
		switch res.StatusCode {
		case 404:
			err = signature.ErrSignatureNotFound
		default:
			err = fmt.Errorf("unknown status code %d (report me: https://github.com/dbadoy/signature)", res.StatusCode)
		}
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return splitBySeperator(string(b)), nil
}

// SignatureWithBytes returns the signature for methodID bytes.
func (c *Client) SignatureWithBytes(id []byte) ([]string, error) {
	return c.Signature(common.Bytes2Hex(id))
}

// parseID checks to see if it follows the correct methodID
// format, and returns "" if not.
func parseID(s string) string {
	if strings.HasPrefix(s, "0x") {
		s = strings.Trim(s, "0x")
	}
	if len(s) != 8 {
		return ""
	}
	return s
}

func splitBySeperator(s string) []string {
	// e.g. getTokenLockersForAccount(address);activateCollection(uint256)
	return strings.Split(s, ";")
}
