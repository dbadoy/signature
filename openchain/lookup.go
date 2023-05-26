package openchain

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/dbadoy/signature"
	"github.com/dbadoy/signature/internal/option"
)

// The original openchain API works like this.
//
//     - A comma-delimited list of function hashes with leading 0x prefix.
//     Method: 0xa9059cbb,0x000000fa,0x00000009
// 	   --> {
// 		   "0x00000009": [
// 			   {
// 			      "name": "getInitializationCodeFromContractRuntime_6CLUNS()",
// 			      "filtered": false
// 			   }
// 		   ],
// 		   "0x000000fa": [
// 			   {
// 			      "name": "getDexAggKeeperWhitelistPosition_IkFc(address)",
// 			      "filtered": false
// 			   }
// 		   ],
// 		   "0xa9059cbb": [
// 			   {
// 			      "name": "transfer(address,uint256)",
// 			      "filtered": false
// 			   }
// 		   ]
// 	   }
//
// However, this library only supports a single query. Just note.

type LookupV1Options struct {
	// You can only retrieve one or the other of Method
	// and Event in a single request.
	//
	// It doesn't matter include '0x' or not.
	Method string `json:"function"`
	Event  string `json:"event"`

	// Whether or not to filter out junk results.
	Filter bool `json:"filter"`
}

func (s *LookupV1Options) Encode() (string, error) {
	if s.Method == "" && s.Event == "" {
		return "", signature.ErrRequiredMissing
	}

	// Only one search is allowed at a time.
	if s.Method != "" && s.Event != "" {
		return "", errors.New("disallowed multiple searches")
	}

	var (
		hex    *string
		expect int
	)

	if s.Method != "" {
		hex = &s.Method
		expect = 8 + 2 /* Include '0x' */
	} else {
		hex = &s.Event
		expect = 64 + 2
	}

	if !strings.HasPrefix(*hex, "0x") {
		*hex = "0x" + *hex
	}

	if len(*hex) != expect {
		return "", signature.ErrInvalidID
	}

	return option.EncodeQueryParam(s), nil
}

// LookupV1 returns a 200 response with a null value, even
// if the retrieved value does not exist in the database.
// It doesn't matter include '0x' or not.
//
// Recommend using the 'Signature' method.
func (c *Client) LookupV1(ctx context.Context, opts LookupV1Options) (*SignatureResponse, int, error) {
	var (
		res        SignatureResponse
		statusCode int
		err        error
	)

	statusCode, err = c.doRequest(ctx, v1, "/lookup", http.MethodGet, &res, nil, &opts)
	return &res, statusCode, err
}
