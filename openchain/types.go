package openchain

const (
	BaseURL = "https://api.openchain.xyz/signature-database/"

	Version = "v1/"
)

type SignatureResponse struct {
	Ok     bool   `json:"ok"`
	Result Result `json:"result"`
}

type Result struct {
	Event  map[string][]Item `json:"event,omitempty"`
	Method map[string][]Item `json:"function,omitempty"`
}

type Item struct {
	Name     string `json:"name"`
	Filtered bool   `json:"filtered"`
}
