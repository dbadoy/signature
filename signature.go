package signature

import "errors"

var (
	ErrInvalidID         = errors.New("invalid ID")
	ErrSignatureNotFound = errors.New("signature not found")
)

type Caller interface {
	// Signature returns the signature for methodID. It
	// doesn't matter if it contains "0x" or not.
	Signature(id string) ([]string, error)

	// SignatureWithBytes returns the signature for
	// methodID bytes.
	SignatureWithBytes(id []byte) ([]string, error)
}
