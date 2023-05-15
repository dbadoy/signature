package fourbytes

import (
	"context"
	"net/http"
	"strings"

	"github.com/dbadoy/signature"
	"github.com/dbadoy/signature/internal/option"
)

type MethodSigV1Options struct {
	HexSignature string `json:"hex_signature"`
}

func (s *MethodSigV1Options) Encode() (string, error) {
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
	// result if no 8 characters are entered.
	if len(s.HexSignature) != 8 {
		return "", signature.ErrInvalidID
	}

	return option.EncodeQueryParam(s), nil
}

func (c *Client) MethodSignatureV1(ctx context.Context, opts MethodSigV1Options) (*SignatureResponse, int, error) {
	var (
		res        SignatureResponse
		statusCode int
		err        error
	)

	statusCode, err = c.doRequest(ctx, "/signatures", http.MethodGet, &res, nil, &opts)
	return &res, statusCode, err
}
