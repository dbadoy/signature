package fourbytes

const (
	BaseURL = "https://www.4byte.directory/api/"

	Version = "v1/"
)

/*
[Example]
{
    "count": 1,
    "next": null,
    "previous": null,
    "results": [
        {
            "id": 238989,
            "created_at": "2021-10-12T07:54:58.838466Z",
            "text_signature": "getInitializationCodeFromContractRuntime_6CLUNS()",
            "hex_signature": "0x00000009",
            "bytes_signature": "\u0000\u0000\u0000\t"
        }
    ]
}
*/
type (
	SignatureResponse struct {
		Count    int             `json:"count"`
		Next     *string         `json:"next"`
		Previous *string         `json:"previous"`
		Results  []SignatureInfo `json:"results"`
	}

	SignatureInfo struct {
		ID             uint64 `json:"id"`
		CreatedAt      string `json:"created_at"`
		TextSignature  string `json:"text_signature"`
		HexSignature   string `json:"hex_signature"`
		BytesSignature string `json:"bytes_signature"`
	}
)
