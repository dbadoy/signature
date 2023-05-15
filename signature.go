package signature

import "errors"

var (
	ErrInvalidID = errors.New("invalid ID")
)

type Caller interface {
	// Signature returns the signature for methodID. It
	// doesn't matter if it contains "0x" or not.
	//
	// The 4byte.directory API handles the '0x' prefix on
	// its own, so '0xa9059cbb' and 'a9059cbb' result in
	// the same thing.
	// The reason this library truncates the '0x' prefix
	// is to check for legitimate requests (length
	// checking) and discard them.
	Signature(id string) ([]string, error)

	// SignatureWithBytes returns the signature for
	// methodID bytes.
	SignatureWithBytes(id []byte) ([]string, error)
}

type Provider struct {
	callers []Caller
}

func (p *Provider) Signature(id string) ([]string, error) {
	for _, caller := range p.callers {
		if sigs, err := caller.Signature(id); err == nil {
			return sigs, err
		}
	}
	return nil, errors.New("failed")
}

func (p *Provider) SignatureWithBytes(id []byte) ([]string, error) {
	for _, caller := range p.callers {
		if sigs, err := caller.SignatureWithBytes(id); err == nil {
			return sigs, err
		}
	}
	return nil, errors.New("failed")
}
