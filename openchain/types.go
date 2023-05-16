package openchain

const (
	BaseURL = "https://api.openchain.xyz/signature-database/"

	Version = "v1/"
)

/*
[Example]
{
  "ok": true,
  "result": {
    "event": {},
    "function": {
      "0xa9059cbb": [
        {
          "name": "transfer(address,uint256)",
          "filtered": false
        }
      ]
    }
  }
}
*/
type (
	SignatureResponse struct {
		Ok     bool            `json:"ok"`
		Result SignatureResult `json:"result"`
	}

	SignatureResult struct {
		Event  map[string][]Item `json:"event"`
		Method map[string][]Item `json:"function"`
	}

	Item struct {
		Name     string `json:"name"`
		Filtered bool   `json:"filtered"`
	}
)

/*
[Example]
{
  "ok": true,
  "result": {
    "count": {
      "event": 373210,
      "function": 2366040
    }
  }
}
*/
type (
	StatsResponse struct {
		Ok     bool        `json:"ok"`
		Result CountResult `json:"result"`
	}

	CountResult struct {
		Count Count `json:"count"`
	}

	Count struct {
		Event  int `json:"event"`
		Method int `json:"function"`
	}
)
