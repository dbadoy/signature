package fourbytes

import (
	"context"
	"net/http"
	"strings"

	"github.com/dbadoy/signature"
	"github.com/dbadoy/signature/internal/option"
)

type EventSigV1Options struct {
	// The event ID to search for.
	//
	// It doesn't matter include '0x' or not.
	HexSignature string `json:"hex_signature"`
}

func (s *EventSigV1Options) Encode() (string, error) {
	// Disallows full search.
	if s.HexSignature == "" || s.HexSignature == "0x" {
		return "", signature.ErrRequiredMissing
	}

	if strings.HasPrefix(s.HexSignature, "0x") {
		s.HexSignature = strings.Trim(s.HexSignature, "0x")
	}

	// Disallows LIKE search.
	//
	// '4bytes.directory' returns an 'LIKE search'
	// result if no 64 characters are entered.
	if len(s.HexSignature) != 64 {
		return "", signature.ErrInvalidID
	}

	return option.EncodeQueryParam(s), nil
}

func (c *Client) EventSignatureV1(ctx context.Context, opts EventSigV1Options) (*SignatureResponse, int, error) {
	var (
		res        SignatureResponse
		statusCode int
		err        error
	)

	statusCode, err = c.doRequest(ctx, v1, "/event-signatures", http.MethodGet, &res, nil, &opts)
	return &res, statusCode, err
}
