package fourbytes

const (
	BaseURL = "https://www.4byte.directory/api/"

	Version = "v1/"
)

type SignatureResponse struct {
	Count    int             `json:"count"`
	Next     *string         `json:"next"`
	Previous *string         `json:"previous"`
	Results  []SignatureInfo `json:"results"`
}

type SignatureInfo struct {
	ID             uint64 `json:"id"`
	CreatedAt      string `json:"created_at"`
	TextSignature  string `json:"text_signature"`
	HexSignature   string `json:"hex_signature"`
	BytesSignature string `json:"bytes_signature"`
}
